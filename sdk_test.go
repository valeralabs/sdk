// additional testing for valera SDK
// this is a fail-safe, most of this is covered elsewhere
package ValeraSDK

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

	principal, err := account.Principal()

	if err != nil {
		test.Fatalf("failed to derive principal from account %v\n", err)
	}

	test.Logf("principal %+v\n", principal)

	stacks, err := principal.Stacks()

	if err != nil {
		test.Fatalf("failed to derive stacks address from principal: %v\n", err)
	}

	if stacks != "SP10WXBBDBSGXM613BM5CT5AGDXFG5511EME1D3M8" {
		test.Fatalf("got invalid stacks address %s\n", stacks)
	}

	test.Logf("stacks %s\n", stacks)

	bitcoin, err := principal.Bitcoin()

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

	transfer, err := NewTokenTransfer(principal, 10, "hello", nil, true)

	if err != nil {
		test.Fatalf("failed to transfer %v\n", err)
	}

	err = transfer.Sign(account, 180, 70)

	if err != nil {
		test.Fatalf("failed to sign %v\n", err)
	}

	encoded, err := transfer.Encode()

	if err != nil {
		test.Fatalf("failed to encode %v\n", err)
	}

	test.Logf("transaction %s\n", encoded)

	err = transfer.Broadcast(account)

	if err != nil {
		test.Fatalf("failed to broadcast %v\n", err)
	}
}

func TestNewSmartContract(test *testing.T) {
	wallet, _ := NewWalletFromPhrase(ExampleGoodSeedPhrase, "")
	account, _ := wallet.Account(0)

	contract, err := NewSmartContract("example", ";; hello world", nil, true)

	if err != nil {
		test.Fatalf("failed to create contract %v\n", err)
	}

	err = contract.Sign(account, 180, 68)

	if err != nil {
		test.Fatalf("failed sign %v\n", err)
	}

	err = contract.Broadcast(account)

	if err != nil {
		test.Fatalf("%v\n", err)
	}
}

func TestContractCall(test *testing.T) {
	wallet, _ := NewWalletFromPhrase(ExampleGoodSeedPhrase, "")
	account, _ := wallet.Account(0)

	principal, _ := NewPrincipal("SP466FNC0P7JWTNM2R9T199QRZN1MYEDTAR0KP27.miamicoin-token")

	list := NewClarityList()

	value, err := NewClarityValue(2, "hello world")

	if err != nil {
		test.Fatalf("failed to encode %v\n", err)
	}

	list.Add(value)

	call, err := NewContractCall(principal, "get-balance", list, nil, true)

	if err != nil {
		test.Fatalf("failed to transfer %v\n", err)
	}

	err = call.Sign(account, 1800, 69)

	if err != nil {
		test.Fatalf("failed sign %v\n", err)
	}

	err = call.Broadcast(account)

	if err != nil {
		test.Fatalf("failed to broadcast %v\n", err)
	}
}
