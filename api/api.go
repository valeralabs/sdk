package api

import (
	"errors"

	"github.com/linden/fetch"
	"github.com/tidwall/gjson"
)

type Balance struct {
	Balance  int `json:",string"`
	Sent     int `json:"total_sent,string"`
	Received int `json:"total_received,string"`
}

type AccountBalance struct {
	STX Balance            `json:"stx"`
	FT  map[string]Balance `json:"fungible_tokens"`
	NFT map[string]Balance `json:"non_fungible_tokens"`
}

var Server = "https://stacks-node-api.mainnet.stacks.co"

// https://hirosystems.github.io/stacks-blockchain-api/#tag/Transactions/operation/post_core_node_transactions
func BroadcastTransaction(raw []byte) error {
	response, err := fetch.Fetch[string](Server+"/v2/transactions", fetch.Options{
		Headers: fetch.Headers{
			"Content-Type": "application/octet-stream",
		},
		Method: "POST",
		Body:   string(raw),
	})

	if err != nil {
		return err
	}

	if response.Status != 200 {
		if response.Body[0] == '"' {
			return errors.New(response.Body)
		}

		return errors.New(gjson.Get(response.Body, "reason").String())
	}

	return nil
}

// https://hirosystems.github.io/stacks-blockchain-api/#tag/Accounts/operation/get_account_balance
func GetBalance(address string) (AccountBalance, error) {
	response, err := fetch.Fetch[AccountBalance](Server+"/extended/v1/address/"+address+"/balances", fetch.Options{})

	if err != nil {
		return AccountBalance{}, err
	}

	if response.Status != 200 {
		return AccountBalance{}, errors.New("failed to get account balance")
	}

	return response.Body, nil
}
