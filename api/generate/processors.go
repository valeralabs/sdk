package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/valeralabs/jenny/jen"
)

func processParameter(file *jen.File, from *openapi3.ParameterRef, opID string) (parameter []jen.Code, argument []jen.Code, value []jen.Code) {
	val := from.Value
	schema := val.Schema.Value

	parameter = append(parameter, jen.Comment(cleanDesc(val.Description)))

	if val.Required == false {
		defaultStr := ""

		if schema.Default != nil {
			defaultStr = fmt.Sprintf("The default value is %v.", schema.Default)
		}

		parameter = append(parameter, jen.Commentf("Optional. "+defaultStr))

		if schema.Max != nil {
			parameter = append(parameter, jen.Comment(fmt.Sprintf("Max value is %v.", *schema.Max)))
		}
	}

	name := cleanID(val.Name)
	lower := name

	if unicode.IsLower(rune(name[1])) {
		lower = strings.ToLower(string(name[0])) + name[1:]
	}

	if name == "Type" {
		lower = "_type"
	}

	value = append(value, jen.ID(lower))

	if schema.Type == "array" || schema.Type == "object" {
		object := &jen.Statement{}

		id, _ := processObjectOrArray(object, schema, opID+name)

		file.Add(object)

		parameter = append(parameter,
			jen.ID(name).ID(id),
		)

		argument = append(argument,
			jen.ID(lower).ID(id),
		)
	} else {
		parameter = append(parameter,
			jen.ID(name).ID(typeReplace(val.Schema.Value.Type)),
		)

		argument = append(argument,
			jen.ID(lower).ID(typeReplace(val.Schema.Value.Type)),
		)
	}

	return parameter, argument, value
}

func processRequestBodyProperty(file *jen.File, prop *openapi3.SchemaRef, opID string, title string) {
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

	file.Type().ID(opID + "Body").Structure(properties)
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

func processPostReq(file *jen.File, opIdTypeName string, operation *openapi3.Operation, arguments []jen.Code, values []jen.Code, name string, responses *jen.Statement) {
	// POST
	inputs := arguments

	// if there is a request body, add it to the input params
	if operation.RequestBody != nil {
		inputs = append(inputs, jen.ID("body").ID(opIdTypeName+"Body"))
	}

	var body []jen.Code

	if len(arguments) > 0 {
		body = append(body,
			// url := fmt.Sprintf("%s%s", Network, fillPath(name, params))
			jen.ID("url").Op(":=").Qualified("fmt", "Sprintf").Call(
				jen.Literal("%s%s"),
				jen.ID("Network"),
				jen.ID("fillPath").Call(
					jen.Literal(name),
					jen.ID(opIdTypeName+"Params").Values(values...),
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
func processGetReq(file *jen.File, opIdTypeName string, operation *openapi3.Operation, arguments []jen.Code, values []jen.Code, name string, responses *jen.Statement) {
	var body []jen.Code

	if len(operation.Parameters) > 0 {
		body = append(body,
			// url := fmt.Sprintf("%s%s", Network, fillPath(name, params))
			jen.ID("url").Op(":=").Qualified("fmt", "Sprintf").Call(
				jen.Literal("%s%s"),
				jen.ID("Network"),
				jen.ID("fillPath").Call(
					jen.Literal(name),
					jen.ID(opIdTypeName+"Params").Values(values...),
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
		arguments...,
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
