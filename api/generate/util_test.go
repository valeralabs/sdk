package main

import (
	"testing"
)

type GetBurnchainRewardsTotalByAddressParams struct {
	// Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format.
	Address string
}

type GetContractEventsByIDParams struct {
	// Max number of contract events to fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Contract identifier formatted as `<contract_address>.<contract_name>`.
	ContractID string
	// Include transaction data from unanchored (i.e. unconfirmed) microblocks.
	// Optional. Use `false` as default.
	Unanchored bool
	// Index of first contract event to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
}

type GetContractsByTraitParams struct {
	// Max number of contracts fetch.
	// Optional. Use `-1` to use the default value.
	Limit int
	// Index of first contract event to fetch.
	// Optional. Use `-1` to use the default value.
	Offset int
	// JSON abi of the trait.
	TraitABI string
}

func TestCleanID(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"get_address_mempool_transactions", "GetAddressMempoolTransactions"},
		{"run_faucet_stx", "RunFaucetSTX"},
		{"get_tx_list_details", "GetTXListDetails"},
	}

	for _, tt := range tests {
		testname := tt.input
		t.Run(testname, func(t *testing.T) {
			ans := cleanID(tt.input)
			if ans != tt.want {
				t.Errorf("got %v from %v, wanted %v", ans, tt.input, tt.want)
			}
		})
	}
}

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

func TestCleanDesc(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"include transac...ed) microblocks", "Include transac...ed) microblocks."},
		{"Include transac...ed) microblocks", "Include transac...ed) microblocks."},
		{"Include transac...ed) microblocks.", "Include transac...ed) microblocks."},
		{"include transac...ed) microblocks.", "Include transac...ed) microblocks."},
	}

	for _, tt := range tests {
		testname := tt.input
		t.Run(testname, func(t *testing.T) {
			ans := cleanDesc(tt.input)
			if ans != tt.want {
				t.Errorf("got %v from %v, wanted %v", ans, tt.input, tt.want)
			}
		})
	}
}
