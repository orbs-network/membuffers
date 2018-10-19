package main

import (
	"github.com/orbs-network/pbparser"
	"io"
	"path"
)

func compileMockFile(w io.Writer, file pbparser.ProtoFile, dependencyData map[string]dependencyData, compilerVersion string, languageGoCtx bool) {
	addMockHeader(w, &file, dependencyData, compilerVersion, languageGoCtx)
	for _, s := range file.Services {
		addMockService(w, file.PackageName, s, &file, languageGoCtx)
	}
}

func addMockHeader(w io.Writer, file *pbparser.ProtoFile, dependencyData map[string]dependencyData, compilerVersion string, languageGoCtx bool) {
	var goPackage string
	implementHandlers := []NameWithAndWithoutImport{}
	registerHandlers := []NameWithAndWithoutImport{}
	for _, option := range file.Options {
		if option.Name == "go_package" {
			goPackage = option.Value
		}
	}
	for _, service := range file.Services {
		for _, option := range service.Options {
			if option.Name == "register_handler" {
				registerHandlers = append(registerHandlers, getNameWithAndWithoutImport(option.Value))
			}
			if option.Name == "implement_handler" {
				implementHandlers = append(implementHandlers, getNameWithAndWithoutImport(option.Value))
			}
		}
	}
	imports := []string{}
	for _, dependency := range file.Dependencies {
		if !doesFileContainHandlers(dependencyData[dependency].path, append(implementHandlers, registerHandlers...)) {
			continue
		}
		relative := dependencyData[dependency].relative
		packageImport := path.Dir(path.Clean(goPackage + "/" + relative + "/" + dependency))
		if packageImport != goPackage {
			imports = append(imports, packageImport)
		}
	}
	t := templateByBoxName("MockFileHeader.template")
	t.Execute(w, struct {
		PackageName       string
		Imports           []string
		HasServiceMethods bool
		CompilerVersion   string
		LanguageGoCtx     bool
	}{
		PackageName:       file.PackageName,
		Imports:           unique(imports),
		HasServiceMethods: fileHasServiceMethods(file),
		CompilerVersion:   compilerVersion,
		LanguageGoCtx:     languageGoCtx,
	})
}

func addMockService(w io.Writer, packageName string, s pbparser.ServiceElement, file *pbparser.ProtoFile, languageGoCtx bool) {
	methods := []ServiceMethod{}
	for _, rpc := range s.RPCs {
		method := ServiceMethod{
			Name:   rpc.Name,
			Input:  removeLocalPackagePrefix(packageName, rpc.RequestType.Name()),
			Output: removeLocalPackagePrefix(packageName, rpc.ResponseType.Name()),
		}
		methods = append(methods, method)
	}
	registerHandlers := []NameWithAndWithoutImport{}
	implementHandlers := []NameWithAndWithoutImport{}
	for _, option := range s.Options {
		if option.Name == "register_handler" {
			registerHandlers = append(registerHandlers, getNameWithAndWithoutImport(option.Value))
		}
		if option.Name == "implement_handler" {
			implementHandlers = append(implementHandlers, getNameWithAndWithoutImport(option.Value))
		}
	}
	t := templateByBoxName("MockService.template")
	t.Execute(w, struct {
		ServiceName       string
		Methods           []ServiceMethod
		RegisterHandlers  []NameWithAndWithoutImport
		ImplementHandlers []NameWithAndWithoutImport
		LanguageGoCtx     bool
	}{
		ServiceName:       s.Name,
		Methods:           methods,
		RegisterHandlers:  registerHandlers,
		ImplementHandlers: implementHandlers,
		LanguageGoCtx:     languageGoCtx,
	})
}
