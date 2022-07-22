package clarity

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"

	"github.com/linden/binstruct"
	"github.com/linden/bite"
	"valera.co/vdk/address"
	"valera.co/vdk/constant"
)

//go:generate go run golang.org/x/tools/cmd/stringer@latest -type=ClarityType -output=clarity_string.go

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
	Content      any
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
	if cursor.PrefixLength == 0 {
		cursor.PrefixLength = constant.DefaultPrefixLength
	}

	reader := bite.NewReader(raw)

	if cursor.Type == ClarityTypeBoolTrue || cursor.Type == ClarityTypeBoolFalse || cursor.Type == ClarityTypeOptionalNone {
		typed = true
	}

	if typed == true {
		cursor.Type = ClarityType(reader.ReadSingle())
	}

	switch cursor.Type {
	case ClarityTypeInt, ClarityTypeUInt:
		if len(reader.Value) < 16 {
			return errors.New("value is an integer but is not long enough")
		}

		low := int64(binary.BigEndian.Uint64(reader.Read(8)))
		high := int64(binary.BigEndian.Uint64(reader.Read(8)))

		value := big.NewInt(low)
		value.Mul(value, big.NewInt(100000000))
		value.Add(value, big.NewInt(high))

		cursor.Content = value

		return nil
	}

	length := readLengthPrefix(&reader, cursor.PrefixLength)
	cursor.Content = reader.Read(length)

	return nil
}

func (cursor *Value) Check() bool {
	switch cursor.Type {
	case ClarityTypeInt, ClarityTypeUInt:
		_, isInt := cursor.Content.(int)
		_, isBig := cursor.Content.(big.Int)

		return isInt || isBig

	case ClarityTypePrincipalStandard, ClarityTypePrincipalContract:
		_, ok := cursor.Content.(address.Address)
		return ok

	case ClarityTypeResponseOk, ClarityTypeResponseErr, ClarityTypeOptionalSome:
		_, ok := cursor.Content.(Value)
		return ok

	case ClarityTypeStringASCII, ClarityTypeStringUTF8, ClarityTypeBuffer:
		_, ok := cursor.Content.([]byte)
		return ok

	}

	return true
}

func (cursor *Value) Marshal(typed bool) ([]byte, error) {
	if cursor.Check() == false {
		return []byte{}, fmt.Errorf("content/type miss-match got \"%T\" from %+v", cursor.Content, cursor)
	}

	var buffer []byte

	if cursor.Type == ClarityTypeBoolTrue || cursor.Type == ClarityTypeBoolFalse || cursor.Type == ClarityTypeOptionalNone {
		typed = true
	}

	if typed == true {
		buffer = append(buffer, byte(cursor.Type))
	}

	switch cursor.Type {
	case ClarityTypeBoolTrue, ClarityTypeBoolFalse, ClarityTypeOptionalNone:
		return buffer, nil

	case ClarityTypeInt, ClarityTypeUInt:
		var content *big.Int

		switch value := cursor.Content.(type) {
		case int:
			content = big.NewInt(int64(value))

		case *big.Int:
			content = value
		}

		value := content.Bytes()

		for len(value) < 16 {
			value = append([]byte{0}, value...)
		}

		buffer = append(buffer, value...)

		return buffer, nil

	case ClarityTypePrincipalStandard, ClarityTypePrincipalContract:
		principal, err := EncodePrincipal(cursor.Content.(address.Address))

		if err != nil {
			return []byte{}, err
		}

		buffer = append(buffer, principal...)

		return buffer, nil

	case ClarityTypeResponseOk, ClarityTypeResponseErr, ClarityTypeOptionalSome:
		content := cursor.Content.(Value)

		underlying, err := content.Marshal(true)

		if err != nil {
			return []byte{}, err
		}

		buffer = append(buffer, underlying...)

		return buffer, nil

	case ClarityTypeStringASCII, ClarityTypeStringUTF8, ClarityTypeBuffer:
		content := cursor.Content.([]byte)

		if cursor.Type == ClarityTypeStringASCII {
			for index, _ := range content {
				content[index] = content[index] & 255
			}
		}

		if len(content) > constant.MaxStringLength {
			return []byte{}, errors.New("string is above the max length")
		}

		buffer = append(buffer, createLengthPrefix(len(cursor.Content.([]byte)), cursor.PrefixLength)...)
		buffer = append(buffer, content...)

		return buffer, nil

	case ClarityTypeList, ClarityTypeTuple:
		return []byte{}, errors.New("list and tuple types are not yet supported")
	}

	return []byte{}, errors.New("type is invalid")
}

func (cursor Value) Length(typed bool) int {
	length := len(cursor.Content.([]byte)) + cursor.PrefixLength

	if typed {
		length += 1
	}

	return length
}

func NewValue(from any, base ClarityType) Value {
	return Value{
		Type:         base,
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
			Type:         list.Type,
			PrefixLength: list.SubPrefixLength,
		}

		err := decoded.Unmarshal(reader.Value, typed)

		if err != nil {
			return err
		}

		list.Content = append(list.Content, decoded)

		reader.Read(decoded.PrefixLength + len(decoded.Content.([]byte)))

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

func NewList(raw [][]byte, base ClarityType) List {
	var list List

	list.Type = base
	list.PrefixLength = constant.DefaultPrefixLength
	list.SubPrefixLength = constant.DefaultPrefixLength

	for _, cursor := range raw {
		list.Content = append(list.Content, NewValue(cursor, list.Type))
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
			Type:         ClarityTypeStringUTF8,
			PrefixLength: 1,
		}

		err = name.Unmarshal(reader.Value, false)

		if err != nil {
			return address.Address{}, errors.New("invalid contract name")
		}

		contract = string(name.Content.([]byte))

		reader.Read(len(name.Content.([]byte)) + 1)
	}

	return address.Address{version, hash, contract}, nil
}

func EncodePrincipal(from address.Address) ([]byte, error) {
	var buffer bytes.Buffer

	writer := binstruct.NewWriter(&buffer, binary.BigEndian, false)

	version, err := from.Version.StacksVersion()

	if err != nil {
		return []byte{}, err
	}

	writer.WriteUint8(uint8(version))
	writer.Write(from.Hash)

	if from.Contract != "" {
		name := Value{
			Type:         ClarityTypeStringUTF8,
			Content:      []byte(from.Contract),
			PrefixLength: 1,
		}

		raw, err := name.Marshal(false)

		if err != nil {
			return []byte{}, err
		}

		writer.Write(raw)
	}

	return buffer.Bytes(), nil
}
