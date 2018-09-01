package main

import (
	"os"
	"strings"
	"fmt"
	"github.com/orbs-network/pbparser"
	"io"
	"errors"
	"path"
)

const MEMBUFC_VERSION = "0.0.19"

type config struct {
	language string
	mock bool
	files []string
}

var conf = config{}

func isFlag(arg string) bool {
	return strings.HasPrefix(arg, "-")
}

func handleFlag(flag string) {
	switch flag {
	case "--version":
		displayVersion()
	case "-v":
		displayVersion()
	case "--help":
		displayUsage()
	case "-h":
		displayUsage()
	case "--go":
		conf.language = "go"
	case "-g":
		conf.language = "go"
	case "-m":
		conf.mock = true
	case "--mock":
		conf.mock = true
	default:
		fmt.Println("ERROR: Unknown command line flag:", flag)
		displayUsage()
	}
}

func displayVersion() {
	fmt.Println("membufc " + MEMBUFC_VERSION)
	os.Exit(0)
}

func displayUsage() {
	fmt.Println("Usage: membufc [OPTION] PROTO_FILES")
	fmt.Println("Parse PROTO_FILES and generate output based on the options given:")
	fmt.Println("  -v, --version    Show version info and exit.")
	fmt.Println("  -h, --help       Show this usage text and exit.")
	fmt.Println("  -g, --go         Set output file language to Go.")
	fmt.Println("  -m, --mock       Generate mocks for services as well.")
	os.Exit(0)
}

func assertFileExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("ERROR: File does not exist:", path)
		os.Exit(1)
	}
}

func outputFileForPath(path string, suffix string) string {
	parts := strings.Split(path, ".")
	return strings.Join(parts[0:len(parts)-1], ".") + suffix
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		displayUsage()
	}
	for _, arg := range args {
		if isFlag(arg) {
			handleFlag(arg)
		} else {
			assertFileExists(arg)
			conf.files = append(conf.files, arg)
		}
	}
	for _, path := range conf.files {
		fmt.Println("Compiling file:\t", path)
		in, err := os.Open(path)
		if err != nil {
			fmt.Println("ERROR:", err.Error())
			os.Exit(1)
		}
		p := importProvider{protoFile: path, moduleToRelative: make(map[string]dependencyData)}
		protoFile, err := pbparser.Parse(in, &p)
		if err != nil {
			fmt.Println("ERROR:", err.Error())
			os.Exit(1)
		}
		outPath := outputFileForPath(path, ".mb.go")
		out, err := os.Create(outPath)
		if err != nil {
			fmt.Println("ERROR:", err.Error())
			os.Exit(1)
		}
		defer out.Close()
		if isInlineFile(&protoFile) {
			compileInlineFile(out, protoFile, p.moduleToRelative, MEMBUFC_VERSION)
		} else {
			compileProtoFile(out, protoFile, p.moduleToRelative, MEMBUFC_VERSION)
		}
		fmt.Println("Created file:\t", outPath)
		if len(protoFile.Services) > 0 && conf.mock {
			outPath := outputFileForPath(path, "_mock.mb.go")
			out, err := os.Create(outPath)
			if err != nil {
				fmt.Println("ERROR:", err.Error())
				os.Exit(1)
			}
			defer out.Close()
			compileMockFile(out, protoFile, p.moduleToRelative, MEMBUFC_VERSION)
			fmt.Println("Created mock file:\t", outPath)
		}
		fmt.Println()
	}
}

type dependencyData struct {
	relative string
	path string
}
type importProvider struct {
	protoFile string
	moduleToRelative map[string]dependencyData
}
func (i *importProvider) Provide(module string) (io.Reader, error) {
	basePath := path.Dir(i.protoFile) + "/"
	relativePath := ""
	attempts := []string{}
	for nesting := 0; nesting < 5; nesting++ {
		attemptPath := basePath + relativePath + module
		f, err := os.Open(attemptPath)
		if err == nil {
			if i.moduleToRelative != nil {
				i.moduleToRelative[module] = dependencyData{relative: relativePath, path: attemptPath}
			}
			return f, nil
		}
		attempts = append(attempts, attemptPath)
		relativePath = "../" + relativePath
	}
	return nil, errors.New(fmt.Sprintf("import %s not found, looked at %v", module, attempts))
}

func isInlineFile(file *pbparser.ProtoFile) bool {
	for _, option := range file.Options {
		if option.Name == "inline" && option.Value == "true" {
			return true
		}
	}
	return false
}