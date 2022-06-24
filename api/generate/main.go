package main

import (
	// "fmt"
	"regexp"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/util"
	"github.com/valeralabs/jenny/jen"
)

//go:generate git submodule update --init --recursive

var (
	output = "../api.v2.gen.go"
	input  = "./source/docs/openapi.yaml"
)

var URLArgumentScope = regexp.MustCompile(`{([^}]*)}`)
var replacements = []string{"nft", "btc", "stx", "api", "id", "tx", "ft", "tld"}

func clean(source string) string {
	for _, cursor := range replacements {
		source = strings.ReplaceAll(source, cursor, strings.ToUpper(cursor))
	}

	return source
}

func main() {
	swagger, err := util.LoadSwagger(input)

	if err != nil {
		panic(err)
	}

	// clean up the swagger object
	for name, path := range swagger.Paths {
		delete(swagger.Paths, name)

		// TODO: add enum cleaning here

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

		for _, operation := range path.Operations() {
			if operation != nil {
				for _, parameter := range operation.Parameters {
					parameter.Value.Name = clean(parameter.Value.Name)
				}

				operation.OperationID = clean(operation.OperationID)
			}
		}

		swagger.Paths[name] = path
	}

	// type generation

	f := jen.NewFile("api")

	// loop through paths
	for _, operations := range swagger.Paths {
		for _, operation := range operations.Operations() {
			if operation != nil {
				typeName := cleanID(operation.OperationID)
				params := []jen.Code{}

				// loop through parameters
				for _, parameter := range operation.Parameters {
					val := parameter.Value
					if val.Schema.Value.Type == "array" {
						if val.Schema.Value.Items.Value.Enum != nil {
							// enum type
							// convert to string array
							values := []string{}
							for _, enum := range val.Schema.Value.Items.Value.Enum {
								values = append(values, cleanID(enum.(string)))
							}
							// fmt.Printf("%v\n", values)

							typeNameIota := typeName + cleanID(val.Name)

							enums := []jen.Code{}
							for index, value := range values {
								if index == 0 {
									enums = append(
										enums,
										jen.ID(values[0]).ID(typeNameIota).Op("=").Iota(),
									)
								} else {
									enums = append(
										enums,
										jen.ID(value),
									)
								}
							}
							
							f.Type().ID(typeNameIota).Int64()
							f.Const().Definitions(enums...)

							params = append(
								params,
								jen.Comment(val.Description),
								jen.Commentf("%v", values),
								jen.ID(cleanID(val.Name)).ID(typeNameIota),
							)
						}
					} else {
						params = append(
							params,
							jen.Comment(val.Description),
							jen.ID(cleanID(val.Name)).ID(typeReplace(val.Schema.Value.Type)),
						)
					}
				}

				if operation.RequestBody != nil {
					for _, body := range operation.RequestBody.Value.Content {
						// fmt.Println(body.Schema.Value.Properties)
						for title, prop := range body.Schema.Value.Properties {
							val := prop.Value
							if val.Description != "" {
								params = append(
									params,
									jen.Comment(val.Description),
								)
							}
							params = append(
								params,
								jen.ID(cleanID(title)).ID(typeReplace(val.Type)),
							)
						}
					}
				}
				f.Commentf("%sParams defines parameters for %v", typeName, typeName)
				f.Type().ID(typeName+"Params").Structure(params...)
			}
		}
	}
	// fmt.Printf("%#v", f)

	err = f.Save(output)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// get_address_mempool_transactions -> GetAddressMempoolTransactions
func cleanID(ID string) string {
	parts := strings.Split(ID, "_")

	for i, part := range parts {
		// i know this has a deprecation warning, but i don't care
		parts[i] = strings.Title(part)
	}

	return strings.Join(parts, "")
}

func typeReplace(src string) string {
	switch src {
	case "integer":
		return "int"
	case "number":
		return "int"
	case "string":
		return "string"
	case "boolean":
		return "bool"
	case "object":
		return "interface{}"
	case "array":
		return "interface{}"
	default:
		return src
	}
}
