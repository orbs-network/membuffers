package main

import (
	"io"
	"github.com/orbs-network/pbparser"
	"fmt"
	"os"
)

var globalInlineTypes = make(map[string]string)

type InlineType struct {
	Name string
	Alias string
	FieldGoType string
}

func compileInlineFile(w io.Writer, file pbparser.ProtoFile, dependencyData map[string]dependencyData, compilerVersion string) {
	inlines := []InlineType{}
	hasBytes := false
	for _, m := range file.Messages {
		if len(m.Options) == 1 && m.Options[0].Name == "inline_type" {
			goTypes := map[string]string{
				"bytes":  "[]byte",
				"string": "string",
				"uint8":  "uint8",
				"uint16": "uint16",
				"uint32": "uint32",
				"uint64": "uint64",
			}
			fieldGoType, found := goTypes[m.Options[0].Value]
			if found {
				if fieldGoType == "[]byte" {
					hasBytes = true
				}
				inlines = append(inlines, InlineType{
					Name: convertFieldNameToGoCase(m.Name),
					Alias: m.Options[0].Value,
					FieldGoType: fieldGoType,
				})
			}
		}
	}
	t := templateByBoxName("InlineFile.template")
	t.Execute(w, struct {
		PackageName string
		InlineType []InlineType
		CompilerVersion string
		HasBytes bool
	}{
		PackageName: file.PackageName,
		InlineType: inlines,
		CompilerVersion: compilerVersion,
		HasBytes: hasBytes,
	})
}

func addInlineFromImports(file *pbparser.ProtoFile, dependencyData map[string]dependencyData) {
	for _, dep := range dependencyData {
		importedFile, err := parseImportedFile(dep.path)
		if err != nil {
			fmt.Println("ERROR:", "imported file cannot be parsed:", dep.path, "\n", err)
			os.Exit(1)
		}
		for _, option := range importedFile.Options {
			if option.Name == "inline" && option.Value == "true" {
				for _, m := range importedFile.Messages {
					if len(m.Options) == 1 && m.Options[0].Name == "inline_type" {
						if importedFile.PackageName != file.PackageName {
							globalInlineTypes[importedFile.PackageName + "." + m.Name] = m.Options[0].Value
						} else {
							globalInlineTypes[m.Name] = m.Options[0].Value
						}
					}
				}
			}
		}
	}
}

func isInlineFileByPath(path string) bool {
	file, err := parseImportedFile(path)
	if err != nil {
		fmt.Println("ERROR:", "imported file cannot be parsed:", path, "\n", err)
		os.Exit(1)
	}
	return isInlineFile(&file)
}
