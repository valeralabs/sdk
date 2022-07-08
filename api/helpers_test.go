package api

import (
	"testing"
)

func TestReverseCleanID(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"GetAddressMempoolTransactions", "get_address_mempool_transactions"},
		{"RunFaucetSTX", "run_faucet_stx"},
		{"GetTxListDetails", "get_tx_list_details"},
	}

	for _, tt := range tests {
		testname := tt.input
		t.Run(testname, func(t *testing.T) {
			ans := reverseCleanID(tt.input)
			if ans != tt.want {
				t.Errorf("got %v from %v, wanted %v", ans, tt.input, tt.want)
			}
		})
	}
}

func TestFillPath(t *testing.T) {
	var tests = []struct {
		inputPath   string
		inputParams interface{}
		want        string
	}{
		{
			"/extended/v1/burnchain/rewards/{address}/total", 
			GetBurnchainRewardsTotalByAddressParams{
				Address: "B58",
			}, 
			"/extended/v1/burnchain/rewards/B58/total",
		},
		{
			"/extended/v1/contract/{contract_id}/events",
			GetContractEventsByIDParams{
				Limit: -1, // optional
				ContractID: "contract.contract",
				Unanchored: false, // optional
				Offset: -1, // optional
			},
			"/extended/v1/contract/contract.contract/events",
		},
		{
			"/extended/v1/contract/by_trait",
			GetContractsByTraitParams{
				Limit: 10, // optional
				Offset: -1, // optional
				TraitABI: "insert trait here", 
			},
			"/extended/v1/contract/by_trait?limit=10&trait_abi=insert+trait+here",
		},
	}

	for _, tt := range tests {
		testname := tt.inputPath
		t.Run(testname, func(t *testing.T) {
			ans := fillPath(tt.inputPath, tt.inputParams)
			if ans != tt.want {
				t.Errorf("got %v from %v, wanted %v", ans, tt.inputPath, tt.want)
			}
		})
	}
}
