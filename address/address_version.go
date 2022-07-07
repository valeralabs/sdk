package address

import (
	"fmt"
)

type AddressVersion struct {
	addressType AddressType
	networkType AddressNetworkType
}

func NewAddressVersion(addressType AddressType, networkType AddressNetworkType) (AddressVersion, error) {
	if !addressType.Check() {
		return AddressVersion{}, fmt.Errorf("hashMode must be between 0 and 1, not %v", addressType)
	}

	if !networkType.Check() {
		return AddressVersion{}, fmt.Errorf("network must be between 0 and 1, not %v", networkType)
	}

	return AddressVersion {
		addressType: addressType,
		networkType: networkType,
	}, nil
}

func (addressVersion AddressVersion) GetVersionByte(addressNetwork AddressNetwork) (byte, error) {
	if !addressVersion.networkType.Check() {
		return byte(0), fmt.Errorf("AddressNetworkType must be 0 or 1, not %v", addressVersion.networkType)
	}

	switch addressVersion.networkType {
	case AddressNetworkTypeMainnet:
		switch addressNetwork {
		case AddressNetworkBitcoin:
			switch addressVersion.addressType {
			case AddressTypeP2PKH:
				return byte(0), nil
			case AddressTypeP2SH:
				return byte(5), nil
			}

		case AddressNetworkStacks:
			switch addressVersion.addressType {
			case AddressTypeP2PKH:
				return byte(22), nil
			case AddressTypeP2SH:
				return byte(20), nil
			}
		}
	case AddressNetworkTypeTestnet:
		switch addressNetwork {
		case AddressNetworkBitcoin:
			switch addressVersion.addressType {
			case AddressTypeP2PKH:
				return byte(111), nil
			case AddressTypeP2SH:
				return byte(196), nil
			}

		case AddressNetworkStacks:
			switch addressVersion.addressType {
			case AddressTypeP2PKH:
				return byte(26), nil
			case AddressTypeP2SH:
				return byte(21), nil
			}
		}
	}

	return byte(0), fmt.Errorf("AddressNetworkType is invalid")
}

func NewAddressVersionFromByte(versionByte byte) (AddressVersion, error) {
	var addressType AddressType
	var addressNetworkType AddressNetworkType

	switch versionByte {
	case 0:
		addressType = AddressTypeP2PKH
		addressNetworkType = AddressNetworkTypeMainnet
	case 5:
		addressType = AddressTypeP2SH
		addressNetworkType = AddressNetworkTypeMainnet
	case 111:
		addressType = AddressTypeP2PKH
		addressNetworkType = AddressNetworkTypeTestnet
	case 196:
		addressType = AddressTypeP2SH
		addressNetworkType = AddressNetworkTypeTestnet
	case 22:
		addressType = AddressTypeP2PKH
		addressNetworkType = AddressNetworkTypeMainnet
	case 20:
		addressType = AddressTypeP2SH
		addressNetworkType = AddressNetworkTypeMainnet
	case 26:
		addressType = AddressTypeP2PKH
		addressNetworkType = AddressNetworkTypeTestnet
	case 21:
		addressType = AddressTypeP2SH
		addressNetworkType = AddressNetworkTypeTestnet
	}

	addressVersion, err := NewAddressVersion(addressType, addressNetworkType)
	if err != nil {
		return AddressVersion{}, fmt.Errorf("Could not create address version: %v", err)
	}

	return addressVersion, nil
}

type AddressNetwork int

const (
	AddressNetworkBitcoin AddressNetwork = iota
	AddressNetworkStacks
)

func (addressNetwork AddressNetwork) Check() bool {
	return addressNetwork >= AddressNetworkBitcoin && addressNetwork <= AddressNetworkStacks
}

type AddressNetworkType int

const (
	AddressNetworkTypeMainnet AddressNetworkType = iota
	AddressNetworkTypeTestnet
)

func (networkType AddressNetworkType) Check() bool {
	return networkType >= AddressNetworkTypeMainnet && networkType <= AddressNetworkTypeTestnet
}

type HashMode int

const (
	HashModeP2PKH HashMode = iota
	HashModeP2SH
	HashModeP2WPKH
	HashModeP2WSH
)

func (hashMode HashMode) Check() bool {
	return hashMode >= HashModeP2PKH && hashMode <= HashModeP2WSH
}

type AddressType int

const (
	AddressTypeP2PKH AddressType = iota
	AddressTypeP2SH
)

func (addressType AddressType) Check() bool {
	return addressType >= AddressTypeP2PKH && addressType <= AddressTypeP2SH
}
