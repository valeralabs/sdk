package main

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/valeralabs/structs"
)

func clean(source string) string {
var LowerCaseReplacements = regexp.MustCompile(`(?P<start>(.*(_|{)|\A))(?P<match>nft|btc|stx|api|id|tx|ft|tld|abi)(?P<end>((_|}).*|\z))`)
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
	parts := strings.Split(clean(ID), "_")

	for i, part := range parts {
		parts[i] = strings.Title(part)
	}

	return strings.Join(parts, "")
}

// reverse cleanID
// GetAddressMempoolTransactions -> get_address_mempool_transactions
func reverseCleanID(ID string) string {
	parts := strings.Split(ID, "")
	for i, part := range parts {
		if part == "" {
			continue
		} else if i == 0 {
			parts[i] = strings.ToLower(part)
		} else if part[0] >= 'A' && part[0] <= 'Z' {
			// if it's preceded by a lowercase letter, proceed
			if i > 1 && parts[i-1] >= "a" && parts[i-1] <= "z" {
				parts[i] = "_" + strings.ToLower(part)
			}
		}
	}

	final := strings.Join(parts, "")
	// convert to lowercase
	return strings.ToLower(final)
}

// accept path and object, filling in params in path with object values
// /extended/v1/tx/{tx_id} -> /extended/v1/tx/123456789
// also adds data not specified in the path to the query string
func fillPath(path string, object interface{}) string {
	// convert object to map
	objectMap := structs.Map(object)

	for key, value := range objectMap {
		// check if value is default value
		switch value.(type) {
		case string:
			if value == "" {
				continue
			}
		case int:
			if value == -1 {
				continue
			}
		case bool:
			if value == false {
				continue
			}
		}

		target := "{" + reverseCleanID(key) + "}"

		if strings.Contains(path, target) {
			// if the target is present in path, replace it
			path = strings.Replace(path, target, url.PathEscape(fmt.Sprintf("%v", value)), -1)
		} else { // add it into the query string
			if strings.Contains(path, "?") {
				path += fmt.Sprintf(
					"&%v=%v",
					url.QueryEscape(reverseCleanID(key)),
					url.QueryEscape(fmt.Sprintf("%v", value)),
				)
			} else {
				path += fmt.Sprintf(
					"?%v=%v",
					url.QueryEscape(reverseCleanID(key)),
					url.QueryEscape(fmt.Sprintf("%v", value)),
				)
			}
		}
	}

	return path
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
