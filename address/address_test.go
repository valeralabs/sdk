package address

import (
	"testing"

	"github.com/valeralabs/racks/constant"
	"github.com/valeralabs/racks/transaction"
)

func TestBasicAddress(test *testing.T) {
	decoded, _ := transaction.HexToBytes("edf9aee84d9b7abc145504dde6726c64f369d37ee34ded868fabd876c26570bc01")
	private, _ := transaction.NewPrivateKey(decoded)
	public := private.PublicKey()

	user, err := NewAddress([]transaction.PublicKey{public}, constant.AddressVersionMainnetSingleSignature, constant.HashModeP2PKH)

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
}
