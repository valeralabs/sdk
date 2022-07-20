package api

import (
	"errors"
	"fmt"

	"github.com/linden/fetch"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type STXBalance struct {
	Balance  int `json:",string"`
	Sent     int `json:"total_sent,string"`
	Received int `json:"total_received,string"`
}

type FTBalance struct {
	Balance int `json:",string"`
}

type NFTBalance struct {
	Balance int `json:"count,string"`
}

type AccountBalance struct {
	STX STXBalance            `json:"stx"`
	FT  map[string]FTBalance  `json:"fungible_tokens"`
	NFT map[string]NFTBalance `json:"non_fungible_tokens"`
}

type rawNextNonce struct {
	NextNonce int `json:"possible_next_nonce"`
}

type rawEstimateFee struct {
	Estimations []struct {
		Fee int
	}
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
		if len(response.Body) == 0 {
			return fmt.Errorf("bad status %s", response.StatusText)
		}

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

// https://hirosystems.github.io/stacks-blockchain-api/#tag/Accounts/operation/get_account_nonces
func NextNonce(address string) (int, error) {
	response, err := fetch.Fetch[rawNextNonce](Server+"/extended/v1/address/"+address+"/nonces", fetch.Options{})

	if err != nil {
		return 0, err
	}

	if response.Status != 200 {
		return 0, errors.New("failed to get next nonce")
	}

	return response.Body.NextNonce, nil
}

// https://hirosystems.github.io/stacks-blockchain-api/#tag/Fees/operation/fetch_fee_rate
func EstimateFee(payload []byte) (int, error) {
	body, _ := sjson.Set("", "transaction_payload", payload)

	response, err := fetch.Fetch[rawEstimateFee](Server+"/v2/fees/transaction", fetch.Options{
		Headers: fetch.Headers{
			"Content-Type": "application/json",
		},
		Method: "POST",
		Body:   body,
	})

	if err != nil {
		return 0, err
	}

	if response.Status != 200 {
		return 0, errors.New("failed to get fee rate")
	}

	if len(response.Body.Estimations) < 2 {
		return 0, errors.New("did not return any estimations")
	}

	for _, estimation := range response.Body.Estimations {
		if estimation.Fee >= 180 {
			return estimation.Fee, nil
		}
	}

	return 0, errors.New("fee rate was too low")
}
