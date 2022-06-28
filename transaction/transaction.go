package transaction

import (
	"github.com/valeralabs/sdk/constant"

	"encoding/binary"
	"encoding/hex"
	"errors"

	"github.com/linden/bite"
)

type StacksTransaction struct {
	Version       constant.TransactionVersion
	ChainID       constant.ChainID
	Payload       Payload
	Authorization Authorization
	AnchorMode    constant.AnchorMode
	// TODO: add post-conditions
	// PostConditionMode constant.PostConditionMode
	// PostConditions    []constant.PostCondition
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

	reader := bite.New(data)

	transaction.Version = constant.TransactionVersion(reader.ReadSingle())

	if transaction.Version.Check() == false {
		return errors.New("transaction version is invalid")
	}

	// TODO (Linden): change when Go hits 1.19 (https://stackoverflow.com/a/67199587)
	transaction.ChainID = constant.ChainID(*(*[4]byte)(reader.Read(4)))

	if transaction.ChainID.Check() == false {
		return errors.New("transaction chainID is invalid")
	}

	authorizationType := reader.ReadSingle()

	switch authorizationType {
	case 0x04:
		var condtion SingleSignatureSpendingCondition

		condtion.HashMode = constant.HashMode(reader.ReadSingle())

		if condtion.HashMode.Check() == false {
			return errors.New("authorization must be signed with either P2PKH, P2SH, P2WPKH-P2SH or P2WSH-P2SH")
		}

		condtion.Signer = *(*[20]byte)(reader.Read(20))
		condtion.Nonce = binary.BigEndian.Uint64(reader.Read(8))
		condtion.Fee = binary.BigEndian.Uint64(reader.Read(8))

		switch condtion.HashMode {
		case constant.HashModeP2PKH, constant.HashModeP2WPKH:
			condtion.KeyEncoding = constant.PublicKeyEncoding(reader.ReadSingle())

			if condtion.KeyEncoding.Check() == false {
				return errors.New("public key encoding must be compressed or uncompressed")
			}

			condtion.Signature = *(*[65]byte)(reader.Read(65))

		case constant.HashModeP2SH, constant.HashModeP2WSH:
			panic("TODO: implment HashModeP2SH and HashModeP2WSH_P2SH")
		}

		transaction.Authorization = StandardAuthorization[SingleSignatureSpendingCondition]{condtion}
	case 0x05:
		// sponsored authorization
		// two spending conditions.

		panic("TODO: implment sponsored authorization")
	}

	transaction.AnchorMode = constant.AnchorMode(reader.ReadSingle())

	if transaction.AnchorMode.Check() == false {
		return errors.New("anchor mode is invalid")
	}

	// // next byte is post condition mode
	// transaction.PostConditionMode = constant.PostConditionMode(reader.ReadSingle())
	// if transaction.PostConditionMode.Check() == false {
	// 	return errors.New("post condition mode is invalid")
	// }
	// postConditionCount := binary.BigEndian.Uint32(reader.Read(4))
	// // TODO: loop through the post conditions
	// for index := uint32(0); index < postConditionCount; index++ {
	// 	// first 1 byte is the post condition type ID
	// 	postConditionType := data[0]
	// 	data = data[1:]
	// 	switch postConditionType {
	// 	case 00: // post condition type is STX
	// 	case 01: // post condition type is fungible token
	// 	case 02: // post condition type is non-fungible token (NFT)
	// 	}
	// }

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
