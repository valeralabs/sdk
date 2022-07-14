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
