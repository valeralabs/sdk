package address

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/btcutil/base58"

	"github.com/valeralabs/sdk/encoding/c32"
	"github.com/valeralabs/sdk/keys"
)

type Address struct {
	Version  AddressVersion
	Hash     []byte
	Contract string
}

func (address Address) B58() (string, error) {
	versionByte, err := address.Version.GetVersionByte(AddressNetworkBitcoin)
	if err != nil {
		return "", err
	}

	return base58.CheckEncode(address.Hash, versionByte), nil
}

func (address Address) C32() (string, error) {
	versionByte, err := address.Version.GetVersionByte(AddressNetworkStacks)

	if err != nil {
		return "", fmt.Errorf("Could not get version byte: %v", err)
	}

	hashAsHex:= make([]byte, hex.EncodedLen(len(address.Hash)))
	hex.Encode(hashAsHex, address.Hash)

	c32AddressBytes, err := c32.ChecksumEncode(hashAsHex, versionByte)
	if err != nil {
		return "", fmt.Errorf("Could not encode in 32: %v", err)
	}

	c32Address := string(c32AddressBytes)

	if address.Contract != "" {
		c32Address += "." + address.Contract
	}
	
	return fmt.Sprintf("S%s", c32Address), nil
}

func NewAddress(publicKeys []keys.PublicKey, numSigs int, version AddressVersion, mode HashMode) (Address, error) {
	var hash []byte

	if mode == HashModeP2PKH {
		if len(publicKeys) > 1 {
			return Address{}, fmt.Errorf("P2PKH can only accept one public key")
		}

		hash = btcutil.Hash160(publicKeys[0].Serialize())

	} else if mode == HashModeP2SH {
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

	} else if mode == HashModeP2WPKH {
		scriptBuilder := txscript.NewScriptBuilder()
		scriptBuilder.AddInt64(0)

		keyHash := btcutil.Hash160(publicKeys[0].Serialize())
		scriptBuilder.AddData(keyHash)

		script, err := scriptBuilder.Script()
		if err != nil {
			return Address{}, err
		}

		hash = btcutil.Hash160(script)

	} else if mode == HashModeP2WSH {
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

		scriptHash := sha256.Sum256(script)

		witnessScriptBuilder := txscript.NewScriptBuilder()
		witnessScriptBuilder.AddInt64(0)
		witnessScriptBuilder.AddData(scriptHash[:])

		witnessScript, err := witnessScriptBuilder.Script()
		if err != nil {
			return Address{}, err
		}

		hash = btcutil.Hash160(witnessScript)
	}

	return Address{
		Version: version,
		Hash:    hash,
	}, nil
}

func NewAddressFromPublicKeyHash(hash []byte, version AddressVersion) Address {
	return Address {
		Version: version,
		Hash: hash,
	}
}

