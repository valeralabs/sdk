// additional testing for valera SDK
// this is a fail-safe, most of this is covered elsewhere
package VDK

import (
	"strings"
	"testing"
)

var (
	ExampleGoodSeedPhrase = "solar label private mandate rare decade appear super similar laundry abstract edge thrive carpet alert clap rapid camera fix pattern slab camera mammal recall"
	ExampleBadSeedPhrase  = "bad_phrase"
)

func TestWallet(test *testing.T) {
	for _, length := range []int{12, 15, 18, 21, 24} {
		phrase, err := NewPhrase(length)

		if err != nil {
			test.Fatalf("failed to create phrase %v\n", err)
		}

		count := len(strings.Split(phrase, " "))

		if length != count {
			test.Fatalf("phrase is not the correct length got %d expected %d\n", count, length)
		}
	}

	wallet, err := NewWalletFromPhrase(ExampleGoodSeedPhrase, "")

	if err != nil {
		test.Fatalf("failed to derive wallet from seed phrase %v\n", err)
	}

	test.Logf("wallet %+v\n", wallet)

	account, err := wallet.Account(0)

	if err != nil {
		test.Fatalf("failed to derive account from wallet %v\n", err)
	}

	test.Logf("account %+v\n", account)
	test.Logf("principal %+v\n", account.Principal)

	stacks, err := account.Principal.Stacks()

	if err != nil {
		test.Fatalf("failed to derive stacks address from principal: %v\n", err)
	}

	if stacks != "SP10WXBBDBSGXM613BM5CT5AGDXFG5511EME1D3M8" {
		test.Fatalf("got invalid stacks address %s\n", stacks)
	}

	test.Logf("stacks %s\n", stacks)

	bitcoin, err := account.Principal.Bitcoin()

	if err != nil {
		test.Fatalf("failed to derive bitcoin address from principal: %v\n", err)
	}

	if bitcoin != "16zxWfrkJJ2vjSRC21omNKfeKp8TL5uNAj" {
		test.Fatalf("got invalid bitcoin address %s\n", stacks)
	}

	test.Logf("bitcoin %s\n", bitcoin)

	test.Run("derive from bad seed phrase", func(test *testing.T) {
		_, err = NewWalletFromPhrase(ExampleBadSeedPhrase, "")

		if err == nil {
			test.Fatal("test was given invalid seed phrase but returned no error")
		}
	})
}

//./valera-sdk-demo token-transfer -P [...] -A 0 -S 1000 -R "SP3DHADHHYXMSP66FJZ057J70AHC5BQW1VQT9YGW8" -M "hello" -M 10 -F 180
func TestTokenTransfer(test *testing.T) {
	wallet, _ := NewWalletFromPhrase(ExampleGoodSeedPhrase, "")
	account, _ := wallet.Account(0)
	principal, _ := NewPrincipal("SP27X4NTRKGZ4C0G1FQ8WATM95JNNZKBQ4NSHDGE")

	transfer, err := NewTokenTransfer(principal, 10, "hello")

	if err != nil {
		test.Fatalf("failed to transfer %v\n", err)
	}

	err = transfer.Sign(account)

	if err != nil {
		test.Fatalf("failed to sign %v\n", err)
	}

	encoded, err := transfer.Encode()

	if err != nil {
		test.Fatalf("failed to encode %v\n", err)
	}

	test.Logf("transaction %s\n", encoded)

	err = transfer.Broadcast()

	if err != nil && err.Error() != "NotEnoughFunds" {
		test.Fatalf("failed to broadcast %v\n", err)
	}
}

func TestNewSmartContract(test *testing.T) {
	wallet, _ := NewWalletFromPhrase(ExampleGoodSeedPhrase, "")
	account, _ := wallet.Account(0)

	contract, err := NewSmartContract("example", ";; hello world")

	if err != nil {
		test.Fatalf("failed to create contract %v\n", err)
	}

	err = contract.Sign(account)

	if err != nil {
		test.Fatalf("failed sign %v\n", err)
	}

	err = contract.Broadcast()

	if err != nil && err.Error() != "NotEnoughFunds" {
		test.Fatalf("%v\n", err)
	}
}

func TestContractCall(test *testing.T) {
	wallet, _ := NewWalletFromPhrase(ExampleGoodSeedPhrase, "")
	account, _ := wallet.Account(0)
	principal, _ := NewPrincipal("SP3TZ3NY4GB3E3Y1K1D40BHE07P20KMS4A8YC4QRJ.straight-ivory-scallop")

	list := NewClarityList()

	value, err := NewClarityValue(1, "100")

	if err != nil {
		test.Fatalf("failed to encode %v\n", err)
	}

	list.Add(value)

	call, err := NewContractCall(principal, "respond", list)

	if err != nil {
		test.Fatalf("failed to transfer %v\n", err)
	}

	err = call.Sign(account)

	if err != nil {
		test.Fatalf("failed sign %v\n", err)
	}

	err = call.Broadcast()

	if err != nil && err.Error() != "NotEnoughFunds" {
		test.Fatalf("failed to broadcast %v\n", err)
	}
}

func TestPostConditions(test *testing.T) {
	wallet, _ := NewWalletFromPhrase(ExampleGoodSeedPhrase, "")
	account, _ := wallet.Account(0)
	principal, _ := NewPrincipal("SP3TZ3NY4GB3E3Y1K1D40BHE07P20KMS4A8YC4QRJ.straight-ivory-scallop")

	list := NewClarityList()

	value, err := NewClarityValue(1, "200")

	if err != nil {
		test.Fatalf("failed to encode %v\n", err)
	}

	list.Add(value)

	call, err := NewContractCall(principal, "respond", list)

	if err != nil {
		test.Fatalf("failed to transfer %v\n", err)
	}

	conditions := NewPostCondition()

	for index := 1; index <= 5; index++ {
		err = conditions.AddSTX(index, index*1000, nil)

		if err != nil {
			test.Fatalf("failed to create condition %d %v", index, err)
		}
	}

	other, err := NewPrincipal("SP27X4NTRKGZ4C0G1FQ8WATM95JNNZKBQ4NSHDGE")

	if err != nil {
		test.Fatalf("failed to other principal %v", err)
	}

	for index := 1; index <= 5; index++ {
		err = conditions.AddSTX(index, index*1000, other)

		if err != nil {
			test.Fatalf("failed to create condition %d %v", index, err)
		}
	}

	contract, err := NewPrincipal("SP27X4NTRKGZ4C0G1FQ8WATM95JNNZKBQ4NSHDGE.example")

	if err != nil {
		test.Fatalf("failed to create contract %v", err)
	}

	for index := 1; index <= 5; index++ {
		err = conditions.AddSTX(index, index*1000, contract)

		if err != nil {
			test.Fatalf("failed to create condition %d %v", index, err)
		}
	}

	asset, err := NewAsset(contract, "EXA")

	if err != nil {
		test.Fatalf("failed to create asset %v", err)
	}

	for index := 1; index <= 5; index++ {
		err = conditions.AddFT(index, index*1000, other, asset)

		if err != nil {
			test.Fatalf("failed to create condition %d %v", index, err)
		}
	}

	call.SetCondition(conditions, false)

	err = call.Sign(account)

	if err != nil {
		test.Fatalf("failed sign %v\n", err)
	}

	err = call.Broadcast()

	if err != nil && err.Error() != "NotEnoughFunds" {
		test.Fatalf("failed to broadcast %v\n", err)
	}
}
