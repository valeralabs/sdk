package main

import (
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
var replacements = []string{"nft", "btc", "stx", "api", "id", "tx", "ft", "tld", "abi"}
var PartialLowerCase = regexp.MustCompile(`(.*|\A)(Nft|Btc|Stx|Api|Id|Tx|Ft|Tld|Abi)([A-Z]|\z)`)

func clean(source string) string {
	for _, cursor := range replacements {
		source = strings.ReplaceAll(source, cursor, strings.ToUpper(cursor))
	}

	source = PartialLowerCase.ReplaceAllStringFunc(source, func(cursor string) string {
		found := PartialLowerCase.FindStringSubmatch(cursor)
		found = found[1:]

		if len(found) == 1 {
			return strings.ToUpper(found[0])
		}

		result := found[0] + strings.ToUpper(found[1])

		if len(found) == 3 {
			result = found[2]
		}

		return result
	})

	return source
}

func main() {
	swagger, err := util.LoadSwagger(input)

	if err != nil {
		panic(err)
	}

	// type generation
	f := jen.NewFile("api")

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
				typeName := cleanID(operation.OperationID)

				var params []jen.Code

				// loop through parameters
				for _, parameter := range operation.Parameters {
					val := parameter.Value

					if val.Schema.Value.Type == "array" {
						if val.Schema.Value.Items.Value.Enum != nil {
							// enum type
							// convert to string array
							var values []string

							for _, enum := range val.Schema.Value.Items.Value.Enum {
								values = append(values, cleanID(enum.(string)))
							}

							typeNameIota := typeName + cleanID(val.Name)

							var enums []jen.Code

							for index, value := range values {
								if index == 0 {
									enums = append(
										enums,
										jen.ID(value).ID(typeNameIota).Op("=").Iota(),
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
				f.Type().ID(typeName + "Parameters").Structure(params...)
			}
		}
	}

	err = f.Save(output)

	if err != nil {
		panic(err)
	}
}

// get_address_mempool_transactions -> GetAddressMempoolTransactions
func cleanID(ID string) string {
	parts := strings.Split(clean(ID), "_")

	for i, part := range parts {
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
		return "any"
	case "array":
		return "any"
	default:
		return src
	}
}
