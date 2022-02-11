package c32check

import (
	"github.com/syvita/racks/encoding/c32check/crock32"

	"crypto/sha256"
)

type Version int

const (
	MainnetP2PKH = Version(0x22) // 'P'
	MainnetP2SH  = Version(0x20) // 'M'
	TestnetP2PKH = Version(0x26) // 'T'
	TestnetP2SH  = Version(0x21) // 'N'
)

// version []byte, payload []byte, checksum []byte
// checksum = sha256(sha256(version + payload)).substring(0,4)

func Encode(version Version, body []byte) ([]byte, error) {
	payload, err := crock32.Decode(body)

	if err != nil {
		return []byte{}, err
	}

	round := sha256.Sum256([]byte{
		byte(int(version) + int(payload)),
	})

	checksum := sha256.Sum256(round[:])

	return checksum[:], nil
}

// c32address = "S" + c32(version + payload) + byteToHexString(checksum)

// TODO
// func CreateC32Address(version Version, payload []byte) ([]byte, error) {
// 	checksum, err := Encode(version, payload)
//
// 	hash(ripemd160.New()
//
// 	// uh okay so payload would be the byte array of the public key
// 	// we sha256 that
// 	// then we RIPEMD160 that
// 	// then we pass it to `payload` in this function
// 	// i think?
// 	//
//
// 	if err != nil {
// 		return []byte{}, err
// 	}
//
// 	return []byte(fmt.Sprintf("S%s%s", crock32.Encode(binary.BigEndian.Uint64(payload)), checksum[:4])), nil
// }
