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

		if b58 != "12z4vLsXnbWC4rwtq962M97p7P2YjNThS4" {
			test.Fatalf("expected \"12z4vLsXnbWC4rwtq962M97p7P2YjNThS4\" got \"%s\"", b58)
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

		c32, err = address.C32()
		if err != nil {
			test.Fatalf("Could not encode multsig address in c32, err: %v", err)
		}

		if c32 != "SMVEM417XQ69X2R07FPZ90Y29T8R63J8QJEQFP6H" {
			test.Fatalf("Expected \"SMVEM417XQ69X2R07FPZ90Y29T8R63J8QJEQFP6H\", got %v", c32);
		}

		b58, err = address.B58()
		if err != nil {
			test.Fatalf("Could not encode multisig address in base 58, err: %v", err)
		}

		if b58 != "2pa6Ai2CoCs15qAMtxZRX1DoUmHH" {
			test.Fatalf("Expected \"2pa6Ai2CoCs15qAMtxZRX1DoUmHH\" got \"%s\"", b58);
		}
	})
}

