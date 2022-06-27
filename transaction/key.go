package transaction

import (
	"errors"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

type PublicKey secp256k1.PublicKey

type PrivateKey secp256k1.PrivateKey

func (key PrivateKey) PublicKey() PublicKey {
	raw := secp256k1.PrivateKey(key)

	return PublicKey(*raw.PubKey())
}

func NewPrivateKey(raw []byte) (PrivateKey, error) {
	if len(raw) == 33 {
		if raw[len(raw)-1] != 1 {
			return PrivateKey(secp256k1.PrivateKey{}), errors.New("private key has a length of 33 but is not compressed")
		}
	} else if len(raw) != 32 {
		return PrivateKey(secp256k1.PrivateKey{}), errors.New("private key must have a length of 33 or 32")
	}

	return PrivateKey(*secp256k1.PrivKeyFromBytes(raw)), nil
}
