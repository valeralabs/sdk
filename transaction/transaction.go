package transaction

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"unicode"

	"github.com/linden/bite"
	"github.com/valeralabs/sdk/constant"
	"github.com/valeralabs/sdk/encoding/clarity"
)

type StacksTransaction struct {
	Version       constant.TransactionVersion
	ChainID       constant.ChainID
	Payload       Payload
	Authorization Authorization
	AnchorMode    constant.AnchorMode
	// TODO: add post-conditions
	PostConditionMode constant.PostConditionMode
	PostConditions    []constant.PostCondition
}

func HexToBytes(plain string) ([]byte, error) {
	// if starts with 0x, remove it
	if plain[:2] == "0x" {
		plain = plain[2:]
	}

	return hex.DecodeString(plain)
}

func (transaction *StacksTransaction) Unmarshal(data []byte) error {
	decoded := make([]byte, hex.DecodedLen(len(data)))

	_, err := hex.Decode(decoded, data)

	if err != nil {
		return err
	}

	data = decoded

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

	transaction.PostConditionMode = constant.PostConditionMode(reader.ReadSingle())

	if transaction.PostConditionMode.Check() == false {
		return errors.New("post condition mode is invalid")
	}

	postConditionCount := binary.BigEndian.Uint32(reader.Read(4))

	//TODO (Linden): post conditions
	for index := uint32(0); index < postConditionCount; index++ {
		postConditionType := constant.PostConditionType(reader.ReadSingle())

		if postConditionType.Check() == false {
			return fmt.Errorf("post condition %d is not valid", index)
		}

		switch postConditionType {
		case constant.PostConditionTypeSTX:
		case constant.PostConditionTypeFT:
		case constant.PostConditionTypeNFT:
		}
	}

	payloadType := PayloadType(reader.ReadSingle())

	if payloadType.Check() == false {
		return errors.New("payload type is invalid")
	}

	switch payloadType {
	case PayloadTypeTokenTransfer:
		var payload PayloadTokenTransfer

		clarityType := clarity.ClarityType(reader.ReadSingle())

		if clarityType.Check() == false {
			return errors.New("address is invalid")
		}

		payload.Address, err = clarity.ParsePrincipal(clarityType, &reader)

		if err != nil {
			return err
		}

		payload.Amount = int(binary.BigEndian.Uint64(reader.Read(8)))
		payload.Memo = string(reader.Read(bytes.IndexByte(reader.Value, 0)))

		transaction.Payload = payload

	case PayloadTypeContractCall:
		var payload ContractCallTransfer

		payload.Address, err = clarity.ParsePrincipal(clarity.ClarityTypePrincipalContract, &reader)

		if err != nil {
			return err
		}

		function := clarity.Value{
			PrefixLength: 1,
		}

		err = function.Unmarshal(reader.Value, false)

		if err != nil {
			return err
		}

		reader.Read(len(function.Content) + 1)

		payload.Function = string(function.Content)

		for _, character := range payload.Function {
			if character > unicode.MaxASCII {
				return errors.New("function name must be ASCII")
			}
		}

		if len(payload.Function) > 128 {
			return errors.New("function name must be >= 128 characters")
		}

		arguments := clarity.List{
			PrefixLength:    4,
			SubPrefixLength: 4,
		}

		err = arguments.Unmarshal(reader.Value, true)

		if err != nil {
			return err
		}

		payload.Arguments = arguments

		transaction.Payload = payload

	default:
		return errors.New("invalid payload type")
	}

	return nil
}

func (transaction *StacksTransaction) Marshal() ([]byte, error) {
	return []byte{}, nil
}
