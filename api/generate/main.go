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
			for method, operation := range path.Operations() {
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

					var bodyParams Code

					// request body
					if operation.RequestBody != nil {
						for _, body := range operation.RequestBody.Value.Content {
							for title, prop := range body.Schema.Value.Properties {
								wg.Add(1)

								title := title
								prop := prop

								go func() {
									defer wg.Done()
									processRequestBodyProperty(prop, &bodyParams, opIdTypeName, title, f)
									fmt.Printf(typePrefix+"Body property `%v` processed\n", title)
								}()
							}
						}
					}

					// responses
					for statusCode, response := range operation.Responses {
						wg.Add(1)

						response := response
						statusCode := statusCode

						go func() {
							defer wg.Done()
							processResponse(response, opIdTypeName, statusCode, f)
							fmt.Printf(typePrefix+"Response %v processed\n", statusCode)
						}()
					}

					// ---- FUNCTIONS

					// funcPrefix := prefix + string(colourCodes["gray"]) + "[FUNC] " + string(colourCodes["reset"])

					possibleResponseTypes := []jen.Code{}
					for statusCode := range operation.Responses {
						if statusCode == "200" {
							possibleResponseTypes = append(possibleResponseTypes, jen.ID(opIdTypeName+"Response"))
						}
						//else {
						// 	possibleResponseTypes = append(possibleResponseTypes, jen.ID(opIdTypeName + statusCode + "Error"))
						// }
					}

					possibleResponseTypes = append(possibleResponseTypes, jen.Error())

					switch method {
					case "GET":
						// GET
					case "POST":
						processPostReq(f, opIdTypeName, operation, bodyParams, name, possibleResponseTypes)
					default:
						panic(fmt.Sprintf("Unsupported method %v", method))
					}

					wg.Wait()

					if len(params.Generated) > 0 {
						f.Commentf("Defines parameters for %v", opIdTypeName)
						f.Type().ID(opIdTypeName + "Params").Structure(params.Generated...)
					}
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
		fmt.Println(err)
		panic(err)
	}

	fmt.Printf(string(colourCodes["green"])+"Finished processing in %v\n"+string(colourCodes["reset"]), time.Since(processingStart))
}
