package main

import (
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/valeralabs/jenny/jen"
)

func processParameter(file *jen.File, parameter *openapi3.ParameterRef, opID string) {
	val := parameter.Value
	schema := val.Schema.Value

	parameters := &jen.Statement{}

	parameters.Add(jen.Comment(cleanDesc(val.Description)))

	if val.Required == false {
		defaultStr := ""

		if schema.Default != nil {
			defaultStr = fmt.Sprintf("The default value is %v.", schema.Default)
		}

		parameters.Add(jen.Commentf("Optional. " + defaultStr))

		if schema.Max != nil {
			parameters.Add(jen.Comment(fmt.Sprintf("Max value is %v.", *schema.Max)))
		}
	}

	if schema.Type == "array" || schema.Type == "object" {
		object := &jen.Statement{}

		id, _ := processObjectOrArray(object, schema, opID+cleanID(val.Name))

		file.Add(object)

		parameters.Add(
			jen.ID(cleanID(val.Name)).ID(id),
		)
	} else {
		parameters.Add(
			jen.ID(cleanID(val.Name)).ID(typeReplace(val.Schema.Value.Type)),
		).Line()
	}

	// file.Commentf("Defines parameters for %v", opID)
	// file.Type().ID(opID + "Params").Structure(parameters)
}

func processRequestBodyProperty(prop *openapi3.SchemaRef, opID string, title string) *jen.Statement {
	properties := &jen.Statement{}

	schema := prop.Value
	titlizedTitle := strings.Title(title)

	// check if it's present in manually added types
	if _, ok := manuallyAddedTypes[opID+titlizedTitle]; ok {
		properties.Add(
			jen.Comment(cleanDesc(schema.Description)),
			jen.ID(cleanID(title)).ID(opID+titlizedTitle),
		)
	} else {
		if schema.Type == "array" || schema.Type == "object" {
			id, _ := processObjectOrArray(properties, schema, opID+cleanID(titlizedTitle))

			properties.Add(
				jen.Comment(cleanDesc(schema.Description)),
				jen.ID(cleanID(title)).ID(id),
			)
		} else {
			properties.Add(
				jen.Comment(cleanDesc(schema.Description)),
				jen.ID(cleanID(title)).ID(typeReplace(schema.Type)),
			)
		}
	}

	return properties
}

func processResponse(file *jen.File, response *openapi3.ResponseRef, opID string, statusCode string) {
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
				object := &jen.Statement{}

				respType, _ = processObjectOrArray(object, schema, backupTitle)

				file.Add(object).Line()
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

			file.Add(
				jen.Commentf("Defines a %v%v response for %v.", statusCode, descMsg, opID),
				jen.Line(),
				jen.Type().ID(backupTitle).ID(respType),
				jen.Line(),
			)
		}
	}
}

func processPostReq(file *jen.File, opIdTypeName string, operation *openapi3.Operation, parameters *jen.Statement, name string, responses *jen.Statement) {
	// POST
	file.Type().ID(opIdTypeName + "Body").Structure(parameters)

	var inputs []jen.Code

	// if there are parameters, add them to the input params
	if len(operation.Parameters) > 0 {
		inputs = append(inputs, jen.ID("params").ID(opIdTypeName+"Params"))
	}
	// if there is a request body, add it to the input params
	if operation.RequestBody != nil {
		inputs = append(inputs, jen.ID("body").ID(opIdTypeName+"Body"))
	}

	var body []jen.Code

	if len(operation.Parameters) > 0 {
		body = append(body,
			// url := fmt.Sprintf("%s%s", Network, fillPath(name, params))
			jen.ID("url").Op(":=").Qualified("fmt", "Sprintf").Call(
				jen.Literal("%s%s"),
				jen.ID("Network"),
				jen.ID("fillPath").Call(
					jen.Literal(name),
					jen.ID("params"),
				),
				jen.Line(),
			),
		)
	} else {
		body = append(body,
			// url := fmt.Sprintf("%s%s", Network, name))
			jen.ID("url").Op(":=").Qualified("fmt", "Sprintf").Call(
				jen.Literal("%s%s"),
				jen.ID("Network"),
				jen.Literal(name),
				jen.Line(),
			),
		)
	}

	body = append(body,
		// var returnedErr error
		jen.Var().ID("returnedErr").Error(),
		jen.Line(),
		// sendBody, _ := json.Marshal(body)
		jen.ID("sendBody").Op(",").ID("_").Op(":=").Qualified("encoding/json", "Marshal").Call(jen.ID("body")),
		jen.Line(),
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
		file.Comment(operation.Description)
	}

	file.Func().ID(opIdTypeName).Parameters(
		inputs...,
	).Parameters(responses).Block(
		body...,
	).Line()
}

