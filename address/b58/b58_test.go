package b58

import (
	"testing"

	"github.com/valeralabs/racks/constant"
	"github.com/valeralabs/racks/transaction"
)

func TestAddress(test *testing.T) {
	decoded, _ := transaction.HexToBytes("edf9aee84d9b7abc145504dde6726c64f369d37ee34ded868fabd876c26570bc01")
	private, _ := transaction.NewPrivateKey(decoded)
	public := private.PublicKey()

	address, err := NewAddress([]transaction.PublicKey{public}, constant.AddressVersionMainnetSingleSignature, constant.HashModeP2PKH)

	if err != nil {
		test.Errorf("failed to parse address: %v", err)
	}

	if address != "12z4vLsXnbWC4rwtq962M97p7P2YjNThS4" {
		test.Errorf("expected \"12z4vLsXnbWC4rwtq962M97p7P2YjNThS4\" got \"%s\"", address)
	}

	test.Logf("address %s\n", address)
}
