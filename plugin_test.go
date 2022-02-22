// MIT License
//
// Copyright (c) 2022 - present Open Swoole Group
// Copyright (c) 2018 SpiralScout
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// When the environment variable RUN_AS_PROTOC_GEN_PHP is set, we skip running
// tests and instead act as protoc-gen-php. This allows the test binary to
// pass itself to protoc.
func init() {
	if os.Getenv("RUN_AS_PROTOC_GEN_PHP") != "" {
		main()
		os.Exit(0)
	}
}

func protoc(t *testing.T, args []string) {
	cmd := exec.Command("protoc", "--plugin=protoc-gen-php-grpc="+os.Args[0]) //nolint:gosec
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = append(os.Environ(), "RUN_AS_PROTOC_GEN_PHP=1")
	out, err := cmd.CombinedOutput()

	if len(out) > 0 || err != nil {
		t.Log("RUNNING: ", strings.Join(cmd.Args, " "))
	}

	if len(out) > 0 {
		t.Log(string(out))
	}

	if err != nil {
		t.Fatalf("protoc: %v", err)
	}
}

func Test_Simple(t *testing.T) {
	workdir, _ := os.Getwd()
	tmpdir, err := ioutil.TempDir("", "proto-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	args := []string{
		"-Itestdata",
		"--php-grpc_out=" + tmpdir,
		"simple/simple.proto",
	}
	protoc(t, args)

	assertEqualFiles(
		t,
		workdir+"/testdata/simple/TestSimple/SimpleServiceInterface.php",
		tmpdir+"/TestSimple/SimpleServiceInterface.php",
	)
}

func Test_PhpNamespaceOption(t *testing.T) {
	workdir, _ := os.Getwd()
	tmpdir, err := ioutil.TempDir("", "proto-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	args := []string{
		"-Itestdata",
		"--php-grpc_out=" + tmpdir,
		"php_namespace/service.proto",
	}
	protoc(t, args)

	assertEqualFiles(
		t,
		workdir+"/testdata/php_namespace/Test/CustomNamespace/ServiceInterface.php",
		tmpdir+"/Test/CustomNamespace/ServiceInterface.php",
	)
}

func Test_UseImportedMessage(t *testing.T) {
	workdir, _ := os.Getwd()
	tmpdir, err := ioutil.TempDir("", "proto-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	args := []string{
		"-Itestdata",
		"--php-grpc_out=" + tmpdir,
		"import/service.proto",
	}
	protoc(t, args)

	assertEqualFiles(
		t,
		workdir+"/testdata/import/Import/ServiceInterface.php",
		tmpdir+"/Import/ServiceInterface.php",
	)
}

func Test_PhpNamespaceOptionInUse(t *testing.T) {
	workdir, _ := os.Getwd()
	tmpdir, err := ioutil.TempDir("", "proto-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)
	args := []string{
		"-Itestdata",
		"--php-grpc_out=" + tmpdir,
		"import_custom/service.proto",
	}
	protoc(t, args)

	assertEqualFiles(
		t,
		workdir+"/testdata/import_custom/Test/CustomImport/ServiceInterface.php",
		tmpdir+"/Test/CustomImport/ServiceInterface.php",
	)
}

func Test_UseOfGoogleEmptyMessage(t *testing.T) {
	workdir, _ := os.Getwd()
	tmpdir, err := ioutil.TempDir("", "proto-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)
	args := []string{
		"-Itestdata",
		"--php-grpc_out=" + tmpdir,
		"use_empty/service.proto",
	}
	protoc(t, args)

	assertEqualFiles(
		t,
		workdir+"/testdata/use_empty/Test/ServiceInterface.php",
		tmpdir+"/Test/ServiceInterface.php",
	)
}

func assertEqualFiles(t *testing.T, original, generated string) {
	assert.FileExists(t, generated)

	originalData, err := ioutil.ReadFile(original)
	if err != nil {
		t.Fatal("Can't find original file for comparison")
	}

	generatedData, err := ioutil.ReadFile(generated)
	if err != nil {
		t.Fatal("Can't find generated file for comparison")
	}

	// every OS has a special boy
	r := strings.NewReplacer("\r\n", "", "\n", "")
	assert.Equal(t, r.Replace(string(originalData)), r.Replace(string(generatedData)))
}
