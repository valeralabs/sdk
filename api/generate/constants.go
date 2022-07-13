package main

import (
	"regexp"

	"github.com/valeralabs/jenny/jen"
)

var (
	output = "../api.v2.gen.go"
	input  = "./source/docs/openapi.yaml"
	URLArgumentScope = regexp.MustCompile(`{([^}]*)}`)
)

var manuallyAddedTypes = map[string]jen.Code{
	"CallReadOnlyFunctionArguments": jen.Type().ID("CallReadOnlyFunctionArguments").Structure(
		jen.Comment("The simulated tx-sender"),
		jen.ID("Sender").String(),
		jen.Comment("An array of hex serialized Clarity values"),
		jen.ID("Arguments").Index().String(),
	),
}

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
