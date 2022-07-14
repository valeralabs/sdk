package main

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/util"
	"github.com/valeralabs/jenny/jen"
)

//go:generate git submodule update --init --recursive

func main() {
	swagger, err := util.LoadSwagger(input)

	if err != nil {
		panic(err)
	}

	file := jen.NewFile("api")

	for name, path := range swagger.Paths {
		// clean up the swagger object
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

		if strings.HasPrefix(name, "/rosetta/") == true {
			continue
		}

		for method, operation := range path.Operations() {
			if operation == nil {
				continue
			}

			opIdTypeName := cleanID(operation.OperationID)

			// parameters
			for _, parameter := range operation.Parameters {
				processParameter(file, parameter, opIdTypeName)
			}

			err = file.Render(ioutil.Discard)

			if err != nil {
				log.Fatalf("%s: failed to parse parameters: %v\n", opIdTypeName, err)
			}

			parameters := &jen.Statement{}

			// request body
			if operation.RequestBody != nil {
				for _, body := range operation.RequestBody.Value.Content {
					for title, prop := range body.Schema.Value.Properties {
						parameters.Add(processRequestBodyProperty(prop, opIdTypeName, title))
					}
				}
			}

			file.Add(parameters)

			err = file.Render(ioutil.Discard)

			if err != nil {
				log.Fatalf("%s: failed to parse request body: %v\n", opIdTypeName, err)
			}

			// responses
			for statusCode, response := range operation.Responses {
				processResponse(file, response, opIdTypeName, statusCode)
			}

			responseTypes := &jen.Statement{}

			for statusCode := range operation.Responses {
				if statusCode == "200" {
					responseTypes.Add(jen.ID(opIdTypeName + "Response"))
				}
			}

			responseTypes.Add(jen.Error())

			err = file.Render(ioutil.Discard)

			if err != nil {
				log.Fatalf("%s: failed to parse response types: %v\n", opIdTypeName, err)
			}

			switch method {
			case "GET":
				processGetReq(file, opIdTypeName, operation, name, responseTypes)

			case "POST":
				processPostReq(file, opIdTypeName, operation, parameters, name, responseTypes)

			default:
				log.Fatalf("Unsupported method %v", method)
			}

			err = file.Render(ioutil.Discard)

			if err != nil {
				log.Fatalf("%s: failed to parse method (%s): %v\n", opIdTypeName, method, err)
			}
		}
	}

	for _, code := range manuallyAddedTypes {
		file.Add(code)
	}

	err = file.Save(output)

	if err != nil {
		panic(err)
	}
}
