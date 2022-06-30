package address

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/chaincfg"

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

func NewAddress(publicKeys []keys.PublicKey, numSigs int, version constant.AddressVersion, mode constant.HashMode) (Address, error) {
	var hash []byte

	if mode == constant.HashModeP2PKH || mode == constant.HashModeP2WPKH {
		if len(publicKeys) > 1 {
			return Address{}, fmt.Errorf("P2PKH can only accept one public key")
		}

		hash = btcutil.Hash160(publicKeys[0].Serialize())
	} else if mode == constant.HashModeP2SH || mode == constant.HashModeP2WSH {
		var addressPublicKeys []*btcutil.AddressPubKey

		for _, publicKey := range publicKeys {
			addressPublicKey, err := btcutil.NewAddressPubKey(publicKey.Serialize(), &chaincfg.MainNetParams)
			if err != nil {
				return Address{}, err
			}

			addressPublicKeys = append(addressPublicKeys, addressPublicKey)
		}

		script, err := txscript.MultiSigScript(addressPublicKeys, numSigs)
		if err != nil {
			return Address{}, err
		}

		hash = btcutil.Hash160(script)
	}

	encoded := make([]byte, hex.EncodedLen(len(hash)))
	hex.Encode(encoded, hash)

	return Address{
		Version: version,
		Hash:    encoded,
	}, nil
}
