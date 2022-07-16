package api

import (
	"encoding/hex"
	"testing"
)

var ExampleTokenTransfer []byte

func init() {
	ExampleTokenTransfer, _ = hex.DecodeString("0000000001040015c31b8c1c11c515e244b75806bac48d1399c775000000000000000a0000000000000001000135230d6a06749377212758f925871c1106fb258abed738df874a858f6f394881786d413dccfe8b134f0a8f820ce72220b8201b1cfc43ac365cf1e23f4f4003ab030200000000000516f5f1d7c482c6e1f8330b4805c5c03d8409d32452000000000000006468656c6c6f0000000000000000000000000000000000000000000000000000000000")
}

func TestBroadcastTransaction(test *testing.T) {
	err := BroadcastTransaction(ExampleTokenTransfer)

	if err != nil {
		test.Fatalf("failed to broadcast transaction because \"%v\"\n", err)
	}
}

func TestGetBalance(test *testing.T) {
	balances, err := GetBalance("SP3TZ3NY4GB3E3Y1K1D40BHE07P20KMS4A8YC4QRJ")

	if err != nil {
		test.Fatalf("failed to get balance because %v\n", err)
	}

	test.Logf("balance %+v\n", balances)
}
