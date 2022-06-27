package c32check

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
)

var C32Alphabet = []byte("0123456789ABCDEFGHJKMNPQRSTVWXYZ")
var HexAlphabet = []byte("0123456789abcdef")

var HexPattern = regexp.MustCompile(`^[0-9a-fA-F]*$`)

func hash(source []byte) []byte {
	sha := sha256.New()
	sha.Write(source)

	return sha.Sum(nil)
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
