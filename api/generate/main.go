package main

import (
	"fmt"
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
var LowerCaseReplacements = regexp.MustCompile(`(?P<start>(.*(_|{)|\A))(?P<match>nft|btc|stx|api|id|tx|ft|tld|abi)(?P<end>((_|}).*|\z))`)

func clean(source string) string {
	for LowerCaseReplacements.MatchString(source) {
		source = LowerCaseReplacements.ReplaceAllStringFunc(source, func(cursor string) string {
			found := LowerCaseReplacements.FindStringSubmatch(cursor)

			start := LowerCaseReplacements.SubexpIndex("start")
			match := LowerCaseReplacements.SubexpIndex("match")
			end := LowerCaseReplacements.SubexpIndex("end")

			return found[start] + strings.ToUpper(found[match]) + found[end]
		})
	}

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
						fmt.Println("[")
						for title, prop := range body.Schema.Value.Properties {
							val := prop.Value
							fmt.Println(title, val)
							
							// if val.Properties != nil {
							if false {
								props := []jen.Code{}

								// for _, prop := range val.Properties {
								// 	if 

								f.Type().ID(cleanID(title)).Interface(props...)
							} else {
								params = append(
									params,
									jen.ID(cleanID(title)).ID(typeReplace(val.Type)),
								)
							}

							if val.Description != "" {
								params = append(
									params,
									jen.Comment(val.Description),
								)
							}


						}
						fmt.Println("]")
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
