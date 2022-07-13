package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/valeralabs/jenny/jen"
)

func processParameter(parameter *openapi3.ParameterRef, params *Code, opID string, f *jen.File, tls *sync.Map) {
	val := parameter.Value
	schema := val.Schema.Value
	var queue Code

	queue.add(jen.Comment(cleanDesc(val.Description)))

	if !val.Required {
		defaultStr := ""
		if schema.Default != nil {
			defaultStr = fmt.Sprintf("The default value is %v.", schema.Default)
		}
		queue.add(jen.Commentf("Optional. " + defaultStr))
		if schema.Max != nil {
			queue.add(jen.Comment(fmt.Sprintf("Max value is %v.", *schema.Max)))
		}
	}

	if schema.Type == "array" || schema.Type == "object" {
		id, _ := processObjectOrArray(schema, tls, f, opID+cleanID(val.Name))
		queue.add(jen.ID(cleanID(val.Name)).ID(id))
	} else {
		queue.add(
			jen.ID(cleanID(val.Name)).ID(typeReplace(val.Schema.Value.Type)),
		)
	}

	params.add(queue.Generated...)
}

func processRequestBodyProperty(prop *openapi3.SchemaRef, params *Code, opID string, title string, f *jen.File, tls *sync.Map) {
	schema := prop.Value
	titlizedTitle := strings.Title(title)

	// check if it's present in manually added types
	if _, ok := manuallyAddedTypes[opID+titlizedTitle]; ok {
		params.add(
			jen.Comment(cleanDesc(schema.Description)),
			jen.ID(cleanID(title)).ID(opID+titlizedTitle),
		)
	} else {
		if schema.Type == "array" || schema.Type == "object" {
			id, _ := processObjectOrArray(schema, tls, f, opID+cleanID(titlizedTitle))
			params.add(
				jen.Comment(cleanDesc(schema.Description)),
				jen.ID(cleanID(title)).ID(id),
			)
		} else {
			params.add(
				jen.Comment(cleanDesc(schema.Description)),
				jen.ID(cleanID(title)).ID(typeReplace(schema.Type)),
			)
		}
	}
}

func processResponse(response *openapi3.ResponseRef, opID string, statusCode string, f *jen.File, addedSchemas *sync.Map) {
	val := response.Value

	for _, content := range val.Content {
		if content.Schema != nil {
			schema := content.Schema.Value

			var backupTitle string

			if statusCode == "200" {
				backupTitle = opID + "Response"
			} else {
				backupTitle = opID + "Error" + statusCode
			}

			var respType string

			if schema.Type == "array" || schema.Type == "object" {
				respType, _ = processObjectOrArray(schema, addedSchemas, f, backupTitle)
			} else {
				respType = typeReplace(schema.Type)
			}
			descMsg := ""
			if val.Description != nil {
				descMsg = fmt.Sprintf(" (%v)", *val.Description)
			}

			if respType == "" {
				respType = "any"
			} 

			f.Add(
				jen.Commentf("Defines a %v%v response for %v.", statusCode, descMsg, opID),
				jen.Line(),
				jen.Type().ID(backupTitle).ID(respType),
			)
		}
	}
}

