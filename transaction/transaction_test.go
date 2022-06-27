package transaction

import (
	"testing"
)

func TestUnmarshalTransaction(test *testing.T) {
	var transaction StacksTransaction

	// from: https://github.com/fungible-systems/micro-stacks/blob/6273a16be549ce44500151d67e44c13d643a702b/src/transactions/transaction.test.ts#L80
	decoded, err := HexToBytes("8000000001040015c31b8c1c11c515e244b75806bac48d1399c775000000000000000000000000000000000000a617adb" +
		"703b474b6569244733e416f9faa2269d48b6a7e0773bfb122bb3de4e24ab61a8ec5456bc6352d8faa04ff6ef335937a29" +
		"f606f4ccad0df191d69ad704030200000001000216df0ba3e79792be7be5e50a370289accfc8c9e032030000000000000" +
		"000000516df0ba3e79792be7be5e50a370289accfc8c9e03200000000002625a06d656d6f20286e6f7420696e636c7564" +
		"656400000000000000000000000000000000")

	err = transaction.Unmarshal(decoded)

	if err != nil {
		test.Errorf("failed to unmarshal transaction: %v", err)
	}

	test.Logf("got transaction %#+v\n", transaction)
}

func TestMarshalTransaction(test *testing.T) {
	var transaction StacksTransaction

	_, err := transaction.Marshal()

	if err != nil {
		test.Errorf("failed to marshal transaction: %v", err)
	}
}

func TestKeys(test *testing.T) {
	decoded, _ := HexToBytes("edf9aee84d9b7abc145504dde6726c64f369d37ee34ded868fabd876c26570bc01")

	private, err := NewPrivateKey(decoded)

	if err != nil {
		test.Errorf("failed to check secret key: %v", err)
	}

	if private.Compressed == false {
		test.Error("expected private key to be compressed")
	}

	test.Logf("got private key %#+v\n", private)

	public := private.PublicKey()

	test.Logf("got public key %#+v\n", public)
}
