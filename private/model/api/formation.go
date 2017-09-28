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
type FormationCreator struct {
	API           *API `json:"-"`
	Name          string
	OperationName string `json:"operation"`
	Operations    map[string]*Formation
	SortOpts      []*Formation
	Referencer    string
	Attrabuter    string
}

type FormationWaiter struct {
	Name      string
	Input     string
	Waiter    *Waiter
	Arguments []*Argument
}

type Formation struct {
	Input         string
	Output        string
	Next          string
	NextFormation *Formation `json:"-"`
	Operation     *Operation `json:"-"`
	Arguments     []*Argument
	Waiter        *FormationWaiter
}

// WaitersGoCode generates and returns Go code for each of the waiters of
// this API.
func (a *API) FormationCreatorGoCode() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "import (\n%q\n\n%q\n%q\n)",
		"time",
		"github.com/the-no/aws-sdk-go/aws",
		"github.com/the-no/aws-sdk-go/aws/request",
	)

	for _, c := range a.Creators {
		buf.WriteString(c.GoCode())
	}
	return buf.String()
}

// used for unmarshaling from the waiter JSON file
type creatorDefinitions struct {
	*API
	Creators map[string]*FormationCreator
}

// AttachWaiters reads a file of waiter definitions, and adds those to the API.
// Will panic if an error occurs.
func (a *API) AttachFormationCreators(filename string) {
	p := creatorDefinitions{API: a}
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

func (p *creatorDefinitions) setup() {
	p.API.Creators = make(map[string]*FormationCreator)
	/*	i, keys := 0, make([]string, len(p.Creators))
		for k := range p.Creators {
			keys[i] = k
			i++
		}*/
	//sort.Strings(keys)
	for k, e := range p.Creators {
		e.API = p.API
		e := p.Creators[k]
		p.API.Creators[k] = e
		nextopt := e.OperationName
		e.Name = strings.Replace(k, "::", "", -1)
		e.SortOpts = make([]*Formation, 0, len(e.Operations))
		e.OperationName = p.ExportableName(e.OperationName)

		for range e.Operations {
			if o, ok := e.Operations[nextopt]; ok {
				o.Operation = p.API.Operations[nextopt]
				o.Input = strings.ToLower(o.Operation.InputRef.ShapeName)
				o.Output = strings.ToLower(o.Operation.OutputRef.ShapeName)

				if o.Waiter != nil {
					o.Waiter.Waiter = p.API.waitersMap[o.Waiter.Name]
					o.Waiter.Input = strings.ToLower(o.Waiter.Waiter.Operation.InputRef.ShapeName)
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
		fmt.Println("--------------------", e.OperationName, len(e.SortOpts))
	}
}

func (p *creatorDefinitions) initFormation() {

}

/*func (p *creatorDefinitions) setup() {
	p.API.Creators = make(map[string]*FormationCreator)
	i, keys := 0, make([]string, len(p.Creators))
	for k := range p.Creators {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, n := range keys {
		e := p.Creators[n]
		e.API = p.API
		n = p.ExportableName(n)
		e.Name = strings.Replace(n, "::", "", -1)
		e.OperationName = p.ExportableName(e.OperationName)
		for k, o := range e.Operations {
			o.Operation = p.API.Operations[k]
			o.Input = strings.ToLower(o.Operation.InputRef.ShapeName)
			o.Output = strings.ToLower(o.Operation.OutputRef.ShapeName)

			if o.Waiter != nil {
				o.Waiter.Waiter = p.API.waitersMap[o.Waiter.Name]
				o.Waiter.Input = strings.ToLower(o.Waiter.Waiter.Operation.InputRef.ShapeName)
				for _, a := range o.Waiter.Arguments {
					inputs := strings.Split(a.Input, ".")
					if inputs[1] == "Output" {
						a.Input = e.Operations[inputs[0]].Output
					} else {
						a.Input = e.Operations[inputs[0]].Input
					}
				}

			}
		}

		for _, o := range e.Operations {
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
		p.API.Creators[n] = e
	}
}
*/
var creatorTmpls = template.Must(template.New("creatorTmpls").Funcs(
	template.FuncMap{
		"titleCase": func(v string) string {
			return strings.Title(v)
		},
		"toLower": func(v string) string {
			return strings.ToLower(v)
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
	return  nil,nil,nil
}
{{- end }}
`))

/*var creatorTmpls = template.Must(template.New("creatorTmpls").Funcs(
	template.FuncMap{
		"titleCase": func(v string) string {
			return strings.Title(v)
		},
		"toLower": func(v string) string {
			return strings.ToLower(v)
		},
	},
).Parse(`
{{ define "creator"}}
{{ $nextOpt := index  .Operations .OperationName -}}
// create{{ .Name }} uses the {{ $nextOpt.Operation.API.NiceName }} API operation
// {{ .OperationName }} to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *{{ .API.StructName }}) create{{ .Name }}(input {{ $nextOpt.Operation.InputRef.GoType }}) (r aws.Referencer,attr aws.Attrabuter,err error) {

	{{ $context := . -}}
	{{ range $_, $_ := .Operations -}}
	  	{{ if eq $nextOpt.Operation.Name $context.OperationName -}}
			{{ toLower $nextOpt.Operation.InputRef.ShapeName -}} := input
		{{ else }}
			{{ range $_, $arg := $nextOpt.Arguments -}}

			{{ end -}}
		{{ end -}}
		{{ toLower $nextOpt.Operation.OutputRef.ShapeName -}} ,err := {{ $nextOpt.Operation.ExportedName -}}({{ toLower $nextOpt.Operation.InputRef.ShapeName -}})
		if err == nil {
			{{ if $nextOpt.Waiter -}}
   			{{ toLower $nextOpt.Waiter.Waiter.Operation.InputRef.ShapeName }} := &{{ $nextOpt.Waiter.Waiter.Operation.InputRef.ShapeName }}{}
   			{{ range $_, $arg := $nextOpt.Waiter.Arguments -}}

			{{ end -}}
   		{{- end }}
		}else{
			return nil,nil,err
		}

   	{{- end }}
	return  nil,nil,nil
}
{{- end }}
`))*/

// InterfaceSignature returns a string representing the Waiter's interface
// function signature.
/*func (f *FormationCreator) InterfaceSignature() string {
	var buf bytes.Buffer
	if err := waiterTmpls.ExecuteTemplate(&buf, "waiter interface", f); err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}*/

// GoCode returns the generated Go code for an individual waiter.
func (f *FormationCreator) GoCode() string {
	var buf bytes.Buffer
	if err := creatorTmpls.ExecuteTemplate(&buf, "creator", f); err != nil {
		panic(err)
	}

	return buf.String()
}
