package address

import (
	"fmt"
)

type AddressVersion struct {
	hashMode HashMode
	networkType AddressNetworkType
}

func NewAddressVersion(hashMode HashMode, networkType AddressNetworkType) (AddressVersion, error) {
	if !hashMode.Check() {
		return AddressVersion{}, fmt.Errorf("hashMode must be between 0 and 1, not %v", hashMode)
	}

	if !networkType.Check() {
		return AddressVersion{}, fmt.Errorf("network must be between 0 and 1, not %v", networkType)
	}

	return AddressVersion {
		hashMode: hashMode,
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
			switch addressVersion.hashMode {
			case HashModeP2PKH:
				return byte(0), nil
			default:
				return byte(5), nil
			}

		case AddressNetworkStacks:
			switch addressVersion.hashMode {
			case HashModeP2PKH:
				return byte(22), nil
			default:
				return byte(20), nil
			}
		}
	case AddressNetworkTypeTestnet:
		switch addressNetwork {
		case AddressNetworkBitcoin:
			switch addressVersion.hashMode {
			case HashModeP2PKH:
				return byte(111), nil
			default:
				return byte(196), nil
			}

		case AddressNetworkStacks:
			switch addressVersion.hashMode {
			case HashModeP2PKH:
				return byte(26), nil
			default:
				return byte(21), nil
			}
		}
	}

	return byte(0), fmt.Errorf("AddressNetworkType is invalid")
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
