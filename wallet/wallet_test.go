package wallet

import (
	"testing"
)

var ExampleMasterKey = []byte{250, 11, 150, 241, 220, 116, 1, 18, 185, 11, 101, 226, 13, 240, 219, 48, 159, 224, 156, 199, 50, 69, 18, 130, 164, 205, 100, 233, 66, 149, 206, 135}

var (
	ExpectedSalt          = "ca85f623fe187fefdb80f63d82f7c6aa177976cb643b5196139203ef6fa530a3"
	ExpectedConfiguration = "777cb566928595cac75dbd3263dcb216af4924675b82c6e04d714e384052d0e2"

	ExpectedDataPrivateKey   = "17f97263b435852bb0bcb7c2b8b7ab031791868d184d59ddc8e30ae9bf6185c0"
	ExpectedStacksPrivateKey = "ccc25c8993bfc4b9beae1d5921c3a2edac4fcad52a1db403eedf20ba6ca112c701"
)

func TestWallet(test *testing.T) {
	wallet, err := NewWallet(ExampleMasterKey)

	if err != nil {
		test.Fatalf("failed to create wallet %v\n", err)
	}

	if string(wallet.Salt) != ExpectedSalt {
		test.Fatalf("failed to derive salt got \"%s\" expected \"%s\"", wallet.Salt, ExpectedSalt)
	}

	if string(wallet.Configuration) != ExpectedConfiguration {
		test.Fatalf("failed to derive configuration got \"%s\" expected \"%s\"", wallet.Configuration, ExpectedConfiguration)
	}

	test.Logf("got wallet %+v\n", wallet)

	test.Run("derive account", func(test *testing.T) {
		account, err := wallet.Account(1)

		if err != nil {
			test.Fatalf("failed to derive account %v", err)
		}

		if string(account.DataPrivateKey) != ExpectedDataPrivateKey {
			test.Fatalf("failed to derive data private key got \"%s\" expected \"%s\"", account.DataPrivateKey, ExpectedDataPrivateKey)
		}

		test.Logf("got account %+v\n", account)
	})
}

func TestPhrase(test *testing.T) {
	phrase, err := NewPhrase(12)

	if err != nil {
		test.Fatalf("failed to generate phrase %v", err)
	}

	test.Logf("phrase %v\n", phrase)

	seed, err := NewSeed(phrase, "password")

	if err != nil {
		test.Fatalf("failed to recover seed %v", err)
	}

	test.Logf("seed %v\n", seed)
}
