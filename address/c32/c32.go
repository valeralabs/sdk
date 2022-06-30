package c32

import (
	"fmt"
	"encoding/hex"

	"github.com/valeralabs/sdk/constant"
	"github.com/valeralabs/sdk/encoding/c32"
)

var versions = map[constant.AddressVersion]byte{
	constant.AddressVersionMainnetSingleSignature:   22,
	constant.AddressVersionMainnetMultipleSignature: 20,
	constant.AddressVersionTestnetSingleSignature:   26,
	constant.AddressVersionTestnetMultipleSignature: 21,
}

func ConvertVersion(version constant.AddressVersion) (byte, error) {
	if value, ok := versions[version]; ok {
		return value, nil
	}

	return 0, fmt.Errorf("invalid version %d", version)
}

func ReverseVersion(version byte) (constant.AddressVersion, error) {
	for key, value := range versions {
		if value == version {
			return key, nil
		}
	}

	return constant.AddressVersion(0), fmt.Errorf("invalid version %d", version)
}

func Encode(hash []byte, version constant.AddressVersion) (string, error) {
	converted, err := ConvertVersion(version)

	if err != nil {
		return "", err
	}

	encodedHash := make([]byte, hex.EncodedLen(len(hash)))
	hex.Encode(encodedHash, hash)

	encoded, err := c32.ChecksumEncode(encodedHash, converted)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("S%s", encoded), nil
}
