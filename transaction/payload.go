package transaction

import (
	"bytes"
	"encoding/binary"

	"github.com/linden/binstruct"

	"valera.co/vdk/address"
	"valera.co/vdk/constant"
	"valera.co/vdk/encoding/clarity"
)

//go:generate go run golang.org/x/tools/cmd/stringer@latest -type=PayloadType -output=payload_string.go

type PayloadType byte

const (
	PayloadTypeTokenTransfer PayloadType = iota
	PayloadTypeSmartContract
	PayloadTypeContractCall
	PayloadTypePoison
	PayloadTypeCoinbase
)

func (payload PayloadType) Check() bool {
	return payload >= PayloadTypeTokenTransfer && payload <= PayloadTypeCoinbase
}

type Payload interface {
	Marshal() ([]byte, error)
}

type PayloadTokenTransfer struct {
	Address address.Address
	Amount  int
	Memo    string
}

type PayloadSmartContract struct {
	Name string
	Body string
}

type PayloadContractCall struct {
	Address   address.Address
	Function  string
	Arguments clarity.List
}

func (payload PayloadTokenTransfer) Marshal() ([]byte, error) {
	var buffer bytes.Buffer

	writer := binstruct.NewWriter(&buffer, binary.BigEndian, false)

	writer.WriteUint8(uint8(PayloadTypeTokenTransfer))

	if payload.Address.Contract == "" {
		writer.WriteUint8(uint8(clarity.ClarityTypePrincipalStandard))
	} else {
		writer.WriteUint8(uint8(clarity.ClarityTypePrincipalContract))
	}

	princpal, err := clarity.EncodePrincipal(payload.Address)

	if err != nil {
		return []byte{}, err
	}

	writer.Write(princpal)

	writer.WriteUint64(uint64(payload.Amount))

	length := writer.Len() + constant.MemoLengthBytes + 2

	writer.Write([]byte(payload.Memo))

	for writer.Len() < length {
		writer.WriteByte(0x00)
	}

	return buffer.Bytes(), nil
}

func (payload PayloadSmartContract) Marshal() ([]byte, error) {
	var buffer bytes.Buffer

	writer := binstruct.NewWriter(&buffer, binary.BigEndian, false)

	writer.WriteUint8(uint8(PayloadTypeSmartContract))

	name := clarity.Value{
		Type:         clarity.ClarityTypeStringUTF8,
		Content:      []byte(payload.Name),
		PrefixLength: 1,
	}

	raw, err := name.Marshal(false)

	if err != nil {
		return []byte{}, err
	}

	writer.Write(raw)

	body := clarity.Value{
		Type:         clarity.ClarityTypeStringUTF8,
		Content:      []byte(payload.Body),
		PrefixLength: 4,
	}

	raw, err = body.Marshal(false)

	if err != nil {
		return []byte{}, err
	}

	writer.Write(raw)

	return buffer.Bytes(), nil
}

func (payload PayloadContractCall) Marshal() ([]byte, error) {
	var buffer bytes.Buffer

	writer := binstruct.NewWriter(&buffer, binary.BigEndian, false)

	writer.WriteUint8(uint8(PayloadTypeContractCall))

	principal, err := clarity.EncodePrincipal(payload.Address)

	if err != nil {
		return []byte{}, err
	}

	writer.Write(principal)

	function := clarity.Value{
		Type:         clarity.ClarityTypeStringUTF8,
		Content:      []byte(payload.Function),
		PrefixLength: 1,
	}

	raw, err := function.Marshal(false)

	if err != nil {
		return []byte{}, err
	}

	writer.Write(raw)

	raw, err = payload.Arguments.Marshal(true)

	if err != nil {
		return []byte{}, err
	}

	writer.Write(raw)

	return buffer.Bytes(), nil
}
