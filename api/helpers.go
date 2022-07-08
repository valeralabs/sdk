package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/valeralabs/structs"
)

type Server string

const (
	ValeraMainnet Server = "mn.stx.valera.sh"
	ValeraTestnet Server = "tn.stx.valera.sh"
	HiroMainnet   Server = "stacks-node-api.mainnet.stacks.co"
	HiroTestnet   Server = "stacks-node-api.testnet.stacks.co"
)

func check(err error, save *error) {
	if err != nil {
		*save = err
	}
}

func makeGetReq(url string, returnedErr *error) ([]byte, int) {
	*returnedErr = nil
	resp, err := http.Get(url)
	check(err, returnedErr)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err, returnedErr)

	return body, resp.StatusCode
}

func makePostReq(url string, sendBody string, returnedErr *error) ([]byte, int) {
	*returnedErr = nil
	resp, err := http.Post(url, "application/json", strings.NewReader(sendBody))
	check(err, returnedErr)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err, returnedErr)

	return body, resp.StatusCode
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
