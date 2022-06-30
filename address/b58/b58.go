package b58

import (
	"fmt"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/valeralabs/sdk/constant"
)

func ConvertVersion(version constant.AddressVersion) (byte, error) {
	switch version {
		case constant.AddressVersionMainnetSingleSignature:
			return 0, nil

		case constant.AddressVersionMainnetMultipleSignature:
			return 5, nil

		case constant.AddressVersionTestnetSingleSignature:
			return 111, nil

		case constant.AddressVersionTestnetMultipleSignature:
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
