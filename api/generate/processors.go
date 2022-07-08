package main

import (
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/valeralabs/jenny/jen"
)

func processParameter(parameter *openapi3.ParameterRef, params *Code, opID string, f *jen.File) {
	val := parameter.Value
	schema := val.Schema.Value
	var queue Code

	queue.add(jen.Comment(cleanDesc(val.Description)))

	if !val.Required {
		switch schema.Type {
		case "string":
			queue.add(jen.Commentf("Optional. Pass an empty string to use the default value."))
		case "integer":
			defaultStr := "."
			if schema.Default != nil {
				defaultStr = fmt.Sprintf(" (%v).", schema.Default)
			}
			queue.add(jen.Commentf("Optional. Use `-1` to use the default value" + defaultStr))
		case "boolean":
			queue.add(jen.Commentf("Optional. Use `" + fmt.Sprint(schema.Default) + "` as default."))
		}
		if schema.Max != nil {
			queue.add(jen.Comment(fmt.Sprintf("Max value is %v.", *schema.Max)))
		}
	}

	if schema.Type == "array" {
		if schema.Items.Value.Enum != nil { // enum type
			var values []string // convert to string array
			for _, enum := range schema.Items.Value.Enum {
				values = append(values, cleanID(enum.(string)))
			}

			enumTypeName := opID + cleanID(val.Name)

			var enums Code

			for index, value := range values {
				if index == 0 {
					enums.add(jen.ID(value).ID(enumTypeName).Op("=").Iota())
				} else {
					enums.add(jen.ID(value))
				}
			}

			f.Type().ID(enumTypeName).Int64()
			f.Const().Definitions(enums.Generated...)

			queue.add(
				jen.Commentf("%v", values),
				jen.ID(cleanID(val.Name)).ID(enumTypeName),
			)
		}
	} else {
		queue.add(
			jen.ID(cleanID(val.Name)).ID(typeReplace(val.Schema.Value.Type)),
		)
	}

	params.add(queue.Generated...)
}

func processRequestBodyProperty(prop *openapi3.SchemaRef, params *Code, opID string, title string, f *jen.File) {
	val := prop.Value
	titlizedTitle := strings.Title(title)

	// check if it's present in manually added types
	if _, ok := manuallyAddedTypes[opID+titlizedTitle]; ok {
		params.add(
			jen.Comment(cleanDesc(val.Description)),
			jen.ID(cleanID(title)).ID(opID+titlizedTitle),
		)
	} else {
		params.add(
			jen.Comment(cleanDesc(val.Description)),
			jen.ID(cleanID(title)).ID(typeReplace(val.Type)),
		)
	}
}

func processResponse(response *openapi3.ResponseRef, opID string, statusCode string, f *jen.File) {
	val := response.Value

	var respTypeName string
	if statusCode == "200" {
		respTypeName = opID + "Response"
	} else {
		return
		// respTypeName = opID + statusCode + "Error"
	}

	var props Code

	for _, content := range val.Content {
		if content.Schema != nil {
			for name, prop := range content.Schema.Value.Properties {
				schema := prop.Value
				tag := map[string]string{"json": name}

				if schema.Description != "" {
					props.add(jen.Comment(cleanDesc(schema.Description)))
				}

				if schema.Type == "array" {
					switch schema.Items.Value.Type {
					case "string":
						props.add(
							jen.ID(cleanID(name)).Index().String().Tag(tag),
						)
					case "integer":
						props.add(
							jen.ID(cleanID(name)).Index().Int().Tag(tag),
						)
					case "boolean":
						props.add(
							jen.ID(cleanID(name)).Index().Bool().Tag(tag),
						)
					default:
						props.add(
							jen.ID(cleanID(name)).Index().Any().Tag(tag),
						)
					}
				} else {
					props.add(
						jen.ID(cleanID(name)).ID(typeReplace(schema.Type)).Tag(tag),
					)
				}
			}
		}
	}

	descMsg := ""
	if val.Description != nil {
		descMsg = fmt.Sprintf(" (%v)", *val.Description)
	}

	f.Add(
		jen.Commentf("Defines a %v%v response for %v.", statusCode, descMsg, opID),
		jen.Line(),
		jen.Type().ID(respTypeName).Structure(props.Generated...),
	)
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

	if (operation.Description != "") {
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
	f.Func().ID(funcName(opIdTypeName)).Parameters(
		inputParams.Generated...,
	).Parameters(possibleResponseTypes...).Block(
		funcCode.Generated...,
	).Line()
}