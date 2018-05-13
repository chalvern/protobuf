package generator

import (
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
	g.P("var ", tmpName, "Mock", " = ", tmpName)
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
		g.P(Annotate(message.file, fieldName), "\t", typename)
	}

	g.P("}")
	g.P("// end")
}
