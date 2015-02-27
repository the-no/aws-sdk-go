package api

import (
	"sort"
	"strings"

	"github.com/awslabs/aws-sdk-go/internal/util"
)

type ShapeRef struct {
	API           *API   `json: "-"`
	Shape         *Shape `json: "-"`
	Documentation string
	ShapeName     string `json:"shape"`
	Location      string
	LocationName  string
	QueryName     string
	Flattened     bool
	ResultWrapper string
	XMLNamespace  XMLInfo
}

type XMLInfo struct {
	Prefix string
	URI    string
}

type Shape struct {
	API           *API `json: "-"`
	ShapeName     string
	Documentation string
	MemberRefs    map[string]*ShapeRef `json:"members"`
	MemberRef     ShapeRef             `json:"member"`
	KeyRef        ShapeRef             `json:"key"`
	ValueRef      ShapeRef             `json:"value"`
	Required      []string
	Payload       string
	Type          string
	Exception     bool
	Enum          []string
	Flattened     bool
	ResultWrapper string
	Streaming     bool

	refs []*ShapeRef
}

func (s *Shape) Rename(newName string) {
	for _, r := range s.refs {
		r.ShapeName = newName
	}

	delete(s.API.Shapes, s.ShapeName)
	s.API.Shapes[newName] = s
	s.ShapeName = newName
}

func (s *Shape) MemberNames() []string {
	i, names := 0, make([]string, len(s.MemberRefs))
	for n, _ := range s.MemberRefs {
		names[i] = n
		i++
	}
	sort.Strings(names)
	return names
}

func (s *Shape) GoTypeElem() string {
	switch s.Type {
	case "structure":
		return s.ShapeName
	case "map":
		return "map[" + s.KeyRef.GoTypeElem() + "]" + s.ValueRef.GoType()
	case "list":
		return "[]" + s.MemberRef.GoType()
	case "boolean":
		return "bool"
	case "string", "character":
		return "string"
	case "blob":
		return "[]byte"
	case "integer":
		return "int"
	case "long":
		return "int64"
	case "float":
		return "float32"
	case "double":
		return "float64"
	case "timestamp":
		s.API.imports["time"] = true
		return "time.Time"
	default:
		panic("Unsupported shape type: " + s.Type)
	}
}

func (s *Shape) GoType() string {
	t := s.GoTypeElem()
	if !strings.HasPrefix(t, "[]") {
		t = "*" + t
	}
	return t
}

func (ref *ShapeRef) GoTypeElem() string {
	return ref.Shape.GoTypeElem()
}

func (ref *ShapeRef) GoType() string {
	return ref.Shape.GoType()
}

func (ref *ShapeRef) GoTags(toplevel bool) string {
	code := "`"
	if ref.Location != "" {
		code += `location:"` + ref.Location + `" `
	}
	if ref.LocationName != "" {
		code += `locationName:"` + ref.LocationName + `" `
	}
	if ref.QueryName != "" {
		code += `queryName:"` + ref.QueryName + `" `
	}
	if ref.Shape.MemberRef.LocationName != "" {
		code += `locationNameList:"` + ref.Shape.MemberRef.LocationName + `" `
	}
	if ref.Shape.KeyRef.LocationName != "" {
		code += `locationNameKey:"` + ref.Shape.KeyRef.LocationName + `" `
	}
	if ref.Shape.ValueRef.LocationName != "" {
		code += `locationNameValue:"` + ref.Shape.ValueRef.LocationName + `" `
	}
	code += `type:"` + ref.Shape.Type + `" `

	// embed the timestamp type for easier lookups
	if ref.Shape.Type == "timestamp" {
		code += `timestampFormat:"`
		if ref.Location == "header" {
			code += "rfc822"
		} else {
			switch ref.API.Metadata.Protocol {
			case "json", "rest-json":
				code += "unix"
			case "rest-xml", "ec2", "query":
				code += "iso8601"
			}
		}
		code += `" `
	}

	if ref.Shape.Flattened || ref.Flattened {
		code += `flattened:"true" `
	}

	if toplevel {
		if ref.ResultWrapper != "" {
			code += `resultWrapper:"` + ref.ResultWrapper + `" `
		} else if ref.Shape.ResultWrapper != "" {
			code += `resultWrapper:"` + ref.Shape.ResultWrapper + `" `
		}
		if ref.Shape.Payload != "" {
			code += `payload:"` + ref.Shape.Payload + `" `
		}
		if len(ref.Shape.Required) > 0 {
			code += `required:"` + strings.Join(ref.Shape.Required, ",") + `" `
		}
	}

	if ref.Location != "" { // omit non-body location elements from JSON/XML
		code += `json:"-" xml:"-" `
	} else if strings.Contains(ref.API.Metadata.Protocol, "json") {
		code += `json:",omitempty"`
	}

	return strings.TrimSpace(code) + "`"
}

func (s *Shape) GoCode() string {
	code := "type " + s.ShapeName + " "
	switch s.Type {
	case "structure":
		code += "struct {\n"
		for _, n := range s.MemberNames() {
			m := s.MemberRefs[n]
			if s.Streaming && s.Payload == n {
				code += n + " io.ReadSeeker " + m.GoTags(false) + "\n"
			} else {
				code += n + " " + m.GoType() + " " + m.GoTags(false) + "\n"
			}
		}
		metaStruct := "metadata" + s.ShapeName
		ref := &ShapeRef{ShapeName: s.ShapeName, API: s.API, Shape: s}
		code += "\n" + metaStruct + "  `json:\"-\", xml:\"-\"`\n"
		code += "}\n\n"
		code += "type " + metaStruct + " struct {\n"
		code += "SDKShapeTraits bool " + ref.GoTags(true)
		code += "}"
	default:
		panic("Cannot generate toplevel shape for " + s.Type)
	}

	return util.GoFmt(code)
}