// GET
func processGetReq(file *jen.File, opIdTypeName string, operation *openapi3.Operation, name string, responses *jen.Statement) {
	var parameters []jen.Code

	// if there are parameters, add them to the input params
	if len(operation.Parameters) > 0 {
		parameters = append(parameters, jen.ID("params").ID(opIdTypeName+"Params"))
	}

	var body []jen.Code

	if len(operation.Parameters) > 0 {
		body = append(body,
			// url := fmt.Sprintf("%s%s", Network, fillPath(name, params))
			jen.ID("url").Op(":=").Qualified("fmt", "Sprintf").Call(
				jen.Literal("%s%s"),
				jen.ID("Network"),
				jen.ID("fillPath").Call(
					jen.Literal(name),
					jen.ID("params"),
				),
			),
		)
	} else {
		body = append(body,
			// url := fmt.Sprintf("%s%s", Network, name))
			jen.ID("url").Op(":=").Qualified("fmt", "Sprintf").Call(
				jen.Literal("%s%s"),
				jen.ID("Network"),
				jen.Literal(name),
			),
		)
	}

	body = append(body,
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
		file.Comment(operation.Description)
	}

	file.Func().ID(opIdTypeName).Parameters(
		parameters...,
	).Parameters(responses).Block(
		body...,
	).Line()
}

func processObjectOrArray(file *jen.Statement, schema *openapi3.Schema, backupTitle string) (string, *jen.Statement) {
	object := &jen.Statement{}

	targetTitle := schema.Title

	if targetTitle == "" {
		targetTitle = backupTitle
	}

	targetTitle = cleanID(targetTitle)

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
					nestID, _ := processObjectOrArray(file, prop, backupTitle+cleanID(k))

					object.Add(
						jen.ID(cleanedKey).ID(nestID).Tag(tag),
					).Line()

				case "array":
					nestID, _ := processObjectOrArray(file, prop, backupTitle+cleanID(k))

					object.Add(
						jen.ID(cleanedKey).ID(nestID).Tag(tag),
					).Line()

				default:
					object.Add(
						jen.ID(cleanedKey).ID(typeReplace(prop.Type)).Tag(tag),
					).Line()
				}
			}
		case "array":
			items := schema.Items.Value

			if items.Enum != nil { // enum type
				var values []string // convert to string array

				for _, enum := range items.Enum {
					values = append(values, cleanID(enum.(string)))
				}

				enums := &jen.Statement{}

				for index, value := range values {
					if index == 0 {
						enums.Add(jen.ID(value).ID(targetTitle).Op("=").Iota()).Line()
					} else {
						enums.Add(jen.ID(value)).Line()
					}
				}

				file.Type().ID(targetTitle).Int64().Line()
				file.Const().Definitions(enums).Line()
			} else if items.AnyOf != nil {
				object.Add(jen.Comment("This is an anyOf type, and so any of the following fields may be set"))

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
						nestID, _ := processObjectOrArray(file, prop, backupTitle+cleanID(k))

						object.Add(
							jen.ID(cleanedKey).ID(nestID).Tag(tag),
						).Line()

					case "array":
						nestID, _ := processObjectOrArray(file, prop, backupTitle+cleanID(k))

						object.Add(
							jen.ID(cleanedKey).ID(nestID).Tag(tag),
						).Line()

					default:
						object.Add(
							jen.ID(cleanedKey).ID(typeReplace(prop.Type)).Tag(tag),
						).Line()
					}
				}
			} else {
				if items.Title != "" {
					targetTitle = items.Title
				}

				cleanedKey := cleanID(targetTitle)

				switch items.Type {
				case "object":
					nestID, _ := processObjectOrArray(file, items, backupTitle+cleanedKey)
					object.Add(jen.ID(cleanedKey).ID(nestID)).Line()

				case "array":
					nestID, _ := processObjectOrArray(file, items, backupTitle+cleanedKey)
					object.Add(jen.ID(cleanedKey).ID(nestID)).Line()

				default:
					object.Add(
						jen.ID(cleanedKey).ID(typeReplace(items.Type)),
					).Line()
				}

				file.Type().ID(targetTitle).Structure(object).Line()

				return "[]" + targetTitle, object
			}
		}
	} else {
		if schema.AnyOf != nil {
			object.Add(jen.Comment("This is an anyOf type, and so any of the following fields may be set"))

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
					nestID, _ := processObjectOrArray(file, prop, backupTitle+cleanID(k))

					object.Add(
						jen.ID(cleanedKey).ID(nestID).Tag(tag),
					).Line()

				case "array":
					nestID, _ := processObjectOrArray(file, prop, backupTitle+cleanID(k))

					object.Add(
						jen.ID(cleanedKey).ID(nestID).Tag(tag),
					).Line()

				default:
					object.Add(
						jen.ID(cleanedKey).ID(typeReplace(prop.Type)).Tag(tag),
					).Line()
				}
			}
		}
	}

	file.Type().ID(targetTitle).Structure(object).Line()

	return targetTitle, object
}