func processPostReq(f *jen.File, opIdTypeName string, operation *openapi3.Operation, bodyParams Code, name string, possibleResponseTypes []jen.Code) {
	// POST
	f.Type().ID(opIdTypeName + "Body").Structure(bodyParams.Generated...)

	var inputParams Code

	inputParams.add(jen.ID("server").ID("Server"))

	// if there are parameters, add them to the input params
	if len(operation.Parameters) > 0 {
		inputParams.add(jen.ID("params").ID(opIdTypeName + "Params"))
	}
	// if there is a request body, add it to the input params
	if operation.RequestBody != nil {
		inputParams.add(jen.ID("body").ID(opIdTypeName + "Body"))
	}

	var funcCode Code

	if len(operation.Parameters) > 0 {
		funcCode.add(
			// url := fmt.Sprintf("%s%s", server, fillPath(name, params))
			jen.ID("url").Op(":=").Qualified("fmt", "Sprintf").Call(
				jen.Literal("%s%s"),
				jen.ID("server"),
				jen.ID("fillPath").Call(
					jen.Literal(name),
					jen.ID("params"),
				),
			),
		)
	} else {
		funcCode.add(
			// url := fmt.Sprintf("%s%s", server, name))
			jen.ID("url").Op(":=").Qualified("fmt", "Sprintf").Call(
				jen.Literal("%s%s"),
				jen.ID("server"),
				jen.Literal(name),
			),
		)
	}

	funcCode.add(
		// var returnedErr error
		jen.Var().ID("returnedErr").Error(),
		// sendBody, _ := json.Marshal(body)
		jen.ID("sendBody").Op(",").ID("_").Op(":=").Qualified("encoding/json", "Marshal").Call(jen.ID("body")),
		// resp, status := makePostReq(url, string(sendBody), &returnedErr)
		jen.ID("resp").Op(",").ID("status").Op(":=").ID("makePostReq").Call(jen.ID("url"), jen.ID("string").Call(jen.ID("sendBody")), jen.ID("&returnedErr")),
		jen.Line(),
		// if returnedErr != nil {
		jen.If(jen.ID("returnedErr").Op("!=").Nil()).Block(
			// return PostFeeTransactionResponse{}, returnedErr
			jen.Return().ID(opIdTypeName+"Response").Block().Op(",").ID("returnedErr"),
		),
		jen.Line(),
		// switch status {
		jen.Switch(jen.ID("status")).Block(
			// case 200:
			jen.Case(jen.Literal(200)).Block(
				// var respObj PostFeeTransactionResponse
				jen.Var().ID("respObj").ID(opIdTypeName+"Response"),
				// err := json.Unmarshal(resp, &respObj)
				jen.ID("err").Op(":=").Qualified("encoding/json", "Unmarshal").Call(jen.ID("resp"), jen.ID("&respObj")),
				// if err != nil {
				// 	return respObj, err
				// }
				jen.If(jen.ID("err").Op("!=").Nil()).Block(
					jen.Return().ID("respObj").Op(",").ID("err"),
				),
				// return respObj, nil
				jen.Return().ID("respObj").Op(",").Nil(),
			),
			jen.Default().Block(
				// var respObj struct{
				// 	Error string `json:"error"`
				// }
				jen.Var().ID("respObj").Structure(
					jen.ID("Error").String().Tag(map[string]string{"json": "error"}),
				),
				// err := json.Unmarshal(resp, &respObj)
				jen.ID("err").Op(":=").Qualified("encoding/json", "Unmarshal").Call(jen.ID("resp"), jen.ID("&respObj")),
				// if err != nil {
				// 	return PostFeeTransactionResponse{}, err
				// }
				jen.If(jen.ID("err").Op("!=").Nil()).Block(
					jen.Return().ID(opIdTypeName+"Response").Block().Op(",").ID("err"),
				),
				// return PostFeeTransactionResponse{}, errors.New(respObj.Error)
				jen.Return().ID(opIdTypeName+"Response").Block().Op(",").Qualified("errors", "New").Call(jen.ID("respObj").Dot("Error")),
			),
		),
	)

	if operation.Description != "" {
		// // detect if there are references (like `transaction_payload`) and replace them with the actual values
		// // pull each reference by backticks and replace with cleanID(reference)
		// splitStr := strings.Split(operation.Description, "`")
		// for i := 0; i < len(splitStr); i++ {
		// 	if i % 2 == 0 {
		// 		// make sure it's not surrounded by no none-whitespace characters
		// 		if strings.TrimSpace(splitStr[i]) == "" {
		// 			splitStr[i] = cleanID(splitStr[i])
		// 		}
		// 	}
		// }

		f.Comment(operation.Description)
	}
	f.Func().ID(opIdTypeName).Parameters(
		inputParams.Generated...,
	).Parameters(possibleResponseTypes...).Block(
		funcCode.Generated...,
	).Line()
}

