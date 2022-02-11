package transaction

import (
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
