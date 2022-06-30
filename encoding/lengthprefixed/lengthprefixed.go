package lengthprefixed

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/linden/bite"
	"github.com/valeralabs/sdk/constant"
)

type String struct {
	Content      []byte
	PrefixLength int
}

type List struct {
	Content         []String
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

func (cursor *String) Unmarshal(raw []byte) error {
	reader := bite.New(raw)

	if cursor.PrefixLength == 0 {
		cursor.PrefixLength = constant.DefaultPrefixLength
	}

	length := readLengthPrefix(&reader, cursor.PrefixLength)
	cursor.Content = reader.Read(length)

	return nil
}

func (cursor String) Marshal() ([]byte, error) {
	if len(cursor.Content) > constant.MaxStringLength {
		return []byte{}, errors.New("string is above the max length")
	}

	var buffer []byte

	prefixed := createLengthPrefix(len(cursor.Content), cursor.PrefixLength)

	buffer = append(buffer, prefixed...)
	buffer = append(buffer, []byte(cursor.Content)...)

	return buffer, nil
}

func NewString(from []byte) String {
	if len(from) > constant.MaxStringLength {
		panic("string is longer then the max length")
	}

	return String{
		Content:      from,
		PrefixLength: constant.DefaultPrefixLength,
	}
}

func (list *List) Unmarshal(data []byte) error {
	reader := bite.New(data)

	if list.PrefixLength == 0 {
		list.PrefixLength = constant.DefaultPrefixLength
	}

	if list.SubPrefixLength == 0 {
		list.SubPrefixLength = constant.DefaultPrefixLength
	}

	length := readLengthPrefix(&reader, list.PrefixLength)

	for index := 0; index < length; index++ {
		decoded := String{
			PrefixLength: list.SubPrefixLength,
		}

		err := decoded.Unmarshal(reader.Value)

		if err != nil {
			return err
		}

		list.Content = append(list.Content, decoded)

		reader.Read(decoded.PrefixLength + len(decoded.Content))
	}

	return nil
}

func (list List) Marshal() ([]byte, error) {
	buffer := createLengthPrefix(len(list.Content), list.PrefixLength)

	for _, encoded := range list.Content {
		decoded, err := encoded.Marshal()

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
		list.Content = append(list.Content, NewString(cursor))
	}

	return list
}
