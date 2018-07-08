package main

import (
	"fmt"
	"os"
	"path"
	"github.com/orbs-network/pbparser"
	"io"
)

type Enum struct{
	Name string
	Values []EnumValue
}

type EnumValue struct{
	Name string
	Value int
}

func addEnumsFromImports(file *pbparser.ProtoFile, dependencyData map[string]dependencyData) {
	for _, dep := range dependencyData {
		importedFile, err := parseImportedFile(dep.path)
		if err != nil {
			fmt.Println("ERROR:", "imported file cannot be parsed:", dep.path, "\n", err)
			os.Exit(1)
		}
		for i, enum := range importedFile.Enums {
			importedFile.Enums[i].Documentation = "imported"
			importedPackageName := path.Base(path.Dir(dep.path))
			if importedPackageName != file.PackageName {
				importedFile.Enums[i].Name = importedPackageName + "." + enum.Name
			} else {
				importedFile.Enums[i].Name = enum.Name
			}

		}
		file.Enums = append(file.Enums, importedFile.Enums...)
	}
}

func addEnums(w io.Writer, enums []pbparser.EnumElement) {
	if len(enums) == 0 {
		return
	}
	messageEnums, _ := getFileEnums(enums)
	t := templateByBoxName("MessageFileEnums.template")
	t.Execute(w, struct {
		Enums []Enum
	}{
		Enums: messageEnums,
	})
}

func getFileEnums(enums []pbparser.EnumElement) ([]Enum, map[string]int) {
	enumByIndex := []Enum{}
	enumNameToIndex := make(map[string]int)
	for _, enum := range enums {
		enumNameToIndex[enum.Name] = len(enumByIndex)
		values := []EnumValue{}
		for _, value := range enum.EnumConstants {
			values = append(values, EnumValue{
				Name: value.Name,
				Value: value.Tag,
			})
		}
		// only add here enums from this package
		if enum.Documentation != "imported" {
			enumByIndex = append(enumByIndex, Enum{
				Name: enum.Name,
				Values: values,
			})
		}
	}
	return enumByIndex, enumNameToIndex
}