func processGetReq(f *jen.File, opIdTypeName string, operation *openapi3.Operation, name string, possibleResponseTypes []jen.Code) {
	// GET

	var inputParams Code

	inputParams.add(jen.ID("server").ID("Server"))

	if opIdTypeName == "GetBlockByHeight" {
		for _, param := range operation.Parameters {
			fmt.Println(param)
		}
	}

	// if there are parameters, add them to the input params
	if len(operation.Parameters) > 0 {
		inputParams.add(jen.ID("params").ID(opIdTypeName + "Params"))
	}

	var funcCode Code

	if len(operation.Parameters) > 0 {
		funcCode.add(
			// url := fmt.Sprintf("%s%s", server, fillPath(name, params))
			jen.ID("url").Op(":=").Qualified("fmt", "Sprintf").Call(
				jen.Literal("%s%s"),
				jen.ID("server"),
				jen.ID("fillPath").Call(
					jen.Literal(name),
					jen.ID("params"),
				),
			),
		)
	} else {
		funcCode.add(
			// url := fmt.Sprintf("%s%s", server, name))
			jen.ID("url").Op(":=").Qualified("fmt", "Sprintf").Call(
				jen.Literal("%s%s"),
				jen.ID("server"),
				jen.Literal(name),
			),
		)
	}

	funcCode.add(
		// var returnedErr error
		jen.Var().ID("returnedErr").Error(),
		// resp, status := makePostReq(url, &returnedErr)
		jen.ID("resp").Op(",").ID("status").Op(":=").ID("makeGetReq").Call(jen.ID("url"), jen.ID("&returnedErr")),
		jen.Line(),
		// if returnedErr != nil {
		jen.If(jen.ID("returnedErr").Op("!=").Nil()).Block(
			// return PostFeeTransactionResponse{}, returnedErr
			jen.Return().ID(opIdTypeName+"Response").Block().Op(",").ID("returnedErr"),
		),
		jen.Line(),
		// switch status {
		jen.Switch(jen.ID("status")).Block(
			// case 200:
			jen.Case(jen.Literal(200)).Block(
				// var respObj PostFeeTransactionResponse
				jen.Var().ID("respObj").ID(opIdTypeName+"Response"),
				// err := json.Unmarshal(resp, &respObj)
				jen.ID("err").Op(":=").Qualified("encoding/json", "Unmarshal").Call(jen.ID("resp"), jen.ID("&respObj")),
				// if err != nil {
				// 	return respObj, err
				// }
				jen.If(jen.ID("err").Op("!=").Nil()).Block(
					jen.Return().ID("respObj").Op(",").ID("err"),
				),
				// return respObj, nil
				jen.Return().ID("respObj").Op(",").Nil(),
			),
			jen.Default().Block(
				// var respObj struct{
				// 	Error string `json:"error"`
				// }
				jen.Var().ID("respObj").Structure(
					jen.ID("Error").String().Tag(map[string]string{"json": "error"}),
				),
				// err := json.Unmarshal(resp, &respObj)
				jen.ID("err").Op(":=").Qualified("encoding/json", "Unmarshal").Call(jen.ID("resp"), jen.ID("&respObj")),
				// if err != nil {
				// 	return PostFeeTransactionResponse{}, err
				// }
				jen.If(jen.ID("err").Op("!=").Nil()).Block(
					jen.Return().ID(opIdTypeName+"Response").Block().Op(",").ID("err"),
				),
				// return PostFeeTransactionResponse{}, errors.New(respObj.Error)
				jen.Return().ID(opIdTypeName+"Response").Block().Op(",").Qualified("errors", "New").Call(jen.ID("respObj").Dot("Error")),
			),
		),
	)

	if operation.Description != "" {
		// // detect if there are references (like `transaction_payload`) and replace them with the actual values
		// // pull each reference by backticks and replace with cleanID(reference)
		// splitStr := strings.Split(operation.Description, "`")
		// for i := 0; i < len(splitStr); i++ {
		// 	if i % 2 == 0 {
		// 		// make sure it's not surrounded by no none-whitespace characters
		// 		if strings.TrimSpace(splitStr[i]) == "" {
		// 			splitStr[i] = cleanID(splitStr[i])
		// 		}
		// 	}
		// }

		f.Comment(operation.Description)
	}
	f.Func().ID(opIdTypeName).Parameters(
		inputParams.Generated...,
	).Parameters(possibleResponseTypes...).Block(
		funcCode.Generated...,
	).Line()
}

