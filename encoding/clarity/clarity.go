package clarity

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/linden/binstruct"
	"github.com/linden/bite"
	"github.com/valeralabs/sdk/address"
	"github.com/valeralabs/sdk/constant"
)

type ClarityType byte

const (
	ClarityTypeInt ClarityType = iota
	ClarityTypeUInt
	ClarityTypeBuffer
	ClarityTypeBoolTrue
	ClarityTypeBoolFalse
	ClarityTypePrincipalStandard
	ClarityTypePrincipalContract
	ClarityTypeResponseOk
	ClarityTypeResponseErr
	ClarityTypeOptionalNone
	ClarityTypeOptionalSome
	ClarityTypeList
	ClarityTypeTuple
	ClarityTypeStringASCII
	ClarityTypeStringUTF8
)

func (_type ClarityType) Check() bool {
	return _type >= ClarityTypeInt && _type <= ClarityTypeStringUTF8
}

type Value struct {
	Type         ClarityType
	Content      []byte
	PrefixLength int
}

type List struct {
	Type            ClarityType
	Content         []Value
	PrefixLength    int
	SubPrefixLength int
}

func createLengthPrefix(length int, total int) []byte {
	buffer := new(bytes.Buffer)

	binary.Write(buffer, binary.BigEndian, int8(length))

	padding := bytes.Repeat([]byte{0}, total-buffer.Len())

	return append(padding, buffer.Bytes()...)
}

func readLengthPrefix(reader *bite.Reader, total int) int {
	var length int

	for _, cursor := range reader.Read(total) {
		if cursor != 0 {
			length = (length * 10) + int(cursor)
		}
	}

	return length
}

func (cursor *Value) Unmarshal(raw []byte, typed bool) error {
	reader := bite.NewReader(raw)

	if cursor.PrefixLength == 0 {
		cursor.PrefixLength = constant.DefaultPrefixLength
	}

	if typed == true {
		cursor.Type = ClarityType(reader.ReadSingle())
	}

	length := readLengthPrefix(&reader, cursor.PrefixLength)
	cursor.Content = reader.Read(length)

	return nil
}

func (cursor Value) Marshal(typed bool) ([]byte, error) {
	if len(cursor.Content) > constant.MaxStringLength {
		return []byte{}, errors.New("string is above the max length")
	}

	var buffer []byte

	if typed == true {
		buffer = append(buffer, byte(cursor.Type))
	}

	buffer = append(buffer, createLengthPrefix(len(cursor.Content), cursor.PrefixLength)...)
	buffer = append(buffer, []byte(cursor.Content)...)

	return buffer, nil
}

func (cursor Value) Length(typed bool) int {
	length := len(cursor.Content) + cursor.PrefixLength

	if typed {
		length += 1
	}

	return length
}

func NewValue(from []byte) Value {
	if len(from) > constant.MaxStringLength {
		panic("string is longer then the max length")
	}

	return Value{
		Content:      from,
		PrefixLength: constant.DefaultPrefixLength,
	}
}

func (list *List) Unmarshal(data []byte, typed bool) error {
	reader := bite.NewReader(data)

	if list.PrefixLength == 0 {
		list.PrefixLength = constant.DefaultPrefixLength
	}

	if list.SubPrefixLength == 0 {
		list.SubPrefixLength = constant.DefaultPrefixLength
	}

	length := readLengthPrefix(&reader, list.PrefixLength)

	for index := 0; index < length; index++ {
		decoded := Value{
			PrefixLength: list.SubPrefixLength,
		}

		err := decoded.Unmarshal(reader.Value, typed)

		if err != nil {
			return err
		}

		list.Content = append(list.Content, decoded)

		reader.Read(decoded.PrefixLength + len(decoded.Content))

		if typed == true {
			reader.Read(1)
		}
	}

	return nil
}

func (list List) Marshal(typed bool) ([]byte, error) {
	buffer := createLengthPrefix(len(list.Content), list.PrefixLength)

	for _, encoded := range list.Content {
		decoded, err := encoded.Marshal(typed)

		if err != nil {
			return []byte{}, err
		}

		buffer = append(buffer, decoded...)
	}

	return buffer, nil
}

func NewList(raw [][]byte) List {
	var list List

	list.PrefixLength = constant.DefaultPrefixLength
	list.SubPrefixLength = constant.DefaultPrefixLength

	for _, cursor := range raw {
		list.Content = append(list.Content, NewValue(cursor))
	}

	return list
}

func DecodePrincipal(from ClarityType, reader *bite.Reader) (address.Address, error) {
	if from != ClarityTypePrincipalStandard && from != ClarityTypePrincipalContract {
		return address.Address{}, errors.New("clarity type can only be standard or contract")
	}

	version, err := address.ReverseStacksVersion(reader.ReadSingle())

	if err != nil {
		return address.Address{}, errors.New("address version is invalid")
	}

	hash := reader.Read(20)

	var contract string

	if from == ClarityTypePrincipalContract {
		name := Value{
			PrefixLength: 1,
		}

		err = name.Unmarshal(reader.Value, false)

		if err != nil {
			return address.Address{}, errors.New("invalid contract name")
		}

		contract = string(name.Content)

		reader.Read(len(name.Content) + 1)
	}

	return address.Address{version, hash, contract}, nil
}

func EncodePrincipal(from address.Address, writer binstruct.Writer) error {
	version, err := from.Version.StacksVersion()

	if err != nil {
		return err
	}

	writer.WriteUint8(uint8(version))
	writer.Write(from.Hash)

	if from.Contract != "" {
		name := Value{
			Content:      []byte(from.Contract),
			PrefixLength: 1,
		}

		raw, err := name.Marshal(false)

		if err != nil {
			return err
		}

		writer.Write(raw)
	}

	return nil
}
