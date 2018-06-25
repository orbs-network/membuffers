package main

import (
	"os"
	"strings"
	"fmt"
	"github.com/tallstoat/pbparser"
)

type config struct {
	language string
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
	default:
		fmt.Println("ERROR: Unknown command line flag:", flag)
		displayUsage()
	}
}

func displayVersion() {
	fmt.Println("membufc 0.0.1")
	os.Exit(0)
}

func displayUsage() {
	fmt.Println("Usage: membufc [OPTION] PROTO_FILES")
	fmt.Println("Parse PROTO_FILES and generate output based on the options given:")
	fmt.Println("  -v, --version    Show version info and exit.")
	fmt.Println("  -h, --help       Show this usage text and exit.")
	fmt.Println("  -g, --go         Set output file language to Go.")
	os.Exit(0)
}

func assertFileExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("ERROR: File does not exist:", path)
		os.Exit(1)
	}
}

func outputFileForPath(path string) string {
	parts := strings.Split(path, ".")
	return strings.Join(parts[0:len(parts)-1], ".") + ".mb.go"
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
		protoFile, err := pbparser.ParseFile(path)
		if err != nil {
			fmt.Println("ERROR:", err.Error())
			os.Exit(1)
		}
		outPath := outputFileForPath(path)
		f, err := os.Create(outPath)
		if err != nil {
			fmt.Println("ERROR:", err.Error())
			os.Exit(1)
		}
		defer f.Close()
		compileProtoFile(f, protoFile)
		fmt.Println("Created file:", outPath)
	}
}