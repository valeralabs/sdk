package transaction

import (
	"fmt"

	"github.com/linden/binstruct"
	"github.com/linden/bite"
	"valera.co/vdk/address"
	"valera.co/vdk/encoding/clarity"
)

//go:generate go run golang.org/x/tools/cmd/stringer@latest -type=PostConditionMode,PostConditionType,FungibleConditionCode,NonFungibleConditionCode -output=postcondition_string.go

type PostCondition any

type PostConditionMode byte

const (
	PostConditionModeLoose PostConditionMode = iota + 1
	PostConditionModeStrict
)

func (mode PostConditionMode) Check() bool {
	return mode == PostConditionModeLoose || mode == PostConditionModeStrict
}

type PostConditionType byte

const (
	PostConditionTypeSTX PostConditionType = iota
	PostConditionTypeFT
	PostConditionTypeNFT
)

func (_type PostConditionType) Check() bool {
	return _type >= PostConditionTypeSTX && _type <= PostConditionTypeNFT
}

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

type PostConditionSTX struct {
	Condition FungibleConditionCode
	Principal PostConditionPrincipal
	Amount    uint64
}

type PostConditionFT struct {
	Condition FungibleConditionCode
	Principal PostConditionPrincipal
	Amount    uint64
	Asset     Asset
}

type PostConditionNFT struct {
	Condition NonFungibleConditionCode
	Principal PostConditionPrincipal
	Value     clarity.Value
	Asset     Asset
}

type PostConditionPrincipal struct {
	Address address.Address
	Type    PostConditionPrincipalType
}

func (principal PostConditionPrincipal) Encode(writer binstruct.Writer) {
	switch principal.Type {
	case PostConditionPrincipalTypeOrigin:
		writer.WriteUint8(uint8(PostConditionPrincipalTypeOrigin))

	case PostConditionPrincipalTypeStandard:
		writer.WriteUint8(uint8(PostConditionPrincipalTypeStandard))
		clarity.EncodePrincipal(principal.Address, writer)

	case PostConditionPrincipalTypeContract:
		writer.WriteUint8(uint8(PostConditionPrincipalTypeContract))
		clarity.EncodePrincipal(principal.Address, writer)
	}
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

func DecodePostConditionPrincipal(reader *bite.Reader, origin address.Address) (PostConditionPrincipal, error) {
	postConditionPrincipalType := PostConditionPrincipalType(reader.ReadSingle())

	switch postConditionPrincipalType {
	case PostConditionPrincipalTypeOrigin:
		return PostConditionPrincipal{
			Address: origin,
			Type:    PostConditionPrincipalTypeOrigin,
		}, nil

	case PostConditionPrincipalTypeStandard:
		principal, err := clarity.DecodePrincipal(clarity.ClarityTypePrincipalStandard, reader)

		if err != nil {
			return PostConditionPrincipal{}, err
		}

		return PostConditionPrincipal{
			Address: principal,
			Type:    PostConditionPrincipalTypeStandard,
		}, nil

	case PostConditionPrincipalTypeContract:
		principal, err := clarity.DecodePrincipal(clarity.ClarityTypePrincipalContract, reader)

		if err != nil {
			return PostConditionPrincipal{}, err
		}

		return PostConditionPrincipal{
			Address: principal,
			Type:    PostConditionPrincipalTypeContract,
		}, nil

	default:
		return PostConditionPrincipal{}, fmt.Errorf("post-condition principal type must be origin (0), standard (1) or contract (2), not %v", byte(postConditionPrincipalType))
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

func (asset Asset) Encode(writer binstruct.Writer) error {
	err := clarity.EncodePrincipal(asset.Address, writer)

	if err != nil {
		return fmt.Errorf("could not encode principal: %v", err)
	}

	writer.WriteUint8(uint8(len(asset.Name)))
	writer.Write([]byte(asset.Name))

	return nil
}
