package address

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/valeralabs/sdk/address/b58"
	"github.com/valeralabs/sdk/address/c32"
	"github.com/valeralabs/sdk/constant"
	"github.com/valeralabs/sdk/keys"
)

type Address struct {
	Version  constant.AddressVersion
	Hash     []byte
	Contract string
}

func (address Address) B58() (string, error) {
	return b58.Encode(address.Hash, address.Version)
}

func (address Address) C32() (string, error) {
	encoded, err := c32.Encode(address.Hash, address.Version)

	if err != nil {
		return "", err
	}

	if address.Contract != "" {
		encoded += "." + address.Contract
	}

	return encoded, nil
}

func NewAddress(publicKeys []keys.PublicKey, version constant.AddressVersion, mode constant.HashMode) (Address, error) {
	var hash []byte

	if mode == constant.HashModeP2PKH {
		hash = btcutil.Hash160(publicKeys[0].Serialize())
	} else {
		return Address{}, fmt.Errorf("TODO: add segwit support")
	}

	encoded := make([]byte, hex.EncodedLen(len(hash)))
	hex.Encode(encoded, hash)

	return Address{
		Version: version,
		Hash:    encoded,
	}, nil
}
