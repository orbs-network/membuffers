// Copyright 2018 the membuffers authors
// This file is part of the membuffers library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package main

import (
	"fmt"
	"github.com/orbs-network/membuffers/go/membufc/api"
	"os"
	"strings"
)

func isFlag(arg string) bool {
	return strings.HasPrefix(arg, "-")
}

func handleFlag(flag string, conf *api.Config) {
	switch flag {
	case "--version":
		displayVersion(conf)
	case "-v":
		displayVersion(conf)
	case "--help":
		displayUsage()
	case "-h":
		displayUsage()
	case "--go":
		conf.Language = "go"
	case "-g":
		conf.Language = "go"
	case "-m":
		conf.Mock = true
	case "--mock":
		conf.Mock = true
	case "--go-ctx":
		conf.LanguageGoCtx = true
	default:
		fmt.Println("ERROR: Unknown command line flag:", flag)
		displayUsage()
	}
}

func displayVersion(conf *api.Config) {
	fmt.Println("membufc " + conf.Version)
	os.Exit(0)
}

func displayUsage() {
	fmt.Println("Usage: membufc [OPTION] PROTO_FILES")
	fmt.Println("Parse PROTO_FILES and generate output based on the options given:")
	fmt.Println("  -v, --version    Show version info and exit.")
	fmt.Println("  -h, --help       Show this usage text and exit.")
	fmt.Println("  -g, --go         Set output file language to Go.")
	fmt.Println("  -m, --mock       Generate mocks for services as well.")
	fmt.Println("  --go-ctx         Add context argument to all output Go service methods.")
	os.Exit(0)
}

func assertFileExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("ERROR: File does not exist:", path)
		os.Exit(1)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		displayUsage()
	}
	conf := api.NewConfig()
	for _, arg := range args {
		if isFlag(arg) {
			handleFlag(arg, conf)
		} else {
			assertFileExists(arg)
			conf.Files = append(conf.Files, arg)
		}
	}
	if err := api.Compile(conf); err != nil {
		fmt.Println("ERROR:", err.Error())
		os.Exit(1)
	}
}
