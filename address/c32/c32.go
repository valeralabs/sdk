package c32

import (
	"fmt"

	"github.com/valeralabs/racks/constant"
	"github.com/valeralabs/racks/encoding/c32"
)

func ConvertVersion(version constant.AddressVersion) (byte, error) {
	switch version {
	case constant.AddressVersionMainnetSingleSignature:
		return 22, nil

	case constant.AddressVersionMainnetMultipleSignature:
		return 20, nil

	case constant.AddressVersionTestnetSingleSignature:
		return 26, nil

	case constant.AddressVersionTestnetMultipleSignature:
		return 21, nil

	default:
		return 0, fmt.Errorf("invalid version %d", version)
	}
}

func Encode(hash []byte, version constant.AddressVersion) (string, error) {
	converted, err := ConvertVersion(version)

	if err != nil {
		return "", err
	}

	encoded, err := c32.ChecksumEncode(hash, converted)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("S%s", encoded), nil
}
