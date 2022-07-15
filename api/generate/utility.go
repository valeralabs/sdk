package main

import (
	"regexp"
	"strings"
)

var LowerCaseReplacements = regexp.MustCompile(`(?P<start>(.*(_|{)|\A))(?P<match>iso|nft|btc|stx|api|id|tx|ft|tld|abi|txid|tld)(?P<end>((_|}).*|\z))`)

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

// get_address_mempool_transactions -> GetAddressMempoolTransactions
func cleanID(ID string) string {
	ID = strings.Replace(ID, " ", "", -1)
	ID = strings.Title(ID)
	parts := strings.Split(clean(ID), "_")

	for i, part := range parts {
		parts[i] = strings.Title(part)
	}

	return clean(strings.Join(parts, ""))
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
	return clean(desc)
}
