package b58

import (
	"fmt"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/valeralabs/sdk/constant"
)

func ConvertVersion(version constant.AddressVersion) (byte, error) {
	switch version {
		case constant.AddressVersionMainnetPublicKeyHash:
			return 0, nil

		case constant.AddressVersionMainnetScriptHash:
			return 5, nil

		case constant.AddressVersionTestnetPublicKeyHash:
			return 111, nil

		case constant.AddressVersionTestnetScriptHash:
			return 196, nil

		default:
			return 0, fmt.Errorf("invalid version %d", version)
	}
}

func Encode(hash []byte, version constant.AddressVersion) (string, error) {
	converted, err := ConvertVersion(version)

	if err != nil {
		return "", err
	}

	return string(base58.CheckEncode(hash, converted)), nil
}
