// Copyright 2018 the membuffers authors
// This file is part of the membuffers library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package tests

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"os/exec"
	"strings"
	"testing"
)

func readInJs(t testing.TB, raw []byte, code string) map[string]interface{} {
	var hexBuf []string
	for _, b := range raw {
		hexBuf = append(hexBuf, "0x"+hex.EncodeToString([]byte{b}))
	}
	hexString := "[" + strings.Join(hexBuf, ",") + "]"
	js := fmt.Sprintf(`
const {InternalMessage, FieldTypes} = require("../javascript/dist/membuffers");

const buf = new Uint8Array(%s);
%s
console.log(JSON.stringify(obj));
`, hexString, code)

	out, err := exec.Command("node", "-e", js).Output()
	if ee, ok := err.(*exec.ExitError); ok {
		t.Log(ee.String())
		t.Log(string(ee.Stderr))
	}
	require.NoError(t, err)
	t.Log(string(out))

	j := make(map[string]interface{})
	require.NoError(t, json.Unmarshal(out, &j))
	return j
}
