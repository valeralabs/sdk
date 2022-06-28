package lengthprefixed

import (
	"github.com/valeralabs/sdk/constant"

	"errors"
	"strconv"
	"strings"
)

type String struct {
	Content      []byte
	PrefixLength int
	Max          int
}

func (cursor *String) Unmarshal(data []byte) error {
	panic("TODO: Unmarshal LP string")
	return nil
}

func (cursor *String) Marshal() ([]byte, error) {
	if len(cursor.Content) > cursor.Max {
		return []byte{}, errors.New("String is above the Max length")
	}

	var buffer []byte

	prefixed := prefix(len(cursor.Content), cursor.PrefixLength)

	buffer = append(buffer, prefixed...)
	buffer = append(buffer, []byte(cursor.Content)...)

	var result []byte

	// TODO: probably a better way to do this
	for _, cursor := range buffer {
		result = append(result, []byte(strconv.Itoa(int(cursor)))...)
	}

	return result, nil
}

func CreateString(from []byte) String {
	if len(from) > constant.MaxStringLength {
		panic("String is above the Max length")
	}

	return String{
		Content:      from,
		PrefixLength: constant.DefaultPrefixLength,
		Max:          constant.MaxStringLength,
	}
}

// TODO: clean up
func prefix(length int, bytes int) []byte {
	encoded := strconv.FormatInt(int64(length), 16)

	var padding string

	if len(encoded) < bytes*2 {
		padding = strings.Repeat("0", ((bytes * 2) - len(encoded)))
	}

	combined := padding + encoded

	var result []byte

	for index := 0; index < len(combined); index += 2 {
		size, _ := strconv.ParseInt(combined[index:index+2], 16, 64)

		result = append(result, byte(size))
	}

	return result
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
