package transaction

import (
	"fmt"

	"github.com/linden/bite"
	"github.com/valeralabs/sdk/address/c32"
	"github.com/valeralabs/sdk/address"
	"github.com/valeralabs/sdk/constant"
	"github.com/valeralabs/sdk/encoding/clarity"
)

type PostCondition any

type FungibleConditionCode int

const (
	FungibleConditionCodeSentEq FungibleConditionCode = iota + 1
	FungibleConditionCodeSentGt
    FungibleConditionCodeSentGe
    FungibleConditionCodeSentLt
    FungibleConditionCodeSentLe
)

func (code FungibleConditionCode) Check() bool {
	return code >= FungibleConditionCodeSentEq && code <= FungibleConditionCodeSentLe
}

type NonFungibleConditionCode int

const (
	NonFungibleConditionCodeSent NonFungibleConditionCode = iota + 0x10
	NonFungibleConditionCodeNotSent
)

type AssetInfo struct {
	contractAddress address.Address
	contractName string
	assetName string
}

type PostConditionSTX struct {
	postConditionPrincipal address.Address
	fungibleConditionCode FungibleConditionCode
	amount uint64
}

type PostConditionFungible struct {
	postConditionPrincipal address.Address
	assetInfo AssetInfo
	fungibleConditionCode FungibleConditionCode
	amount uint64
}

type PostConditionNonFungible struct {
	postConditionPrincipal address.Address
	assetInfo AssetInfo
	value uint64
	nonFungibleConditionCode NonFungibleConditionCode
}

type PostConditionPrincipalType byte

const (
	PostConditionPrincipalTypeOrigin PostConditionPrincipalType = 0x01
	PostConditionPrincipalTypeStandard PostConditionPrincipalType = 0x02
	PostConditionPrincipalTypeContract PostConditionPrincipalType = 0x03
)

func (postConditionPrincipalType PostConditionPrincipalType) Check() bool {
	return postConditionPrincipalType >= PostConditionPrincipalTypeOrigin && postConditionPrincipalType <= PostConditionPrincipalTypeContract
}

 func DeserializePostConditionPrincipal(reader *bite.Reader, originAddress address.Address) (address.Address, error) {
	postConditionPrincipalType := PostConditionPrincipalType(reader.ReadSingle()) // -1, because rust enums start at 1

	if !postConditionPrincipalType.Check() {
		return address.Address{}, fmt.Errorf("post condition principal type must be between or equal to 0 and 2, not %v", byte(postConditionPrincipalType))
	}

	if postConditionPrincipalType == PostConditionPrincipalTypeOrigin {
		return originAddress, nil
	} else if postConditionPrincipalType == PostConditionPrincipalTypeStandard {
		version, err := c32.ReverseVersion(reader.ReadSingle())

		if err != nil {
			return address.Address{}, fmt.Errorf("Could not deserialize post condition principal address: %v", err)
		}

		addressHash := reader.Read(20)

		principalAddress := address.NewAddressFromPublicKeyHash(addressHash, version)

		return principalAddress, nil
	} else if postConditionPrincipalType == PostConditionPrincipalTypeContract {
		principalAddress, err := clarity.DecodePrincipal(clarity.ClarityTypePrincipalContract, reader)
		if err != nil {
			return address.Address{}, err
		}

		return principalAddress, nil
	}

	return address.Address{}, fmt.Errorf("Could not parse post condition principal")
 }

 func HashModeToAddressVersion(hashMode constant.HashMode, transactionVersion constant.TransactionVersion) constant.AddressVersion {
	switch transactionVersion {
	case constant.TransactionVersionMainnet:
		switch hashMode {
		case constant.HashModeP2PKH:
			return constant.AddressVersionMainnetPublicKeyHash
		default:
			return constant.AddressVersionMainnetScriptHash
		}
	case constant.TransactionVersionTestnet:
		switch hashMode {
		case constant.HashModeP2PKH:
			return constant.AddressVersionTestnetPublicKeyHash
		default:
			return constant.AddressVersionTestnetScriptHash
		}
	}

	return constant.AddressVersionMainnetPublicKeyHash
 }
