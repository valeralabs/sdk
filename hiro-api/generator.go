package main

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/azarc-io/json-schema-to-go-struct-generator/pkg/inputs"
	"github.com/pkg/errors"
)

func Convert(inputFiles []string, packageName string, outputFile string, debug bool) error {
	//ensure that files are aways processed in deterministic order
	sort.Strings(inputFiles)

	schemas, err := inputs.ReadInputFiles(inputFiles, false) // passing true will check for schema key in the file
	if err != nil {
		return errors.Wrapf(err, "error while reading input file")

	}
	generatorInstance := inputs.New(schemas...) // instance of generator which will produce structs
	err = generatorInstance.CreateTypes()
	if err != nil {
		return errors.Wrapf(err, "error while generating instance for producing structs")
	}

	packageDirectory := filepath.Dir(outputFile)
	err = os.MkdirAll(packageDirectory, 0755)
	if err != nil {
		return errors.Wrapf(err, "error while creating directory")
	}

	var w io.Writer
	w, err = os.Create(outputFile)
	if err != nil {
		return errors.Wrapf(err, "error while creating output file")
	}

	return inputs.Output(w, generatorInstance, packageName, inputFiles, debug)
}

// this is a script that reads all **/*.schema.json files in hiro-api and generates go structs

func main() {
	out, err := exec.Command("find", ".", "-name", "*.schema.json").Output()
	if err != nil {
		panic(err)
	}
	inputFiles := strings.Split(string(out), "\n")
	inputFiles = inputFiles[:len(inputFiles)-1]

	Convert(inputFiles, "main", "structs.go", true)
}
