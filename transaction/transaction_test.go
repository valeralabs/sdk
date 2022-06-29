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

// func TestUnmarshalContractCallTransaction(test *testing.T) {
// 	var transaction StacksTransaction
//
// 	// blockstack-cli contract-call edf9aee84d9b7abc145504dde6726c64f369d37ee34ded868fabd876c26570bc01 1 10 SP3TZ3NY4GB3E3Y1K1D40BHE07P20KMS4A8YC4QRJ demo example
// 	err := transaction.Unmarshal([]byte("0000000001040015c31b8c1c11c515e244b75806bac48d1399c775000000000000000a00000000000000010000b02ed8fbaef6c9c9f359feb9b535ca60b43c068cdd56cd1412eb6fb76147f10b461a36265fa7620a4dc0179c7c3820a81be6378c5ee3fd1f1a767804633572e60302000000000216f5f1d7c482c6e1f8330b4805c5c03d8409d324520464656d6f076578616d706c6500000000"))
//
// 	if err != nil {
// 		test.Fatalf("failed to unmarshal transaction: %v", err)
// 	}
//
// 	test.Logf("got transaction %#+v\n", transaction)
// }
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
