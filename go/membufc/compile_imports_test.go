package main

import (
	"bytes"
	"fmt"
	"github.com/orbs-network/pbparser"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"testing"
)

func TestCompilingEmptyProto(t *testing.T) {
	proto := pbparser.ProtoFile{
		PackageName: "foo",
	}

	compiled, err := compile(proto)
	if err != nil {
		t.Fatalf("Parsing compiled file has failed: %s", err)
	}

	if compiled.Name.Name != "foo" {
		t.Fatalf("Package name was '%s' and not 'foo'", compiled.Name.Name)
	}

	if len(compiled.Imports) > 0 {
		t.Fatalf("Expected 0 imports, got %s", importsToString(compiled.Imports))
	}
}

func TestCompilingProtoWithMessage(t *testing.T) {
	stringType, _ := pbparser.NewScalarDataType("string")
	proto := pbparser.ProtoFile{
		PackageName: "foo",
		Messages: []pbparser.MessageElement{
			{
				Name: "FooMessage",
				Fields: []pbparser.FieldElement{
					{
						Name: "aString",
						Type: stringType,
						Tag:  1,
					},
				},
			},
		},
	}

	compiled, err := compile(proto)
	if err != nil {
		t.Fatalf("Parsing compiled file has failed: %s", err)
	}

	assertImport(t, "bytes", compiled)
	assertImport(t, "fmt", compiled)
	assertImport(t, "github.com/orbs-network/membuffers/go", compiled)
}

func assertImport(t *testing.T, importName string, compiled *ast.File) {
	if !hasImport(compiled, importName) {
		t.Fatalf("Expected '%s' import, got %s", importName, importsToString(compiled.Imports))
	}
}

func hasImport(compiled *ast.File, importName string) bool {
	for _, spec := range compiled.Imports {
		if spec.Path.Value == fmt.Sprintf("\"%s\"", importName) {
			return true
		}
	}
	return false
}

func compile(proto pbparser.ProtoFile) (*ast.File, error) {
	w := bytes.Buffer{}
	compileProtoFile(&w, proto, make(map[string]dependencyData), MEMBUFC_VERSION, false)
	fset := token.NewFileSet()
	compiled, err := parser.ParseFile(fset, "", w.String(), parser.ImportsOnly)
	return compiled, err
}

func importsToString(specs []*ast.ImportSpec) string {
	w := bytes.Buffer{}
	fset := token.NewFileSet()
	for i, spec := range specs {
		format.Node(&w, fset, spec)
		if i < len(specs)-1 {
			w.WriteByte(',')
		}
	}

	return w.String()
}
