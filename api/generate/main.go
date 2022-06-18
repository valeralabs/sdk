package main

import (
	"io/ioutil"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/codegen"
	"github.com/deepmap/oapi-codegen/pkg/util"
)

//go:generate git submodule update --init --recursive

var (
	output = "../api.gen.go"
	input  = "./source/docs/openapi.yaml"
)

var replacements = []string{"nft", "btc", "stx", "api", "id", "tx", "ft"}

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
			ChiServer:    false,
			EchoServer:   false,
			GinServer:    false,
			Client:       false,
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
		if path.Connect != nil {
			path.Connect.OperationID = clean(path.Connect.OperationID)
		}

		if path.Delete != nil {
			path.Delete.OperationID = clean(path.Delete.OperationID)
		}

		if path.Get != nil {
			path.Get.OperationID = clean(path.Get.OperationID)
		}

		if path.Head != nil {
			path.Head.OperationID = clean(path.Head.OperationID)
		}

		if path.Options != nil {
			path.Options.OperationID = clean(path.Options.OperationID)
		}

		if path.Patch != nil {
			path.Patch.OperationID = clean(path.Patch.OperationID)
		}

		if path.Post != nil {
			path.Post.OperationID = clean(path.Post.OperationID)
		}

		if path.Put != nil {
			path.Put.OperationID = clean(path.Put.OperationID)
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
