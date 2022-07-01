package main

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/util"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/valeralabs/jenny/jen"
)

//go:generate git submodule update --init --recursive

var colourCodes = map[string]string{
	"reset":   "\033[0m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"blue":    "\033[34m",
	"yellow":  "\033[33m",
	"cyan":    "\033[36m",
	"magenta": "\033[35m",
	"white":   "\033[37m",
	"black":   "\033[30m",
	"gray":    "\033[90m",
}

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

var manuallyAddedTypes = map[string]jen.Code{
	"CallReadOnlyFunctionArguments": jen.Type().ID("CallReadOnlyFunctionArguments").Structure(
		jen.Comment("The simulated tx-sender"),
		jen.ID("Sender").String(),
		jen.Comment("An array of hex serialized Clarity values"),
		jen.ID("Arguments").Index().String(),
	),
}

func main() {
	fmt.Println("Loading OpenAPI 3.0 spec...")
	rootStart := time.Now()
	swagger, err := util.LoadSwagger(input)

	if err != nil {
		panic(err)
	}
	
	fmt.Printf("Loaded API spec in %v\n", time.Since(rootStart))
	processingStart := time.Now()

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

		if !strings.HasPrefix(name, "/rosetta/") {
			for _, operation := range path.Operations() {
				if operation != nil {
					startTime := time.Now()
					prefix := "➤ │ "
					opIdTypeName := cleanID(operation.OperationID)
					fmt.Printf("➤ ┌ "+string(colourCodes["blue"])+"%s\n"+string(colourCodes["reset"]), opIdTypeName)

					var params Code

					var wg sync.WaitGroup

					// ---- TYPES

					typePrefix := prefix + string(colourCodes["gray"]) + "[TYPE] " + string(colourCodes["reset"])

					// parameters
					for _, parameter := range operation.Parameters {
						wg.Add(1)
						parameter := parameter

						go func() {
							defer wg.Done()
							var queue Code
							processParameter(parameter, &params, opIdTypeName, f, &queue)
							fmt.Printf(typePrefix+"Parameter `%v` processed\n", parameter.Value.Name)
							if !parameter.Value.Required {
								fmt.Printf(prefix + "       └ " + string(colourCodes["yellow"]) + "Optional\n" + string(colourCodes["reset"]))
							}
						}()
					}

					// request body
					if operation.RequestBody != nil {
						for _, body := range operation.RequestBody.Value.Content {
							for title, prop := range body.Schema.Value.Properties {
								wg.Add(1)

								title := title
								prop := prop

								go func() {
									defer wg.Done()
									processRequestBodyProperty(prop, &params, opIdTypeName, title, f)
									fmt.Printf(typePrefix+"Body property `%v` processed\n", title)
								}()
							}
						}
					}

					// ---- FUNCTIONS

					// funcPrefix := prefix + string(colourCodes["gray"]) + "[FUNC] " + string(colourCodes["reset"])

					wg.Wait()

					ParamsTypeName := opIdTypeName + "Params"

					f.Commentf("%s defines parameters for %v", ParamsTypeName, opIdTypeName)
					f.Type().ID(ParamsTypeName).Structure(params.Generated...)

					fmt.Printf("➤ └ Completed in %v\n", time.Since(startTime))
				}
			}
		}
	}

	for _, code := range manuallyAddedTypes {
		f.Add(code)
	}

	fmt.Printf("➤ Rendering output to %v\n", output)
	err = f.Save(output)

	if err != nil {
		panic(err)
	}

	fmt.Printf(string(colourCodes["green"])+"Finished processing in %v\n"+string(colourCodes["reset"]), time.Since(processingStart))
}

func processParameter(parameter *openapi3.ParameterRef, params *Code, opID string, f *jen.File, queue *Code) {
	val := parameter.Value
	schema := val.Schema.Value

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

func cleanDesc(desc string) string {
	if desc == "" {
		return ""
	}
	desc = strings.Replace(desc, "\n", "", -1)
	desc = strings.ToUpper(string(desc[0])) + desc[1:]
	if !strings.HasSuffix(desc, ".") {
		desc += "."
	}
	return desc
}
