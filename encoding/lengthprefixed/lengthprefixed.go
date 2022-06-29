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

func (cursor *String) Unmarshal(raw []byte) error {
	reader := bite.New(raw)

	if cursor.PrefixLength == 0 {
		cursor.PrefixLength = constant.DefaultPrefixLength
	}

	var length int

	for _, cursor := range reader.Read(cursor.PrefixLength) {
		if cursor != 0 {
			length = (length * 10) + int(cursor)
		}
	}

	cursor.Content = reader.Read(length)

	return nil
}

func (cursor *String) Marshal() ([]byte, error) {
	if len(cursor.Content) > constant.MaxStringLength {
		return []byte{}, errors.New("string is above the max length")
	}

	var buffer []byte

	prefixed := prefix(len(cursor.Content), cursor.PrefixLength)

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

func prefix(length int, total int) []byte {
	buffer := new(bytes.Buffer)

	binary.Write(buffer, binary.BigEndian, int8(length))

	padding := bytes.Repeat([]byte{0}, total-buffer.Len())

	return append(padding, buffer.Bytes()...)
}

// type List struct {
// 	Content []String
// 	Length  int64
// 	Max     int64
// }
//
// func (list *List) Unmarshal(data []byte) error {
// 	return nil
// }
//
// func (list *List) Marshal() ([]byte, error) {
// 	return []byte{}, nil
// }
