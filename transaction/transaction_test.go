package transaction

import (
	"bytes"
	"encoding/hex"
	"testing"

	"valera.co/vdk/constant"
	"valera.co/vdk/wallet/keys"
)

var (
	// blockstack-cli token-transfer edf9aee84d9b7abc145504dde6726c64f369d37ee34ded868fabd876c26570bc01 1 10 SP3TZ3NY4GB3E3Y1K1D40BHE07P20KMS4A8YC4QRJ 100 'hello'
	ExampleTokenTransfer = []byte("0000000001040015c31b8c1c11c515e244b75806bac48d1399c775000000000000000a0000000000000001000135230d6a06749377212758f925871c1106fb258abed738df874a858f6f394881786d413dccfe8b134f0a8f820ce72220b8201b1cfc43ac365cf1e23f4f4003ab030200000000000516f5f1d7c482c6e1f8330b4805c5c03d8409d32452000000000000006468656c6c6f0000000000000000000000000000000000000000000000000000000000")

	// blockstack-cli publish edf9aee84d9b7abc145504dde6726c64f369d37ee34ded868fabd876c26570bc01 1 20 example demo.clar
	// demo.clar: (print "Hello World")
	ExampleSmartContract = []byte("0000000001040015c31b8c1c11c515e244b75806bac48d1399c775000000000000001400000000000000010000a5e33acf930ec7837f24403411f53e44474ea31a5f6069f15802348d0e113b850bcf586ffc2065031c2e3948cf9e1e9f313f915c2f3305d073d722dec5e4de1303020000000001076578616d706c6500000016287072696e74202268656c6c6f20776f726c6422290a")

	// blockstack-cli contract-call edf9aee84d9b7abc145504dde6726c64f369d37ee34ded868fabd876c26570bc01 1 10 SP3TZ3NY4GB3E3Y1K1D40BHE07P20KMS4A8YC4QRJ demo example -e \"example\"
	ExampleContractCall = []byte("0000000001040015c31b8c1c11c515e244b75806bac48d1399c775000000000000000a000000000000000100015eb19a8598f72ac81fa96cb11c59edac0620236f534f7bff2363c413e099a4ea630f2cdc53cc28810e969bb1a4003d1b34ee6dfca7f87fe44ac81563a37abbb10302000000000216f5f1d7c482c6e1f8330b4805c5c03d8409d324520464656d6f076578616d706c65000000010d000000076578616d706c65")
)

func TestTokenTransfer(test *testing.T) {
	var transaction StacksTransaction

	test.Run("unmarshal token transfer", func(test *testing.T) {
		err := transaction.Unmarshal(ExampleTokenTransfer)

		if err != nil {
			test.Fatalf("failed to unmarshal token transfer: %v", err)
		}

		if transaction.Version != constant.TransactionVersionMainnet {
			test.Fatal("expected mainnet got testnet")
		}

		if transaction.ChainID != constant.ChainIDMainnet {
			test.Fatal("expected mainnet got testnet")
		}

		test.Logf("unmarshaled %#+v\n", transaction)
	})

	test.Run("marshal token transfer", func(test *testing.T) {
		plain, err := transaction.Marshal()

		if err != nil {
			test.Fatalf("failed to marshal token transfer: %v", err)
		}

		if bytes.Compare(plain, ExampleTokenTransfer) != 0 {
			test.Fatalf("failed to marshal token transfer got %s expected %s", plain, ExampleTokenTransfer)
		}

		test.Logf("marshaled %s\n", plain)
	})
}

func TestSmartContract(test *testing.T) {
	var transaction StacksTransaction

	test.Run("unmarshal smart contract", func(test *testing.T) {
		err := transaction.Unmarshal(ExampleSmartContract)

		if err != nil {
			test.Fatalf("failed to unmarshal smart contract: %v", err)
		}

		test.Logf("unmarshaled %#+v\n", transaction)
	})

	test.Run("marshal smart contract", func(test *testing.T) {
		plain, err := transaction.Marshal()

		if err != nil {
			test.Fatalf("failed to marshal smart contract: %v", err)
		}

		if bytes.Compare(plain, ExampleSmartContract) != 0 {
			test.Fatalf("failed to marshal smart contract got %s expected %s", plain, ExampleTokenTransfer)
		}

		test.Logf("marshaled %s\n", plain)
	})
}

func TestContractCall(test *testing.T) {
	var transaction StacksTransaction

	test.Run("unmarshal contract call", func(test *testing.T) {
		err := transaction.Unmarshal(ExampleContractCall)

		if err != nil {
			test.Fatalf("failed to unmarshal transaction: %v", err)
		}

		payload := transaction.Payload.(PayloadContractCall)

		if payload.Function != "example" {
			test.Fatalf("got wrong function name: %s", payload.Function)
		}

		if payload.Address.Contract != "demo" {
			test.Fatalf("got wrong contract name: %s", payload.Address.Contract)
		}

		test.Logf("unmarshaled %#+v\n", transaction)
	})

	test.Run("marshal smart contract", func(test *testing.T) {
		plain, err := transaction.Marshal()

		if err != nil {
			test.Fatalf("failed to marshal contract call: %v", err)
		}

		if bytes.Compare(plain, ExampleContractCall) != 0 {
			test.Fatalf("failed to marshal contract call got %s expected %s", plain, ExampleContractCall)
		}

		test.Logf("marshaled %s\n", plain)
	})
}

func TestSigning(test *testing.T) {
	var transaction StacksTransaction

	transaction.Unmarshal(ExampleTokenTransfer)

	test.Logf("got transaction %+v\n", transaction.Authorization.GetCondition())

	// zeroing out the signature, in case the signature never gets changed resulting in a false positive
	empty := transaction.Authorization.GetCondition()
	empty.SetSignature([65]byte{})

	transaction.Authorization = transaction.Authorization.SetCondition(empty)

	decoded, err := hex.DecodeString("edf9aee84d9b7abc145504dde6726c64f369d37ee34ded868fabd876c26570bc01")

	if err != nil {
		test.Fatalf("could not decode private key hex: %v", err)
	}

	key, err := keys.NewPrivateKey(decoded)

	if err != nil {
		test.Fatalf("could not create private key: %v", err)
	}

	err = transaction.Sign(key)

	if err != nil {
		test.Fatalf("could not sign transaction: %v", err)
	}

	decoded, err = transaction.Marshal()

	if err != nil {
		test.Fatalf("could not marshal transaction: %v", err)
	}

	test.Logf("got transaction %+v %s\n", transaction.Authorization.GetCondition(), decoded)

	if bytes.Equal(decoded, ExampleTokenTransfer) == false {
		test.Fatalf("marshaled transaction is not same as initial transaction")
	}

	test.Logf("got transaction %s\n", decoded)
}
