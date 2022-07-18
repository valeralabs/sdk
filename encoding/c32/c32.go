package c32

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var C32Alphabet = []byte("0123456789ABCDEFGHJKMNPQRSTVWXYZ")
var HexAlphabet = []byte("0123456789abcdef")

var HexPattern = regexp.MustCompile(`^[0-9a-fA-F]*$`)
var C32Pattern = regexp.MustCompile(`^[0-9A-Z]*$`)

var ZeroPattern = regexp.MustCompile(`(0)`)
var OnePattern = regexp.MustCompile(`(L|I)`)

func hash(source []byte) []byte {
	sha := sha256.New()
	sha.Write(source)

	return sha.Sum(nil)
}

func Normalize(raw []byte) []byte {
	result := bytes.ToUpper(raw)
	result = ZeroPattern.ReplaceAll(result, []byte("0"))
	result = OnePattern.ReplaceAll(result, []byte("1"))

	return result
}

func Checksum(raw []byte) []byte {
	dehexed := make([]byte, hex.DecodedLen(len(raw)))

	hex.Decode(dehexed, raw)

	hashed := hash(hash(dehexed))[:4]

	hexed := make([]byte, hex.EncodedLen(len(hashed)))

	hex.Encode(hexed, hashed)

	return hexed
}

func Encode(raw []byte) ([]byte, error) {
	if HexPattern.Match(raw) == false {
		return []byte{}, fmt.Errorf("not a hex string got %s", raw)
	}

	if (len(raw) % 2) != 0 {
		raw = append([]byte("0"), raw...)
	}

	raw = bytes.ToLower(raw)

	var result []byte
	var carry int

	for index := len(raw) - 1; index >= 0; index-- {
		if carry < 4 {
			current := bytes.IndexByte(HexAlphabet, raw[index]) >> carry

			if current < 0 {
				return []byte{}, fmt.Errorf("got invalid digit %s (%d) at %d\n", string(raw[index]), raw[index], index)
			}

			var next int

			if index != 0 {
				next = int(bytes.IndexByte(HexAlphabet, raw[index-1]))
			}

			bits := carry + 1
			low := (next % (1 << bits)) << (5 - bits)

			digit := C32Alphabet[int(current)+low]
			carry += 1

			result = append([]byte{digit}, result...)
		} else {
			carry = 0
		}
	}

	var leading int

	for index := 0; index < len(result); index++ {
		if result[index] != '0' {
			break
		} else {
			leading += 1
		}
	}

	return result[leading:], nil
}

func ChecksumEncode(raw []byte, version byte) ([]byte, error) {
	if version < 0 || version > 31 {
		return []byte{}, fmt.Errorf("version has to be between 0 - 31 got %d", version)
	}

	if HexPattern.Match(raw) == false {
		return []byte{}, fmt.Errorf("expected hex got %s", raw)
	}

	raw = bytes.ToLower(raw)

	if (len(raw) % 2) != 0 {
		raw = append([]byte("0"), raw...)
	}

	prefix := strconv.FormatInt(int64(version), 16)

	if (len(prefix) % 2) == 1 {
		prefix = "0" + prefix
	}

	checksum := Checksum(append([]byte(prefix), raw...))
	encoded, err := Encode(append(raw, checksum...))

	if err != nil {
		return []byte{}, fmt.Errorf("failed to encode got %v", err)
	}

	return append([]byte{C32Alphabet[version]}, encoded...), nil
}

func Decode(raw []byte) ([]byte, error) {
	raw = Normalize(raw)

	if C32Pattern.Match(raw) == false {
		return []byte{}, fmt.Errorf("expected C32 got %s", raw)
	}

	var prefix int

	for index := 0; raw[index] == '0'; index++ {
		prefix += 1
	}

	var result []byte
	var carry int
	var carryBits int

	for index := len(raw) - 1; index >= 0; index-- {
		if carryBits == 4 {
			result = append([]byte{HexAlphabet[carry]}, result...)

			carry = 0
			carryBits = 0
		}

		code := bytes.IndexByte(C32Alphabet, raw[index]) << carryBits
		value := code + carry

		digit := HexAlphabet[(value % 16)]

		carryBits += 1
		carry = value >> 4

		if carry > 1<<carryBits {
			return []byte{}, errors.New("failed to decode, invalid string")
		}

		result = append([]byte{digit}, result...)
	}

	result = append([]byte{HexAlphabet[carry]}, result...)

	if len(result)%2 == 1 {
		result = append([]byte{'0'}, result...)
	}

	var resultPrefix int

	for index := 0; result[index] == '0'; index++ {
		resultPrefix += 1
	}

	result = result[(resultPrefix - (resultPrefix % 2)):]

	for index := 0; index < prefix; index++ {
		result = append([]byte("00"), result...)
	}

	return result, nil
}

func ChecksumDecode(raw []byte) ([]byte, byte, error) {
	data, err := Decode(raw[1:])

	if err != nil {
		return []byte{}, 0, err
	}

	version := byte(bytes.IndexByte(C32Alphabet, raw[0]))

	checksum := data[len(data)-8:]
	content := data[:len(data)-8]

	prefix := strconv.FormatInt(int64(version), 16)

	if len(prefix) == 1 {
		prefix = "0" + prefix
	}

	encoded := Checksum(append([]byte(prefix), content...))

	if bytes.Compare(encoded, checksum) != 0 {
		return []byte{}, 0, errors.New("checksum is invalid")
	}

	return content, version, nil
}
