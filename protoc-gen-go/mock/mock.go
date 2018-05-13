package mock

import (
	"fmt"

	pb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

const generatedCodeVersion = 4

func init() {
	generator.RegisterPlugin(new(mock))
}

// mock is an implementation of the Go protocol buffer compiler's
// plugin architecture.  It generates bindings for mock support.
type mock struct {
	gen *generator.Generator
}

// Name returns the name of this plugin, "mock".
func (m *mock) Name() string {
	return "mock"
}

// Init initializes the plugin.
func (m *mock) Init(gen *generator.Generator) {
	m.gen = gen
}

// P forwards to g.gen.P.
func (m *mock) P(args ...interface{}) { m.gen.P(args...) }

// Generate generates code for the services in the given file.
func (m *mock) Generate(file *generator.FileDescriptor) {
	m.P("// Mock plugin start.")
	m.P()
	for _, desc := range m.gen.desc {
		// Don't generate virtual messages for maps.
		if desc.GetOptions().GetMapEntry() {
			continue
		}
		m.generateMessage(desc)
	}
}

// GenerateImports generates the import declaration for this file.
func (m *mock) GenerateImports(file *generator.FileDescriptor) {
	m.P("// GenerateImports")
}

// generateService generates all the code for the named service.
func (m *mock) generateService(file *generator.FileDescriptor, service *pb.ServiceDescriptorProto, index int) {
	path := fmt.Sprintf("6,%d", index) // 6 means service.
	m.P("// generateService", path)
}
