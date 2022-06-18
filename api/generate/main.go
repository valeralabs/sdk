package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/codegen"
	"github.com/deepmap/oapi-codegen/pkg/util"
)

//go:generate git submodule update --init --recursive

var (
	output = "../api.gen.go"
	input  = "./source/docs/openapi.yaml"
)

var URLArgumentScope = regexp.MustCompile(`{([^}]*)}`)
var replacements = []string{"nft", "btc", "stx", "api", "id", "tx", "ft", "tld"}

func clean(source string) string {
	for _, cursor := range replacements {
		source = strings.ReplaceAll(source, cursor, strings.ToUpper(cursor))
	}

	return source
}

func main() {
	options := codegen.Configuration{
		PackageName: "api",
		Generate: codegen.GenerateOptions{
			ChiServer:  false,
			EchoServer: false,
			GinServer:  false,
			// API client functions
			Client: false,
			// API client types
			Models:       true,
			EmbeddedSpec: false,
		},
	}

	options = options.UpdateDefaults()

	swagger, err := util.LoadSwagger(input)

	if err != nil {
		panic(err)
	}

	for name, path := range swagger.Paths {
		delete(swagger.Paths, name)

		if path.Parameters != nil {
			for _, parameter := range path.Parameters {
				parameter.Value.Name = clean(parameter.Value.Name)
			}
		}

		for _, match := range URLArgumentScope.FindAllStringSubmatch(name, -1) {
			modified := clean(match[0])

			if modified != match[0] {
				name = strings.ReplaceAll(name, match[0], modified)
			}
		}

		for _, operation := range path.Operations() {
			if operation != nil {
				for _, parameter := range operation.Parameters {
					parameter.Value.Name = clean(parameter.Value.Name)
				}

				operation.OperationID = clean(operation.OperationID)
			}
		}

		swagger.Paths[name] = path
	}

	code, err := codegen.Generate(swagger, options)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(output, []byte(code), 0644)

	if err != nil {
		panic(err)
	}
}
