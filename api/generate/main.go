package main

import (
	"fmt"
	"regexp"
	"strings"
	// "sync"

	"github.com/deepmap/oapi-codegen/pkg/util"
	"github.com/getkin/kin-openapi/openapi3"
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

type Code struct {
	Generated []jen.Code
}

func (p *Code) add(code ...jen.Code) {
	p.Generated = append(p.Generated, code...)
}

func main() {
	swagger, err := util.LoadSwagger(input)

	if err != nil {
		panic(err)
	}

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

		// if path doesn't start with "/rosetta/", process the path
		if !strings.HasPrefix(name, "/rosetta/") {
			for _, operation := range path.Operations() {
				if operation != nil {
					opIdTypeName := cleanID(operation.OperationID)

					params := Code{}

					// var wg sync.WaitGroup

					// loop through parameters
					for index, parameter := range operation.Parameters {
						// wg.Add(1)

						// index := index
						// parameter := parameter

						// go func() {
						// 	defer wg.Done()
						// 	processParameter(parameter, &params, index)
						// }()
						val := parameter.Value
						schema := val.Schema.Value

						if schema.Type == "array" {
							if schema.Items.Value.Enum != nil { // enum type
								var values []string // convert to string array
								for _, enum := range schema.Items.Value.Enum {
									values = append(values, cleanID(enum.(string)))
								}

								// create the enum type name
								enumTypeName := cleanID(val.Name)

								enums := Code{}

								for index, value := range values {
									if index == 0 {
										enums.add(jen.ID(value).ID(enumTypeName).Op("=").Iota())
									} else {
										enums.add(jen.ID(value))
									}
								}

								params.add(
									jen.Comment(val.Description),
									jen.Commentf("%v", values),
									jen.ID(cleanID(val.Name)).ID(enumTypeName),
								)
							}
						} else {
							params.add(
								jen.Comment(val.Description),
								jen.ID(cleanID(val.Name)).ID(typeReplace(val.Schema.Value.Type)),
							)
						}

						fmt.Printf("Parameter %v processed\n", index)
					}

					if operation.RequestBody != nil {
						for _, body := range operation.RequestBody.Value.Content {
							// fmt.Println("[")
							for title, prop := range body.Schema.Value.Properties {
								val := prop.Value
								// fmt.Println(title, val)

								if val.Description != "" {
									params.add(jen.Comment(val.Description))
								}

								// if val.Properties != nil {
								if false {
									props := []jen.Code{}

									// for _, prop := range val.Properties {
									// 	if

									f.Type().ID(cleanID(title)).Interface(props...)
								} else {
									params.add(jen.ID(cleanID(title)).ID(typeReplace(val.Type)))
								}
							}
							// fmt.Println("]")
						}
					}

					ParamsTypeName := opIdTypeName + "Params"

					f.Commentf("%s defines parameters for %v", ParamsTypeName, opIdTypeName)
					f.Type().ID(ParamsTypeName).Structure(params.Generated...)
				}
			}
		}
	}

	err = f.Save(output)

	if err != nil {
		panic(err)
	}
}

func processParameter(parameter *openapi3.ParameterRef, params *Code, index int) {
	val := parameter.Value
	schema := val.Schema.Value

	if schema.Type == "array" {
		if schema.Items.Value.Enum != nil { // enum type
			var values []string // convert to string array
			for _, enum := range schema.Items.Value.Enum {
				values = append(values, cleanID(enum.(string)))
			}

			// create the enum type name
			enumTypeName := cleanID(val.Name)

			enums := Code{}

			for index, value := range values {
				if index == 0 {
					enums.add(jen.ID(value).ID(enumTypeName).Op("=").Iota())
				} else {
					enums.add(jen.ID(value))
				}
			}

			params.add(
				jen.Comment(val.Description),
				jen.Commentf("%v", values),
				jen.ID(cleanID(val.Name)).ID(enumTypeName),
			)
		}
	} else {
		params.add(
			jen.Comment(val.Description),
			jen.ID(cleanID(val.Name)).ID(typeReplace(val.Schema.Value.Type)),
		)
	}

	fmt.Printf("Parameter %v processed\n", index)
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
