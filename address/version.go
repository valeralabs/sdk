package address

import "errors"

type AddressVersion struct {
	HashMode HashMode
	Network  Network
}

type Network int

const (
	NetworkMainnet Network = iota
	NetworkTestnet
)

func (networkType Network) Check() bool {
	return networkType >= NetworkMainnet && networkType <= NetworkTestnet
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

func (version AddressVersion) BitcoinVersion() (byte, error) {
	switch version.Network {
	case NetworkMainnet:
		if version.HashMode == HashModeP2PKH {
			return byte(0), nil
		}

		return byte(5), nil

	case NetworkTestnet:
		if version.HashMode == HashModeP2PKH {
			return byte(111), nil
		}

		return byte(196), nil
	}

	return byte(0), errors.New("address version is invalid")
}

func (version AddressVersion) StacksVersion() (byte, error) {
	switch version.Network {
	case NetworkMainnet:
		if version.HashMode == HashModeP2PKH {
			return byte(22), nil
		}

		return byte(20), nil

	case NetworkTestnet:
		if version.HashMode == HashModeP2PKH {
			return byte(26), nil
		}

		return byte(21), nil
	}

	return byte(0), errors.New("address version is invalid")
}

func ReverseBitcoinVersion(raw byte) (AddressVersion, error) {
	switch raw {
	case 0:
		return AddressVersion{HashModeP2PKH, NetworkMainnet}, nil

	case 5:
		return AddressVersion{HashModeP2SH, NetworkMainnet}, nil

	case 111:
		return AddressVersion{HashModeP2PKH, NetworkTestnet}, nil

	case 196:
		return AddressVersion{HashModeP2SH, NetworkTestnet}, nil
	}

	return AddressVersion{}, errors.New("address version is invalid")
}

func ReverseStacksVersion(raw byte) (AddressVersion, error) {
	switch raw {
	case 22:
		return AddressVersion{HashModeP2PKH, NetworkMainnet}, nil

	case 20:
		return AddressVersion{HashModeP2SH, NetworkMainnet}, nil

	case 26:
		return AddressVersion{HashModeP2PKH, NetworkTestnet}, nil

	case 21:
		return AddressVersion{HashModeP2SH, NetworkTestnet}, nil

	default:
		return AddressVersion{}, errors.New("address version is invalid")
	}
}
