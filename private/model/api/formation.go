// +build codegen

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	//	"sort"
	"strings"
	"text/template"
)

// WaiterAcceptor is the acceptors defined in the model the SDK will use
// to wait on resource states with.
type Argument struct {
	Key   string
	Input string
	Value string
}

// ExpectedString returns the string that was expected by the WaiterAcceptor
/*func (a *WaiterAcceptor) ExpectedString() string {
	switch a.Expected.(type) {
	case string:
		return fmt.Sprintf("%q", a.Expected)
	default:
		return fmt.Sprintf("%v", a.Expected)
	}
}*/

// A Waiter is an individual waiter definition.
type Formation struct {
	API           *API `json:"-"`
	Name          string
	OperationName string `json:"operation"`
	Operations    map[string]*FormationOpt
	SortOpts      []*FormationOpt
	Referencer    string
	Attrabuter    string
}

type FormationWaiter struct {
	Name      string
	Input     string
	Waiter    *Waiter
	Arguments []*Argument
}

type FormationOpt struct {
	Input         string
	Output        string
	Next          string
	NextFormation *FormationOpt `json:"-"`
	Operation     *Operation    `json:"-"`
	Arguments     []*Argument
	Waiter        *FormationWaiter
}

// WaitersGoCode generates and returns Go code for each of the waiters of
// this API.
func (a *API) FormationGoCode() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "import (\n%q\n\n%q\n%q\n)",
		"time",
		"github.com/the-no/aws-sdk-go/aws",
		"github.com/the-no/aws-sdk-go/aws/request",
	)

	for _, c := range a.Creators {
		buf.WriteString(c.GoCreatorCode())
	}

	for _, c := range a.Deleters {
		buf.WriteString(c.GoDeleterCode())
	}
	return buf.String()
}

// used for unmarshaling from the waiter JSON file
type creatorDeleterDefinitions struct {
	*API
	Creators map[string]*Formation
	Deleters map[string]*Formation
}

// AttachWaiters reads a file of waiter definitions, and adds those to the API.
// Will panic if an error occurs.
func (a *API) AttachFormations(filename string) {
	p := creatorDeleterDefinitions{API: a}
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	err = json.NewDecoder(f).Decode(&p)
	if err != nil {
		panic(err)
	}

	p.setup()
}

func (p *creatorDeleterDefinitions) setup() {
	p.API.Creators = make([]*Formation, 0, len(p.Creators))
	p.API.Deleters = make([]*Formation, 0, len(p.Deleters))
	for k, e := range p.Creators {
		e.Name = strings.Replace(k, "::", "", -1)
		p.resolveFormation(e)
		refs := strings.Split(e.Referencer, ".")
		if refs[1] == "Output" {
			e.Referencer = e.Operations[refs[0]].Output
		} else {
			e.Referencer = e.Operations[refs[0]].Input
		}

		attrs := strings.Split(e.Attrabuter, ".")
		if attrs[1] == "Output" {
			e.Attrabuter = e.Operations[attrs[0]].Output
		} else {
			e.Attrabuter = e.Operations[attrs[0]].Input
		}
		p.API.Creators = append(p.API.Creators, e)

	}

	for k, e := range p.Deleters {
		e.Name = strings.Replace(k, "::", "", -1)
		p.resolveFormation(e)
		p.API.Deleters = append(p.API.Deleters, e)
	}
}

func (p *creatorDeleterDefinitions) resolveFormation(e *Formation) {
	e.API = p.API
	nextopt := e.OperationName
	e.SortOpts = make([]*FormationOpt, 0, len(e.Operations))
	e.OperationName = p.ExportableName(e.OperationName)

	for range e.Operations {
		if o, ok := e.Operations[nextopt]; ok {
			o.Operation = p.API.Operations[nextopt]
			o.Input = strings.ToLower(o.Operation.InputRef.ShapeName)
			o.Output = strings.ToLower(o.Operation.OutputRef.ShapeName)

			if o.Waiter != nil {
				o.Waiter.Waiter = p.API.waitersMap[o.Waiter.Name]
				o.Waiter.Input = strings.ToLower(o.Waiter.Waiter.Operation.InputRef.ShapeName)
				fmt.Println(o)
				for _, a := range o.Waiter.Arguments {
					inputs := strings.Split(a.Input, ".")
					if inputs[1] == "Output" {
						a.Input = e.Operations[inputs[0]].Output
					} else {
						a.Input = e.Operations[inputs[0]].Input
					}
				}

			}

			nextopt = o.Next
			e.SortOpts = append(e.SortOpts, o)
			o.NextFormation = e.Operations[o.Next]
			if len(o.Arguments) > 0 {
				for _, a := range o.Arguments {
					inputs := strings.Split(a.Input, ".")
					if inputs[1] == "Output" {
						a.Input = e.Operations[inputs[0]].Output
					} else {
						a.Input = e.Operations[inputs[0]].Input
					}
				}
			}
		}
	}
}

