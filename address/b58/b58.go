package b58

import (
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/valeralabs/racks/constant"
	"github.com/valeralabs/racks/transaction"
)

func NewAddress(publics []transaction.PublicKey, version constant.AddressVersion, hash constant.HashMode) (string, error) {
	var prefix byte

	switch version {
	case constant.AddressVersionMainnetSingleSignature:
		prefix = 0

	case constant.AddressVersionMainnetMultipleSignature:
		prefix = 5

	case constant.AddressVersionTestnetSingleSignature:
		prefix = 111

	case constant.AddressVersionTestnetMultipleSignature:
		prefix = 196
	}

	if hash == constant.HashModeP2WPKH || hash == constant.HashModeP2WSH {
		return "", fmt.Errorf("TODO: add segwit support")
	}

	var address []byte

	if hash == constant.HashModeP2PKH {
		address = btcutil.Hash160(publics[0].Serialize())
	} else {
		return "", fmt.Errorf("TODO: add more address support support")
	}

	return string(base58.CheckEncode(address, prefix)), nil
}
