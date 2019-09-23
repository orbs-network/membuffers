// Copyright 2018 the membuffers authors
// This file is part of the membuffers library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"path"
	"strings"
	//	"strings"
)

type dependencyData struct {
	relative string
	path     string
}

type importProvider struct {
	protoFile        string
	moduleToRelative map[string]dependencyData
}

func (i *importProvider) Provide(module string) (io.Reader, error) {
	if r := i.satisfies(module); r != nil {
		return r, nil
	}
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

func (i *importProvider) satisfies(moduleName string) io.Reader {
	if moduleName == "membuffers" {
		return strings.NewReader(`
syntax = "proto3";
package membuffers;
message bytes32 {}
`)
	}

	return nil
}
