package keys

import (
	"errors"
	"fmt"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	libsecp256k1 "github.com/ethereum/go-ethereum/crypto/secp256k1"
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

func (key PrivateKey) PublicKey() PublicKey {
	return PublicKey{key.Value.PubKey(), key.Compressed}
}

func (key PrivateKey) Serialize() []byte {
	return key.Value.Serialize()
}

// Message must be 32 bytes. Libsecp256k1 will not allow a non 32 byte message. Message should be a hash of what is actually being signed.
func (key PrivateKey) SignRecoverable(message []byte) ([]byte, error) {
	if len(message) != 32 {
		return []byte{}, errors.New("message must be 32 bytes")
	}

	signature, err := libsecp256k1.Sign(message, key.Serialize())
	if err != nil {
		return []byte{}, fmt.Errorf("libsecp256k1 could not sign message: %v", err)
	}

	// Reording the serialization of each part of the signature
	// This is necessary to comply with how stacks serialized recoverable signatures
	V := signature[64]
	RS := signature[:64]

	reorderedSignature := []byte{}
	reorderedSignature = append(reorderedSignature, V)
	reorderedSignature = append(reorderedSignature, RS...)

	return reorderedSignature, nil
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
