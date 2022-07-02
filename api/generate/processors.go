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

func processResponse(response *openapi3.ResponseRef, opID string, statusCode string, queue *Code) {
	val := response.Value
	respTypeName := opID + cleanID(statusCode) + "Response"

	queue.add(jen.Commentf("Defines a %v response for %v.", statusCode, opID))

	if val.Description != nil {
		queue.add(jen.Comment(cleanDesc(*val.Description)))
	}

	queue.add(
		jen.Type().ID(respTypeName).Structure(),
	)
}
