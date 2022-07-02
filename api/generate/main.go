package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/util"
	"github.com/valeralabs/jenny/jen"
)

//go:generate git submodule update --init --recursive

type Code struct {
	Generated []jen.Code
}

func (p *Code) add(code ...jen.Code) {
	p.Generated = append(p.Generated, code...)
}

func main() {
	fmt.Println("Loading OpenAPI 3.0 spec...")
	rootStart := time.Now()
	swagger, err := util.LoadSwagger(input)

	if err != nil {
		panic(err)
	}
	
	fmt.Printf("Loaded API spec in %v\n", time.Since(rootStart))
	processingStart := time.Now()

	// type generation
	f := jen.NewFile("api")

	for name, path := range swagger.Paths {
		delete(swagger.Paths, name)

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

		if !strings.HasPrefix(name, "/rosetta/") {
			for _, operation := range path.Operations() {
				if operation != nil {
					startTime := time.Now()
					prefix := "➤ │ "
					opIdTypeName := cleanID(operation.OperationID)
					fmt.Printf("➤ ┌ "+string(colourCodes["blue"])+"%s\n"+string(colourCodes["reset"]), opIdTypeName)

					var params Code

					var wg sync.WaitGroup

					// ---- TYPES

					typePrefix := prefix + string(colourCodes["gray"]) + "[TYPE] " + string(colourCodes["reset"])

					// parameters
					for _, parameter := range operation.Parameters {
						wg.Add(1)
						parameter := parameter

						go func() {
							defer wg.Done()
							processParameter(parameter, &params, opIdTypeName, f)
							fmt.Printf(typePrefix+"Parameter `%v` processed\n", parameter.Value.Name)
							if !parameter.Value.Required {
								fmt.Printf(prefix + "       └ " + string(colourCodes["yellow"]) + "Optional\n" + string(colourCodes["reset"]))
							}
						}()
					}

					// request body
					if operation.RequestBody != nil {
						for _, body := range operation.RequestBody.Value.Content {
							for title, prop := range body.Schema.Value.Properties {
								wg.Add(1)

								title := title
								prop := prop

								go func() {
									defer wg.Done()
									processRequestBodyProperty(prop, &params, opIdTypeName, title, f)
									fmt.Printf(typePrefix+"Body property `%v` processed\n", title)
								}()
							}
						}
					}

					// responses
					// responses := make(map[string]Code)
					// for statusCode, response := range operation.Responses {
					// 	wg.Add(1)

					// 	response := response
					// 	statusCode := statusCode

					// 	go func() {
					// 		defer wg.Done()
					// 		var queue Code
					// 		processResponse(response, opIdTypeName, statusCode, &queue)
					// 		responses[statusCode] = queue
					// 		fmt.Printf(typePrefix+"Response %v processed\n", statusCode)
					// 	}()
					// }

					// ---- FUNCTIONS

					// funcPrefix := prefix + string(colourCodes["gray"]) + "[FUNC] " + string(colourCodes["reset"])

					wg.Wait()

					ParamsTypeName := opIdTypeName + "Params"

					f.Commentf("Defines parameters for %v", opIdTypeName)
					f.Type().ID(ParamsTypeName).Structure(params.Generated...)

					// // loop through responses and add them to file
					// for statusCode, response := range responses {
					// 	f.Commentf("Defines a %v response for %v.", statusCode, opIdTypeName)
					// 	f.Type().ID(opIdTypeName + statusCode + "Response").Structure(response.Generated...)
					// }

					fmt.Printf("➤ └ Completed in %v\n", time.Since(startTime))
				}
			}
		}
	}

	for _, code := range manuallyAddedTypes {
		f.Add(code)
	}

	fmt.Printf("➤ Rendering output to %v\n", output)
	err = f.Save(output)

	if err != nil {
		panic(err)
	}

	fmt.Printf(string(colourCodes["green"])+"Finished processing in %v\n"+string(colourCodes["reset"]), time.Since(processingStart))
}
