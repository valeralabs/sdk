package keys

import (
	"encoding/hex"
	"testing"
)

func TestKeys(test *testing.T) {
	decoded, _ := hex.DecodeString("edf9aee84d9b7abc145504dde6726c64f369d37ee34ded868fabd876c26570bc01")
	private, err := NewPrivateKey(decoded)

	if err != nil {
		test.Fatalf("failed to check secret key: %v", err)
	}

	if private.Compressed == false {
		test.Fatalf("expected private key to be compressed")
	}

	test.Logf("got private key %#+v\n", private)

	public := private.PublicKey()

	test.Logf("got public key %#+v\n", public)
}
