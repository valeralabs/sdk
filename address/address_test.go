package address

import (
	"encoding/hex"
	"testing"

	"github.com/valeralabs/sdk/constant"
	"github.com/valeralabs/sdk/keys"
)

func TestBasicAddress(test *testing.T) {
	decoded, _ := hex.DecodeString("edf9aee84d9b7abc145504dde6726c64f369d37ee34ded868fabd876c26570bc01")
	private, _ := keys.NewPrivateKey(decoded)
	public := private.PublicKey()

	user, err := NewAddress([]keys.PublicKey{public}, 1, constant.AddressVersionMainnetSingleSignature, constant.HashModeP2PKH)

	publicKeyHex := hex.EncodeToString(public.Value.SerializeUncompressed())
	test.Logf("%v", publicKeyHex)


	test.Run("can create address", func(test *testing.T) {
		if err != nil {
			test.Fatalf("failed to parse address: %v", err)
		} else {
			test.Logf("got address %#+v\n", user)
		}
	})

	c32, err := user.C32()

	test.Run("can derive c32 address", func(test *testing.T) {
		if err != nil {
			test.Fatalf("failed to derive c32 address: %v", err)
		}

		if c32 != "SPAW66WC3G8WA5F28JVNG1NTRJ6H76E7EN5H6QQD" {
			test.Fatalf("expected \"SPAW66WC3G8WA5F28JVNG1NTRJ6H76E7EN5H6QQD\" got \"%s\"", c32)
		}

		test.Logf("got c32 address: %s\n", c32)
	})

	b58, err := user.B58()

	test.Run("can derive b58 address", func(test *testing.T) {
		if err != nil {
			test.Fatalf("failed to derive b58 address: %v", err)
		}

		if b58 != "1GzySdy8Es3QPmuF7osW13pN3UEFfHrLbC8VhjuNf2Qqd5Y4J9tLQFVwAfmKN" {
			test.Fatalf("expected \"1GzySdy8Es3QPmuF7osW13pN3UEFfHrLbC8VhjuNf2Qqd5Y4J9tLQFVwAfmKN\" got \"%s\"", c32)
		}

		test.Logf("got b58 address: %s\n", b58)
	})

	test.Run("can create multisig address", func(test *testing.T) {
		decodedOne, _ := hex.DecodeString("c3e1c944086ea6d61e0b9948a62d9608018c00a67424a817f005cf6bba39ce9001")
		privateOne, _ := keys.NewPrivateKey(decodedOne)
		publicOne := privateOne.PublicKey()

		decodedTwo, _ := hex.DecodeString("ee65f9526a229cff575fecdb2a06c565a51c0d466dbe207c6de683413259154901")
		privateTwo, _ := keys.NewPrivateKey(decodedTwo)
		publicTwo := privateTwo.PublicKey()

		publicKeys := []keys.PublicKey{publicOne, publicTwo}

		address, err := NewAddress(publicKeys, 2, constant.AddressVersionMainnetMultipleSignature, constant.HashModeP2SH)

		if err != nil {
			test.Fatalf("Could not create address, err: %v", err)
		}

		b58, err = address.B58()
		if err != nil {
			test.Fatalf("Could not encode multisig address in base 58, err: %v", err)
		}

		c32, err = address.C32()
		if err != nil {
			test.Fatalf("Could not encode multsig address in c32, err: %v", err)
		}

		test.Logf("Multisig Address:\nBase58 %v\nC32 %v", b58, c32)

	})
}

