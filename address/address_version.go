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

func NewAddressVersionFromByte(versionByte byte) (AddressVersion, error) {
	var hashMode HashMode
	var addressNetworkType AddressNetworkType

	switch versionByte {
	case 0:
		hashMode = HashModeP2PKH
		addressNetworkType = AddressNetworkTypeMainnet
	case 5:
		hashMode = HashModeP2SH
		addressNetworkType = AddressNetworkTypeMainnet
	case 111:
		hashMode = HashModeP2PKH
		addressNetworkType = AddressNetworkTypeTestnet
	case 196:
		hashMode = HashModeP2SH
		addressNetworkType = AddressNetworkTypeTestnet
	case 22:
		hashMode = HashModeP2PKH
		addressNetworkType = AddressNetworkTypeMainnet
	case 20:
		hashMode = HashModeP2SH
		addressNetworkType = AddressNetworkTypeMainnet
	case 26:
		hashMode = HashModeP2PKH
		addressNetworkType = AddressNetworkTypeTestnet
	case 21:
		hashMode = HashModeP2SH
		addressNetworkType = AddressNetworkTypeTestnet
	}

	addressVersion, err := NewAddressVersion(hashMode, addressNetworkType)
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
