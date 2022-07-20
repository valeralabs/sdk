package VDK

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"valera.co/vdk/address"
	"valera.co/vdk/constant"
	"valera.co/vdk/encoding/c32"
	"valera.co/vdk/encoding/clarity"
)

// Wrapper around [address.Address].
type Principal struct {
	value *address.Address
}

// Wrapper around [encoding/clarity.Value].
type ClarityValue struct {
	value *clarity.Value
}

// Wrapper around [encoding/clarity.List].
type ClarityList struct {
	value *clarity.List
}

// FT or NFT
type Asset struct {
	Contract *Principal
	Name     string
}

// Derive a Stacks address (C32) from a [Principal].
func (principal *Principal) Stacks() (string, error) {
	c32, err := (*principal.value).C32()

	if err != nil {
		return "", err
	}

	return c32, nil
}

// Derive a bitcoin address (B58) from a [Principal].
func (principal *Principal) Bitcoin() (string, error) {
	b58, err := (*principal.value).B58()

	if err != nil {
		return "", err
	}

	return b58, nil
}

// Create a principal from a Stacks (c32) address.
func NewPrincipal(value string) (*Principal, error) {
	var principal address.Address
	var err error

	split := strings.Split(value, ".")
	value = split[0]

	if len(split) > 1 {
		principal.Contract = split[1]
	}

	raw := []byte(value)
	raw = raw[1:]

	var version byte

	hexed, version, err := c32.ChecksumDecode(raw)

	hash := make([]byte, hex.DecodedLen(len(hexed)))
	hex.Decode(hash, hexed)

	principal.Hash = hash

	if err != nil {
		return &Principal{}, err
	}

	principal.Version, err = address.ReverseStacksVersion(version)

	if err != nil {
		return &Principal{}, err
	}

	return &Principal{&principal}, nil
}

func (value *ClarityValue) Raw() string {
	return fmt.Sprintf("%#+v\n", *value.value)
}

func (value *ClarityValue) SetPrefix(prefix int) {
	(*value.value).PrefixLength = prefix
}

// `base`:
// 	  0  is Int
// 	  1  is UInt
// 	  2  is Buffer
// 	  3  is BoolTrue
// 	  4  is BoolFalse
// 	  5  is PrincipalStandard
// 	  6  is PrincipalContract
// 	  7  is ResponseOk
// 	  8  is ResponseErr
// 	  9  is OptionalNone
// 	  10 is OptionalSome
// 	  11 is List
// 	  12 is Tuple
// 	  13 is StringASCII
// 	  14 is StringUTF8
// `content`: value
func NewClarityValue(base int, content string) (*ClarityValue, error) {
	typed := clarity.ClarityType(base)

	if typed.Check() == false {
		return &ClarityValue{}, errors.New("base is not a clarity type")
	}

	return &ClarityValue{&clarity.Value{
		Type:         typed,
		Content:      []byte(content),
		PrefixLength: constant.DefaultPrefixLength,
	}}, nil
}

func (list *ClarityList) SetPrefix(prefix int) {
	(*list.value).PrefixLength = prefix
}

func (list *ClarityList) SetSubPrefix(prefix int) {
	(*list.value).SubPrefixLength = prefix
}

func (list *ClarityList) Add(value *ClarityValue) {
	list.value.Content = append(list.value.Content, *value.value)
}

func (list *ClarityList) Raw() string {
	return fmt.Sprintf("%#+v\n", *list.value)
}

func NewClarityList() *ClarityList {
	return &ClarityList{&clarity.List{
		PrefixLength:    constant.DefaultPrefixLength,
		SubPrefixLength: constant.DefaultPrefixLength,
	}}
}

// Get the `Asset` as a string.
func (asset *Asset) String() string {
	stacks, _ := (*asset.Contract.value).C32()
	return stacks + "::" + asset.Name
}

// Create a new Asset (FT or NFT)
func NewAsset(contract *Principal, name string) (*Asset, error) {
	if contract == nil {
		return &Asset{}, errors.New("contract is nil")
	}

	if (*contract.value).Contract == "" {
		return &Asset{}, errors.New("contract is a standard principal")
	}

	if name == "" {
		return &Asset{}, errors.New("asset name is invalid")
	}

	return &Asset{
		Contract: contract,
		Name:     name,
	}, nil
}
