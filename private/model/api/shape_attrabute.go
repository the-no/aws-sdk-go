// +build codegen

package api

import (
	"bytes"
	"fmt"
	"text/template"
)

type ShapeAttrabute ShapeReference

var attrabuteGoCodeTmpls = template.Must(template.New("attrabuteGoCodeTmpls").Parse(`
{{ define "baseType" -}}
    return s.{{ .Name }}
{{- end }}
{{ define "structure" -}}
	if  s.{{ .Name }} != nil{
		 s.{{ .Name }}.Attrabute(attr)
	}
	return nil
}
{{- end }}
{{ define "list" -}}
 	 if len(s.{{ .Name }}) > 0 {
 	 	s.{{ .Name }}[0].Attrabute(attr)
 	 }
	return nil
{{- end }}
`))

// GoCode returns the generated Go code for the Shape with its validation type.
func (sa ShapeAttrabute) GoCode() string {
	var err error
	w := &bytes.Buffer{}
	switch sa.Type {
	case ShapeBaseTye:
		err = attrabuteGoCodeTmpls.ExecuteTemplate(w, "baseType", sa)
	case ShapeStructure:
		err = attrabuteGoCodeTmpls.ExecuteTemplate(w, "structure", sa)
	case ShapeList:
		err = attrabuteGoCodeTmpls.ExecuteTemplate(w, "list", sa)
	}

	if err != nil {
		panic(fmt.Sprintf("ShapeValidation.GoCode failed, err: %v", err))
	}
	return w.String()
}

// A ShapeValidations is a collection of shape validations needed nested within
// a parent shape
type ShapeAttrabutes []ShapeAttrabute

var attrabuteShapeTmpl = template.Must(template.New("attrabuteShapeTmpl").Parse(`
// Validate inspects the fields of the type to determine if they are valid.
func (s {{ .Shape.ShapeName }}) Attrabute(attr string) interface{} {
	switch attr {
  	{{ range $_, $v := .Attrabutes -}}
  		case "{{ $v.Name }}":
			{{ $v.GoCode }}
	{{ end }}
	}
	return nil
}
`))

// GoCode generates the Go code needed to perform validations for the
// shape and its nested fields.
func (sa ShapeAttrabutes) GoCode(shape *Shape) string {
	buf := &bytes.Buffer{}
	attrabuteShapeTmpl.Execute(buf, map[string]interface{}{
		"Shape":      shape,
		"Attrabutes": sa,
	})
	return buf.String()
}

// Has returns true or false if the ShapeValidations already contains the
// the reference and validation type.
func (sa ShapeAttrabutes) Has(ref *ShapeRef) bool {
	for _, v := range sa {
		if v.Ref == ref {
			return true
		}
	}
	return false
}
