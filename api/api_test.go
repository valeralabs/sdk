package api

import (
	"encoding/hex"
	"testing"
)

var ExampleTokenTransfer []byte

func init() {
	ExampleTokenTransfer, _ = hex.DecodeString("0000000001040041cead6d5e61da18235d0acd15506f5f02942175000000000000000400000000000000b400003d9c35971f6f6dd763e24e6c1494732223eb3cb0c8363fc9dad442841d41ba176c489c69159588c13e2021c090a7dafdccd3f668a59943fa0a95e63dc4e73980030200000000000516db153631f7699b18cf97c053c8e0545855df81dd00000000000003e868656c6c6f0000000000000000000000000000000000000000000000000000000000")
}

func TestBroadcastTransaction(test *testing.T) {
	err := BroadcastTransaction(ExampleTokenTransfer)

	if err != nil && err.Error() != "NotEnoughFunds" {
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

func TestNextNonce(test *testing.T) {
	nonce, err := NextNonce("SP3TZ3NY4GB3E3Y1K1D40BHE07P20KMS4A8YC4QRJ")

	if err != nil {
		test.Fatalf("failed to get next nonce because %v\n", err)
	}

	test.Logf("nonce %+v\n", nonce)
}

func TestEstimateFee(test *testing.T) {
	fee, err := EstimateFee([]byte("010568656c6c6f0000000e3b3b2068656c6c6f20776f726c64"))

	if err != nil {
		test.Fatalf("failed to estimate fee because %v\n", err)
	}

	test.Logf("fee %+v\n", fee)
}
