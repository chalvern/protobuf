package generator

import (
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// Generate the type and default constant definitions for this Descriptor.
func (g *Generator) generateMessageMock(message *Descriptor) {
	// The full type name
	typeName := message.TypeName()

	usedNames := make(map[string]bool)
	for _, n := range methodNames {
		usedNames[n] = true
	}
	// The full type name, CamelCased.
	ccTypeName := CamelCaseSlice(typeName)
	fieldNames := make(map[*descriptor.FieldDescriptorProto]string)
	fieldTypes := make(map[*descriptor.FieldDescriptorProto]string)
	// mapFieldTypes := make(map[*descriptor.FieldDescriptorProto]string)

	// mock comments
	g.P("// mock parts")
	tmpName := Annotate(message.file, message.path, ccTypeName)
	g.P("var ", tmpName, "Mock", " = ", tmpName, "{")
	g.In()

	// allocNames finds a conflict-free variation of the given strings,
	// consistently mutating their suffixes.
	// It returns the same number of strings.
	allocNames := func(ns ...string) []string {
	Loop:
		for {
			for _, n := range ns {
				if usedNames[n] {
					for i := range ns {
						ns[i] += "_"
					}
					continue Loop
				}
			}
			for _, n := range ns {
				usedNames[n] = true
			}
			return ns
		}
	}

	for _, field := range message.Field {
		base := CamelCase(*field.Name)
		ns := allocNames(base)
		fieldName := ns[0]
		typename, _ := g.GoType(message, field)

		fieldNames[field] = fieldName
		fieldTypes[field] = typename
		g.P(fieldName, " : ", valueOfType(field, typename), ",")
	}

	g.P("}")
	g.P("// end")
}

func valueOfType(field *descriptor.FieldDescriptorProto, typeName string) string {
	switch *field.Type {
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		return "true"
	case descriptor.FieldDescriptorProto_TYPE_FLOAT,
		descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		return "2018.0513"
	case descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32,
		descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		return "2018"
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		return "\"abcdefg\""
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		// This is only possible for oneof fields.
		return strings.Replace(typeName, "*", "&", -1) + "Mock"
	default:
		return "nil"
	}
}
