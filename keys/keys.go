package keys

import (
	"errors"
	"fmt"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

type PublicKey struct {
	Value      *secp256k1.PublicKey
	Compressed bool
}

type PrivateKey struct {
	Value      *secp256k1.PrivateKey
	Compressed bool
}

func (key PublicKey) Serialize() []byte {
	if key.Compressed == true {
		return key.Value.SerializeCompressed()
	} else {
		return key.Value.SerializeUncompressed()
	}
}

func ParsePublicKey(bytes []byte, compressed bool) (PublicKey, error) {
	value, err := secp256k1.ParsePubKey(bytes)
	if err != nil {
		return PublicKey{}, fmt.Errorf("Could not parse public key: %v", err)
	}

	return PublicKey {
		Value: value,
		Compressed: compressed,
	}, nil
}

func (key PrivateKey) PublicKey() PublicKey {
	return PublicKey{key.Value.PubKey(), key.Compressed}
}

func NewPrivateKey(raw []byte) (PrivateKey, error) {
	compressed := false

	if len(raw) == 33 {
		if raw[len(raw)-1] != 1 {
			return PrivateKey{}, errors.New("private key has a length of 33 but is not compressed")
		} else {
			compressed = true
		}
	} else if len(raw) != 32 {
		return PrivateKey{}, errors.New("private key must have a length of 33 or 32")
	}

	return PrivateKey{secp256k1.PrivKeyFromBytes(raw), compressed}, nil
}
