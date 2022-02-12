package transaction

import (
	"github.com/syvita/racks/constant"

	"encoding/binary"
	"encoding/hex"
	"errors"
)

type StacksTransaction struct {
	Version           constant.TransactionVersion
	ChainID           constant.ChainID
	Payload           Payload
	Authorization     Authorization
	AnchorMode        constant.AnchorMode
	PostConditionMode constant.PostConditionMode
	//TODO: figure out what is in a PostCondtion
	PostConditions []constant.PostCondition
}

func HexToBytes(plain string) ([]byte, error) {
	// if starts with 0x, remove it
	if plain[:2] == "0x" {
		plain = plain[2:]
	}

	return hex.DecodeString(plain)
}

func (transaction *StacksTransaction) Unmarshal(data []byte) error {
	// TODO: find minimum size, currently in place to prevent bounds errors
	if len(data) < 3 {
		return errors.New("transaction is too short to be valid")
	}

	// version byte is first byte
	transaction.Version = constant.TransactionVersion(data[0])

	if transaction.Version.Check() == false {
		return errors.New("transaction version is invalid")
	}

	// version byte is first byte
	data = data[1:]

	// next 4 bytes are the chain ID
	// TODO (Linden): change when Go hits 1.18 (https://stackoverflow.com/a/67199587)
	transaction.ChainID = constant.ChainID(*(*[4]byte)(data[:4]))

	if transaction.ChainID.Check() == false {
		return errors.New("transaction chainID is invalid")
	}

	data = data[4:]

	// next byte is the transaction authorization type
	authorizationType := data[0]
	data = data[1:]

	switch authorizationType {
	case 0x04:
		var condtion SingleSigSpendingCondition

		condtion.HashMode = HashMode(data[0])

		if condtion.HashMode.Check() == false {
			return errors.New("authorization must be signed with either P2PKH, P2SH, P2WPKH-P2SH or P2WSH-P2SH")
		}

		data = data[1:]

		// 20-byte public key hash
		condtion.Signer = *(*[20]byte)(data[:20])
		data = data[20:]

		// 8-byte nonce
		condtion.Nonce = binary.BigEndian.Uint64(data[:8])
		data = data[8:]

		// 8-byte fee
		condtion.Fee = binary.BigEndian.Uint64(data[:8])
		data = data[8:]

		switch condtion.HashMode {
		// 00 or 02 - single signature
		case HashModeP2PKH, HashModeP2WPKH_P2SH:
			// next byte is the public key encoding
			condtion.KeyEncoding = constant.PublicKeyEncoding(data[0])
			data = data[1:]

			if condtion.KeyEncoding.Check() == false {
				return errors.New("public key encoding must be compressed or uncompressed")
			}

			// next 65-byte is a recoverable ECDSA signature
			condtion.Signature = *(*[65]byte)(data[:65])
			data = data[65:]

		// 01 or 03 - multi signature
		case HashModeP2SH, HashModeP2WSH_P2SH:
			panic("TODO: implment HashModeP2SH and HashModeP2WSH_P2SH")
		}

		transaction.Authorization = StandardAuthorization{condtion}
	case 0x05:
		// sponsored authorization
		// two spending conditions.

		panic("TODO: implment sponsored authorization")
	}

	// assuming we've set authorization and updated data to the next byte

	// next byte is anchor mode
	transaction.AnchorMode = constant.AnchorMode(data[0])
	data = data[1:]

	if transaction.AnchorMode.Check() == false {
		return errors.New("anchor mode is invalid")
	}

	// next byte is post condition mode
	transaction.PostConditionMode = constant.PostConditionMode(data[0])
	data = data[1:]

	if transaction.PostConditionMode.Check() == false {
		return errors.New("post condition mode is invalid")
	}

	postConditionCount := binary.BigEndian.Uint32(data[:4])

	// TODO: loop through the post conditions
	for index := uint32(0); index < postConditionCount; index++ {
		// first 1 byte is the post condition type ID
		postConditionType := data[0]
		data = data[1:]
		switch postConditionType {
		case 00: // post condition type is STX
		case 01: // post condition type is fungible token
		case 02: // post condition type is non-fungible token (NFT)
		}
	}

	data = data[4:]

	return nil

	// export function deserializeTransaction(data: BufferReader | Uint8Array | string) {
	//   const version = bufferReader.readUInt8Enum(TransactionVersion, n => {
	//     throw new Error(`Could not parse ${n} as TransactionVersion`);
	//   });
	//   const chainId = bufferReader.readUInt32BE();
	//   const auth = deserializeAuthorization(bufferReader);
	//
	//   const anchorMode = bufferReader.readUInt8Enum(AnchorMode, n => {
	//     throw new Error(`Could not parse ${n} as AnchorMode`);
	//   });
	//   const postConditionMode = bufferReader.readUInt8Enum(PostConditionMode, n => {
	//     throw new Error(`Could not parse ${n} as PostConditionMode`);
	//   });
	//   const postConditions = deserializeLPList(bufferReader, StacksMessageType.PostCondition);
	//   const payload = deserializePayload(bufferReader);
	//
	//   return new StacksTransaction(
	//     version,
	//     auth,
	//     payload,
	//     postConditions,
	//     postConditionMode,
	//     anchorMode,
	//     chainId
	//   );
	// }
}

func (transaction *StacksTransaction) Marshal() ([]byte, error) {
	return []byte{}, nil
}
