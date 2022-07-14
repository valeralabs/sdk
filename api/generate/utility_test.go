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
