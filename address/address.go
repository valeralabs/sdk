package address

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"

	"github.com/valeralabs/sdk/encoding/c32"
	"github.com/valeralabs/sdk/wallet/keys"
)

type Address struct {
	Version  AddressVersion
	Hash     []byte
	Contract string
}

func (address Address) B58() (string, error) {
	version, err := address.Version.BitcoinVersion()

	if err != nil {
		return "", err
	}

	return base58.CheckEncode(address.Hash, version), nil
}

func (address Address) C32() (string, error) {
	version, err := address.Version.StacksVersion()

	if err != nil {
		return "", err
	}

	hexed := make([]byte, hex.EncodedLen(len(address.Hash)))
	hex.Encode(hexed, address.Hash)

	c32AddressBytes, err := c32.ChecksumEncode(hexed, version)

	if err != nil {
		return "", err
	}

	c32Address := string(c32AddressBytes)

	if address.Contract != "" {
		c32Address += "." + address.Contract
	}

	return fmt.Sprintf("S%s", c32Address), nil
}

// func NewAddress(publicKeys []keys.PublicKey, numSigs int, version AddressVersion) (Address, error) {
// 	var hash []byte

// 	switch version.HashMode {
// 	// case HashModeP2PKH:
// 	// 	if len(publicKeys) > 1 {
// 	// 		return Address{}, fmt.Errorf("P2PKH can only accept one public key")
// 	// 	}

// 	// 	hash = btcutil.Hash160(publicKeys[0].Serialize())

// 	case HashModeP2SH:
// 		var addressPublicKeys []*btcutil.AddressPubKey

// 		for _, publicKey := range publicKeys {
// 			addressPublicKey, err := btcutil.NewAddressPubKey(publicKey.Serialize(), &chaincfg.MainNetParams)
// 			if err != nil {
// 				return Address{}, err
// 			}

// 			addressPublicKeys = append(addressPublicKeys, addressPublicKey)
// 		}

// 		script, err := txscript.MultiSigScript(addressPublicKeys, numSigs)
// 		if err != nil {
// 			return Address{}, err
// 		}

// 		hash = btcutil.Hash160(script)

// 	// case HashModeP2WPKH:
// 	// 	scriptBuilder := txscript.NewScriptBuilder()
// 	// 	scriptBuilder.AddInt64(0)

// 	// 	keyHash := btcutil.Hash160(publicKeys[0].Serialize())
// 	// 	scriptBuilder.AddData(keyHash)

// 	// 	script, err := scriptBuilder.Script()
// 	// 	if err != nil {
// 	// 		return Address{}, err
// 	// 	}

// 	// 	hash = btcutil.Hash160(script)

// 	case HashModeP2WSH:
// 		var addressPublicKeys []*btcutil.AddressPubKey

// 		for _, publicKey := range publicKeys {
// 			addressPublicKey, err := btcutil.NewAddressPubKey(publicKey.Serialize(), &chaincfg.MainNetParams)
// 			if err != nil {
// 				return Address{}, err
// 			}

// 			addressPublicKeys = append(addressPublicKeys, addressPublicKey)
// 		}

// 		script, err := txscript.MultiSigScript(addressPublicKeys, numSigs)
// 		if err != nil {
// 			return Address{}, err
// 		}

// 		scriptHash := sha256.Sum256(script)

// 		witnessScriptBuilder := txscript.NewScriptBuilder()
// 		witnessScriptBuilder.AddInt64(0)
// 		witnessScriptBuilder.AddData(scriptHash[:])

// 		witnessScript, err := witnessScriptBuilder.Script()
// 		if err != nil {
// 			return Address{}, err
// 		}

// 		hash = btcutil.Hash160(witnessScript)
// 	}

// 	return Address{
// 		Version: version,
// 		Hash:    hash,
// 	}, nil
// }

func NewAddressSingleSignature(public keys.PublicKey, version AddressVersion) (Address, error) {
	var hash []byte

	switch version.HashMode {
	case HashModeP2PKH:
		hash = btcutil.Hash160(public.Serialize())

	case HashModeP2WPKH:
		scriptBuilder := txscript.NewScriptBuilder()
		scriptBuilder.AddInt64(0)

		keyHash := btcutil.Hash160(public.Serialize())
		scriptBuilder.AddData(keyHash)

		script, err := scriptBuilder.Script()
		if err != nil {
			return Address{}, err
		}

		hash = btcutil.Hash160(script)

	default:
		return Address{}, errors.New("hash mode is not single signature")

	}

	return Address{
		Version: version,
		Hash:    hash,
	}, nil
}

func NewAddressMultipleSignature(publics []keys.PublicKey, signatures int, version AddressVersion) (Address, error) {
	var hash []byte

	switch version.HashMode {
	case HashModeP2SH:
		var addressPublicKeys []*btcutil.AddressPubKey

		for _, publicKey := range publics {
			addressPublicKey, err := btcutil.NewAddressPubKey(publicKey.Serialize(), &chaincfg.MainNetParams)
			if err != nil {
				return Address{}, err
			}

			addressPublicKeys = append(addressPublicKeys, addressPublicKey)
		}

		script, err := txscript.MultiSigScript(addressPublicKeys, signatures)

		if err != nil {
			return Address{}, err
		}

		hash = btcutil.Hash160(script)

	case HashModeP2WSH:
		var addressPublicKeys []*btcutil.AddressPubKey

		for _, publicKey := range publics {
			addressPublicKey, err := btcutil.NewAddressPubKey(publicKey.Serialize(), &chaincfg.MainNetParams)
			if err != nil {
				return Address{}, err
			}

			addressPublicKeys = append(addressPublicKeys, addressPublicKey)
		}

		script, err := txscript.MultiSigScript(addressPublicKeys, signatures)

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

	default:
		return Address{}, errors.New("hash mode is not multiple signature")
	}

	return Address{
		Version: version,
		Hash:    hash,
	}, nil
}

func NewAddress(publics []keys.PublicKey, signatures int, version AddressVersion) (Address, error) {
	if len(publics) == 1 && (version.HashMode == HashModeP2PKH || version.HashMode == HashModeP2WPKH) {
		return NewAddressSingleSignature(publics[0], version)
	}

	return NewAddressMultipleSignature(publics, signatures, version)
}