var creatorTmpls = template.Must(template.New("creatorTmpls").Funcs(
	template.FuncMap{
		"titleCase": func(v string) string {
			return strings.Title(v)
		},
	},
).Parse(`
{{ define "creator"}}
{{ $firstOpt := index  .Operations .OperationName -}}
// create{{ .Name }} uses the {{ $firstOpt.Operation.API.NiceName }} API operation
// {{ .OperationName }} to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *{{ .API.StructName }}) create{{ .Name }}(input {{ $firstOpt.Operation.InputRef.GoType }}) (r aws.Referencer,attr aws.Attrabuter,err error) {
	
	{{ $firstName := .OperationName -}}
	{{ range $_, $nextOpt := .SortOpts -}}
	  	{{ if eq $nextOpt.Operation.Name $firstName -}}
			{{ $nextOpt.Input -}} := input
		{{ else }}
			{{ $nextOpt.Input}} := &{{ $nextOpt.Operation.InputRef.ShapeName }}{}
			{{ range $_, $arg := $nextOpt.Arguments -}}
				if err := awsutil.CopyValue({{ $nextOpt.Input }} ,"{{ $arg.Key -}}",{{ $arg.Input }},"{{ $arg.Value }}");
				 err != nil {
					return nil,nil,err
				}
			{{ end -}}
		{{ end -}}
		{{ $nextOpt.Output -}} ,err := {{ $nextOpt.Operation.ExportedName -}}({{ $nextOpt.Input -}})
		if err == nil {
			{{ if $nextOpt.Waiter -}}
   			{{ $nextOpt.Waiter.Input}} := &{{ $nextOpt.Waiter.Waiter.Operation.InputRef.ShapeName }}{}
   			{{ range $_, $arg := $nextOpt.Waiter.Arguments -}}
				if err := awsutil.CopyValue({{ $nextOpt.Waiter.Input }} ,"{{ $arg.Key -}}",{{ $arg.Input }},"{{ $arg.Value }}");
				 err != nil {
					return nil,nil,err
				}
			{{ end -}}
			if err :=  WaitUntil{{ $nextOpt.Waiter.Waiter.Name }}({{ $nextOpt.Waiter.Input}});err != nil{
				return nil,nil,err
			}
   		{{- end }}

		}else{
			return nil,nil,err
		}
   	{{- end }}
	return  {{ .Referencer }},{{ .Attrabuter }},nil
}
{{- end }}
`))

var deleterTmpls = template.Must(template.New("deleterTmpls").Funcs(
	template.FuncMap{
		"titleCase": func(v string) string {
			return strings.Title(v)
		},
	},
).Parse(`
{{ define "deleter"}}
{{ $firstOpt := index  .Operations .OperationName -}}
// delete{{ .Name }} uses the {{ $firstOpt.Operation.API.NiceName }} API operation
// {{ .OperationName }} to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *{{ .API.StructName }}) delete{{ .Name }}(input {{ $firstOpt.Operation.InputRef.GoType }}) (err error) {
	
	{{ $firstName := .OperationName -}}
	{{ range $_, $nextOpt := .SortOpts -}}
	  	{{ if eq $nextOpt.Operation.Name $firstName -}}
			{{ $nextOpt.Input -}} := input
		{{ else }}
			{{ $nextOpt.Input}} := &{{ $nextOpt.Operation.InputRef.ShapeName }}{}
			{{ range $_, $arg := $nextOpt.Arguments -}}
				if err := awsutil.CopyValue({{ $nextOpt.Input }} ,"{{ $arg.Key -}}",{{ $arg.Input }},"{{ $arg.Value }}");
				 err != nil {
					return err
				}
			{{ end -}}
		{{ end -}}
		{{ $nextOpt.Output -}} ,err := {{ $nextOpt.Operation.ExportedName -}}({{ $nextOpt.Input -}})
		if err == nil {
			{{ if $nextOpt.Waiter -}}
   			{{ $nextOpt.Waiter.Input}} := &{{ $nextOpt.Waiter.Waiter.Operation.InputRef.ShapeName }}{}
   			{{ range $_, $arg := $nextOpt.Waiter.Arguments -}}
				if err := awsutil.CopyValue({{ $nextOpt.Waiter.Input }} ,"{{ $arg.Key -}}",{{ $arg.Input }},"{{ $arg.Value }}");
				 err != nil {
					return err
				}
			{{ end -}}
			if err :=  WaitUntil{{ $nextOpt.Waiter.Waiter.Name }}({{ $nextOpt.Waiter.Input}});err != nil{
				return err
			}
   		{{- end }}
		}else{
			return err
		}
   	{{- end }}
	return nil
}
{{- end }}
`))

// GoCode returns the generated Go code for an individual waiter.
func (f *Formation) GoCreatorCode() string {
	var buf bytes.Buffer
	if err := creatorTmpls.ExecuteTemplate(&buf, "creator", f); err != nil {
		panic(err)
	}

	return buf.String()
}

func (f *Formation) GoDeleterCode() string {
	var buf bytes.Buffer
	if err := deleterTmpls.ExecuteTemplate(&buf, "deleter", f); err != nil {
		panic(err)
	}

	return buf.String()
}
