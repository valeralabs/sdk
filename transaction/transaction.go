package transaction

import (
	"hash"

	"github.com/syvita/racks/constant"

	"errors"
	"fmt"
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

func (transaction *StacksTransaction) Unmarshal(data []byte) error {
	//TODO: find minimum size, currently in place to prevent bounds errors
	if len(data) < 3 {
		return errors.New("transaction is too short to be valid")
	}

	if string(data[:2]) == "0x" {
		data = data[2:]
	}

	//originalData := data

	// version byte is first byte
	version := data[0]
	data = data[1:]

	// next 4 bytes are the chain id
	chainID := constant.ChainID(data[:4])
	data = data[4:]

	// next byte is the transaction authorization type
	authorizationType := data[0]
	switch authorizationType {
		case 04: // standard authorization
			// one spending condition.

			// hashMode byte
			hashMode := data[1]
			data = data[2:]
			// 20-byte public key hash
			publicKeyHash := data[:20]
			data = data[20:]
			// 8-byte nonce
			nonce := data[:8]
			data = data[8:]
			// 8-byte fee
			fee := data[:8]
			data = data[8:]
			
			switch hashMode {
				// 00 or 02 - single signature
				case 00, 02:
					// next byte is the public key encoding
					publicKeyEncoding := constant.PublicKeyEncoding(data[0])
					data = data[1:]
					// next 65-byte is a recoverable ECDSA signature
					signature := data[:65]
					data = data[65:]

				// 01 or 03 - multi signature
				case 01, 03:
					
			}
		case 05: // sponsored authorization
			// two spending conditions.
	}

	fmt.Println(string(data))

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
