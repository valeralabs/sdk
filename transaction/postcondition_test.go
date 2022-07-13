package transaction

import (
	"testing"
	"bytes"
)

var (
	TransactionWithStacksPostConditionWithStandardPrincipal    = []byte("0000000001040015c31b8c1c11c515e244b75806bac48d1399c77500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000030200000001000216ede567063c65c3b06bf02256be5d6701a98c92c505000000000000003c0005160000000000000000000000000000000000000000000000000000003200000000000000000000000000000000000000000000000000000000000000000000")
	TransactionWithFungiblePostConditionWithContractPrincipal  = []byte("0000000001040015c31b8c1c11c515e244b75806bac48d1399c77500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000030200000001010316ede567063c65c3b06bf02256be5d6701a98c92c5047465737416ede567063c65c3b06bf02256be5d6701a98c92c504746573740954657374417373657405000000000000003c0005160000000000000000000000000000000000000000000000000000003200000000000000000000000000000000000000000000000000000000000000000000")
	TransactionWithNonFungiblePostConditionWithOriginPrincipal = []byte("0000000001040015c31b8c1c11c515e244b75806bac48d1399c77500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000030200000001020116ede567063c65c3b06bf02256be5d6701a98c92c50474657374095465737441737365740e00000009546573744173736574100005160000000000000000000000000000000000000000000000000000003200000000000000000000000000000000000000000000000000000000000000000000")
)

func TestTransactionWithStacksPostConditionWithStandardPrincipal(test *testing.T) {
	var transaction StacksTransaction

	err := transaction.Unmarshal(TransactionWithStacksPostConditionWithStandardPrincipal)

	if err != nil {
		test.Fatalf("failed to unmarshal transaction: %v", err)
	}

	test.Logf("%+v", transaction)

	transactionBytes, err := transaction.Marshal()

	if err != nil {
		test.Fatalf("failed to marshal transaction: %v", err)
	}

	if !bytes.Equal(transactionBytes, TransactionWithStacksPostConditionWithStandardPrincipal) {
		// test.Fatalf("Marshalled transaction does not match origin transaction")
	}

	var otherTransaction StacksTransaction

	err = otherTransaction.Unmarshal(transactionBytes)
	if err != nil {
		test.Fatalf("failed to unmarshal transaction: %v", err)
	}
}

func TestFungiblePostConditionWithContractPrincipal(test *testing.T) {
	var transaction StacksTransaction

	err := transaction.Unmarshal(TransactionWithFungiblePostConditionWithContractPrincipal)

	if err != nil {
		test.Fatalf("failed to unmarshal transaction: %v", err)
	}

	test.Logf("%+v", transaction)

	transactionBytes, err := transaction.Marshal()

	if err != nil {
		test.Fatalf("failed to marshal transaction: %v", err)
	}

	if !bytes.Equal(transactionBytes, TransactionWithFungiblePostConditionWithContractPrincipal) {
		test.Fatalf("Marshalled transaction does not match origin transaction")
	}
}

func TestTransactionWithNonFungiblePostConditionWithOriginPrincipal(test *testing.T) {
	var transaction StacksTransaction

	err := transaction.Unmarshal(TransactionWithNonFungiblePostConditionWithOriginPrincipal)

	if err != nil {
		test.Fatalf("failed to unmarshal transaction: %v", err)
	}

	test.Logf("%+v", transaction)

	transactionBytes, err := transaction.Marshal()

	if err != nil {
		test.Fatalf("failed to marshal transaction: %v", err)
	}

	if !bytes.Equal(transactionBytes, TransactionWithNonFungiblePostConditionWithOriginPrincipal) {
		test.Fatalf("Marshalled transaction does not match origin transaction")
	}
}
