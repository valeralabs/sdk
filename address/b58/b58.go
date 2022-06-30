package b58

import (
	"fmt"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/valeralabs/sdk/constant"
)

var versions = map[constant.AddressVersion]byte{
	constant.AddressVersionMainnetPublicKeyHash:   0,
	constant.AddressVersionMainnetScriptHash: 5,
	constant.AddressVersionTestnetPublicKeyHash:   111,
	constant.AddressVersionTestnetScriptHash: 196,
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

	return string(base58.CheckEncode(hash, converted)), nil
}
