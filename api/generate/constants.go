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

var httpStatusCodes = map[int]string{
	100: "Continue",
	101: "Switching Protocols",
	102: "Processing",
	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "Non-Authoritative Information",
	204: "No Content",
	205: "Reset Content",
	206: "Partial Content",
	207: "Multi-Status",
	208: "Already Reported",
	226: "IM Used",
	300: "Multiple Choices",
	301: "Moved Permanently",
	302: "Found",
	303: "See Other",
	304: "Not Modified",
	305: "Use Proxy",
	306: "(Unused)",
	307: "Temporary Redirect",
	308: "Permanent Redirect",
	400: "Bad Request",
	401: "Unauthorized",
	402: "Payment Required",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	406: "Not Acceptable",
	407: "Proxy Authentication Required",
	408: "Request Timeout",
	409: "Conflict",
	410: "Gone",
	411: "Length Required",
	412: "Precondition Failed",
	413: "Payload Too Large",
	414: "URI Too Long",
	415: "Unsupported Media Type",
	416: "Range Not Satisfiable",
	417: "Expectation Failed",
	418: "I'm a teapot",
	421: "Misdirected Request",
	422: "Unprocessable Entity",
	423: "Locked",
	424: "Failed Dependency",
	426: "Upgrade Required",
	428: "Precondition Required",
	429: "Too Many Requests",
	431: "Request Header Fields Too Large",
	451: "Unavailable For Legal Reasons",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Timeout",
	505: "HTTP Version Not Supported",
	506: "Variant Also Negotiates",
	507: "Insufficient Storage",
	508: "Loop Detected",
	510: "Not Extended",
	511: "Network Authentication Required",
}
