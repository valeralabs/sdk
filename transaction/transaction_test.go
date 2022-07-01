package transaction

import (
	"testing"

	"github.com/valeralabs/sdk/constant"
)

func TestUnmarshalTokenTransferTransaction(test *testing.T) {
	var transaction StacksTransaction

	// blockstack-cli token-transfer edf9aee84d9b7abc145504dde6726c64f369d37ee34ded868fabd876c26570bc01 1 10 SP3TZ3NY4GB3E3Y1K1D40BHE07P20KMS4A8YC4QRJ 100 'hello'
	err := transaction.Unmarshal([]byte("0000000001040015c31b8c1c11c515e244b75806bac48d1399c775000000000000000a0000000000000001000135230d6a06749377212758f925871c1106fb258abed738df874a858f6f394881786d413dccfe8b134f0a8f820ce72220b8201b1cfc43ac365cf1e23f4f4003ab030200000000000516f5f1d7c482c6e1f8330b4805c5c03d8409d32452000000000000006468656c6c6f0000000000000000000000000000000000000000000000000000000000"))

	if err != nil {
		test.Fatalf("failed to unmarshal transaction: %v", err)
	}

	if transaction.Version != constant.TransactionVersionMainnet {
		test.Fatal("expected mainnet got testnet")
	}

	if transaction.ChainID != constant.ChainIDMainnet {
		test.Fatal("expected mainnet got testnet")
	}

	test.Logf("got transaction %#+v\n", transaction)
}

func TestUnmarshalSmartContractTransaction(test *testing.T) {
	var transaction StacksTransaction

	// blockstack-cli publish edf9aee84d9b7abc145504dde6726c64f369d37ee34ded868fabd876c26570bc01 1 20 example demo.clar
	// demo.clar: (print "Hello World")
	err := transaction.Unmarshal([]byte("0000000001040015c31b8c1c11c515e244b75806bac48d1399c775000000000000001400000000000000010000a5e33acf930ec7837f24403411f53e44474ea31a5f6069f15802348d0e113b850bcf586ffc2065031c2e3948cf9e1e9f313f915c2f3305d073d722dec5e4de1303020000000001076578616d706c6500000016287072696e74202268656c6c6f20776f726c6422290a"))

	if err != nil {
		test.Fatalf("failed to unmarshal smart contract: %v", err)
	}

	test.Logf("got transaction %#+v\n", transaction)
}

func TestUnmarshalContractCallTransaction(test *testing.T) {
	var transaction StacksTransaction

	// blockstack-cli contract-call edf9aee84d9b7abc145504dde6726c64f369d37ee34ded868fabd876c26570bc01 1 10 SP3TZ3NY4GB3E3Y1K1D40BHE07P20KMS4A8YC4QRJ demo example -e \"example\"
	err := transaction.Unmarshal([]byte("0000000001040015c31b8c1c11c515e244b75806bac48d1399c775000000000000000a000000000000000100015eb19a8598f72ac81fa96cb11c59edac0620236f534f7bff2363c413e099a4ea630f2cdc53cc28810e969bb1a4003d1b34ee6dfca7f87fe44ac81563a37abbb10302000000000216f5f1d7c482c6e1f8330b4805c5c03d8409d324520464656d6f076578616d706c65000000010d000000076578616d706c65"))

	if err != nil {
		test.Fatalf("failed to unmarshal transaction: %v", err)
	}

	payload := transaction.Payload.(ContractCallTransfer)

	if payload.Function != "example" {
		test.Fatalf("got wrong function name: %s", payload.Function)
	}

	if payload.Address.Contract != "demo" {
		test.Fatalf("got wrong contract name: %s", payload.Address.Contract)
	}

	test.Logf("got transaction %#+v\n", transaction)
}

//
// func TestMarshalTransaction(test *testing.T) {
// 	var transaction StacksTransaction
//
// 	_, err := transaction.Marshal()
//
// 	if err != nil {
// 		test.Fatalf("failed to marshal transaction: %v", err)
// 	}
// }
