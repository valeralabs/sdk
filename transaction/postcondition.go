package transaction

import (
	"fmt"

	"github.com/linden/bite"
	"github.com/valeralabs/sdk/address"
	"github.com/valeralabs/sdk/constant"
	"github.com/valeralabs/sdk/encoding/clarity"
)

type PostCondition any
type FungibleConditionCode int

const (
	FungibleConditionCodeSentEqual FungibleConditionCode = iota + 1
	FungibleConditionCodeSentGreaterThan
	FungibleConditionCodeSentGreaterOrEqual
	FungibleConditionCodeSentLessThen
	FungibleConditionCodeSentLessThenOrEqual
)

func (code FungibleConditionCode) Check() bool {
	return code >= FungibleConditionCodeSentEqual && code <= FungibleConditionCodeSentLessThenOrEqual
}

type NonFungibleConditionCode int

const (
	NonFungibleConditionCodeSent NonFungibleConditionCode = 0x10 + iota
	NonFungibleConditionCodeNotSent
)

func (code NonFungibleConditionCode) Check() bool {
	return code == NonFungibleConditionCodeSent || code == NonFungibleConditionCodeNotSent
}

type Asset struct {
	Address address.Address
	Name    string
}

type PostConditionStacks struct {
	Condition FungibleConditionCode
	Principal address.Address
	Amount    uint64
}

type PostConditionFungible struct {
	Condition FungibleConditionCode
	Principal address.Address
	Amount    uint64
	Asset     Asset
}

type PostConditionNonFungible struct {
	Condition NonFungibleConditionCode
	Principal address.Address
	Value     clarity.Value
	Asset     Asset
}

type PostConditionPrincipalType byte

const (
	PostConditionPrincipalTypeOrigin PostConditionPrincipalType = iota + 1
	PostConditionPrincipalTypeStandard
	PostConditionPrincipalTypeContract
)

func (postConditionPrincipalType PostConditionPrincipalType) Check() bool {
	return postConditionPrincipalType >= PostConditionPrincipalTypeOrigin && postConditionPrincipalType <= PostConditionPrincipalTypeContract
}

func DecodePostConditionPrincipal(reader *bite.Reader, origin address.Address) (address.Address, error) {
	postConditionPrincipalType := PostConditionPrincipalType(reader.ReadSingle())

	switch postConditionPrincipalType {
	case PostConditionPrincipalTypeOrigin:
		return origin, nil

	case PostConditionPrincipalTypeStandard:
		principal, err := clarity.DecodePrincipal(clarity.ClarityTypePrincipalStandard, reader)

		if err != nil {
			return address.Address{}, err
		}

		return principal, nil

	case PostConditionPrincipalTypeContract:
		principal, err := clarity.DecodePrincipal(clarity.ClarityTypePrincipalContract, reader)

		if err != nil {
			return address.Address{}, err
		}

		return principal, nil

	default:
		return address.Address{}, fmt.Errorf("post-condition principal type must be origin, standard or contract")
	}
}

func DecodeAsset(reader *bite.Reader) (Asset, error) {
	var asset Asset
	var err error

	asset.Address, err = clarity.DecodePrincipal(clarity.ClarityTypePrincipalContract, reader)

	if err != nil {
		return Asset{}, fmt.Errorf("address is not valid")
	}

	nameLength := int(reader.ReadSingle())
	asset.Name = string(reader.Read(nameLength))

	return asset, nil
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
