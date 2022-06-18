package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/codegen"
	"github.com/deepmap/oapi-codegen/pkg/util"
	"github.com/getkin/kin-openapi/openapi3"
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
		// clean Parameters
		if path.Parameters != nil {
			// type Parameters []*ParameterRef
			for _, parameter := range path.Parameters {
				// parameter.Value.Name = clean(parameter.Value.Name)
				// parameter.Value.Description = clean(parameter.Value.Description)
				// log parameter as json
				fmt.Printf("%s\n", clean(parameter.Value.Name))

			}
		}

		// clean OperationIDs
		for _, op := range nilOps(path) {
			op.OperationID = clean(op.OperationID)
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

// returns an array of non-nil *openapi3.Operation objects in a given *openapi3.PathItem
func nilOps (itm *openapi3.PathItem) []*openapi3.Operation {
	var ops []*openapi3.Operation
	if itm.Connect != nil {
		ops = append(ops, itm.Connect)
	}
	if itm.Delete != nil {
		ops = append(ops, itm.Delete)
	}
	if itm.Get != nil {
		ops = append(ops, itm.Get)
	}
	if itm.Head != nil {
		ops = append(ops, itm.Head)
	}
	if itm.Options != nil {
		ops = append(ops, itm.Options)
	}
	if itm.Patch != nil {
		ops = append(ops, itm.Patch)
	}
	if itm.Post != nil {
		ops = append(ops, itm.Post)
	}
	if itm.Put != nil {
		ops = append(ops, itm.Put)
	}
	if itm.Trace != nil {
		ops = append(ops, itm.Trace)
	}
	return ops
}