func processObjectOrArray(schema *openapi3.Schema, tls *sync.Map, f *jen.File, backupTitle string) (string, Code) {
	var q Code

	var targetTitle string

	if schema.Title != "" {
		targetTitle = schema.Title
	} else {
		targetTitle = backupTitle
	}
	targetTitle = cleanID(targetTitle)

	// has the schema been processed already?
	// if _, ok := tls.Load(targetTitle); !ok {
		if schema.Type != "" {
			switch schema.Type {
			case "object":
				props := schema.Properties

				for k, v := range props {
					prop := v.Value
					tag := map[string]string{"json": k}

					cleanedKey := cleanID(k)

					switch prop.Type {
					case "object":
						nestID, _ := processObjectOrArray(prop, tls, f, backupTitle+cleanID(k))
						q.add(jen.ID(cleanedKey).ID(nestID).Tag(tag))
					case "array":
						nestID, _ := processObjectOrArray(prop, tls, f, backupTitle+cleanID(k))
						q.add(jen.ID(cleanedKey).ID(nestID).Tag(tag))
					default:
						q.add(
							jen.ID(cleanedKey).ID(typeReplace(prop.Type)).Tag(tag),
						)
					}
				}
			case "array":
				items := schema.Items.Value
				if items.Enum != nil { // enum type
					var values []string // convert to string array
					for _, enum := range items.Enum {
						values = append(values, cleanID(enum.(string)))
					}

					var enums Code
			
					for index, value := range values {
						if index == 0 {
							enums.add(jen.ID(value).ID(targetTitle).Op("=").Iota())
						} else {
							enums.add(jen.ID(value))
						}
					}
			
					f.Type().ID(targetTitle).Int64()
					f.Const().Definitions(enums.Generated...)
				} else {
					return "[]"+targetTitle, q
				}
			}
		} else {
			if schema.AnyOf != nil {
				q.add(jen.Comment("This is an anyOf type, and so any of the following fields may be set"))

				keys := make(map[string]*openapi3.Schema)
				for _, anyOf := range schema.AnyOf {
					val := anyOf.Value
					for k, v := range val.Properties {
						keys[k] = v.Value
					}
				}
				for k, prop := range keys {
					tag := map[string]string{"json": k}
					cleanedKey := cleanID(k)
					switch prop.Type {
					case "object":
						nestID, _ := processObjectOrArray(prop, tls, f, backupTitle+cleanID(k))
						q.add(jen.ID(cleanedKey).ID(nestID).Tag(tag))
					case "array":
						nestID, _ := processObjectOrArray(prop, tls, f, backupTitle+cleanID(k))
						q.add(jen.ID(cleanedKey).ID(nestID).Tag(tag))
					default:
						q.add(
							jen.ID(cleanedKey).ID(typeReplace(prop.Type)).Tag(tag),
						)
					}
				}
			}
		}

		fmt.Println(targetTitle, q)

		tls.Store(targetTitle, q)
	// }

	return targetTitle, q
}
