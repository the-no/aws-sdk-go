// +build codegen

package api

import (
	"bytes"
	"fmt"
	"text/template"
)

// A ShapeValidationType is the type of validation that a shape needs
type ShapeReferenceType int

const (
	// ShapeValidationRequired states the shape must be set
	ShapeBaseTye = iota

	ShapeStructure
	ShapeList
)

// A ShapeValidation contains information about a shape and the type of validation
// that is needed
type ShapeReference struct {
	// Name of the shape to be validated
	Name string
	// Reference to the shape within the context the shape is referenced
	Ref *ShapeRef
	// Type of validation needed
	Type ShapeReferenceType
}

var referenceGoCodeTmpls = template.Must(template.New("referenceGoCodeTmpls").Parse(`
{{ define "baseType" -}}
    	return   s.{{ .Name }}
{{- end }}
{{ define "structure" -}}
	if  s.{{ .Name }} != nil{
		 s.{{ .Name }}.Reference()
	}
	return   ""
}
{{- end }}
{{ define "list" -}}
 	 if len(s.{{ .Name }}) > 0 {
 	 	s.{{ .Name }}[0].Reference()
 	 }
	return   ""
{{- end }}
`))

// GoCode returns the generated Go code for the Shape with its validation type.
func (sr ShapeReference) GoCode(shape *Shape) string {
	var err error

	w := &bytes.Buffer{}
	switch sr.Type {
	case ShapeBaseTye:
		err = referenceGoCodeTmpls.ExecuteTemplate(w, "baseType", sr)
	case ShapeStructure:
		err = referenceGoCodeTmpls.ExecuteTemplate(w, "structure", sr)
	case ShapeList:
		err = referenceGoCodeTmpls.ExecuteTemplate(w, "list", sr)
	}

	if err != nil {
		panic(fmt.Sprintf("ShapeValidation.GoCode failed, err: %v", err))
	}

	return w.String()
}

/*var ReferenceShapeTmpl = template.Must(template.New("ReferenceShape").Parse(`
// Validate inspects the fields of the type to determine if they are valid.
func (s *{{ .ShapeName }}) Reference() interface{} {
	{{ .Shape.ReferenceAction.GoCode }}
	}
`))*/
