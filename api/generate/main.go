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

	for name, path := range swagger.Paths {
		delete(swagger.Paths, name)

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

	// now to generate the types

	f := jen.NewFile("api")

	// loop through paths
	for _, operations := range swagger.Paths {
		for _, operation := range operations.Operations() {
			if operation != nil {
				typeName := cleanOpID(operation.OperationID)
				params := []jen.Code{}

				f.Commentf("%sParams defines parameters for %v", typeName, typeName)

				// loop through parameters
				for _, parameter := range operation.Parameters {
					val := parameter.Value
					// create the parameter statement
					// create the comment
					params = append(
						params,
						jen.Comment(val.Description),
						jen.ID(replaceKeyword(val.Name)).ID(typeReplace(val.Schema.Value.Type)),
						// jen.Line(),
					)
				}

				f.Type().ID(typeName+"Params").Structure(params...)
			}
		}
	}
	// fmt.Printf("%#v", f)

	f.Save(output)
}

// get_address_mempool_transactions -> GetAddressMempoolTransactionsParams
func cleanOpID(opID string) string {
	parts := strings.Split(opID, "_")

	for i, part := range parts {
		// i know this has a deprecation warning, but i don't care
		parts[i] = strings.Title(part)
	}

	opID = strings.Join(parts, "")
	return opID
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func replaceKeyword(src string) string {
	// ensure src doesn't collide with a go keyword (like "type")
	switch src {
	case "type":
		return "type_"
	default:
		return src
	}
}