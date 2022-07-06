package transaction

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"unicode"

	"github.com/linden/binstruct"
	"github.com/linden/bite"
	"github.com/valeralabs/sdk/address"
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
	PostConditions    []PostCondition
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

	reader := bite.NewReader(data)

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

		transaction.Authorization = StandardAuthorization(SingleSignatureSpendingCondition(condtion))
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

	postConditions := []PostCondition{}

	//TODO (Linden): post conditions
	for index := uint32(0); index < postConditionCount; index++ {
		postConditionType := constant.PostConditionType(reader.ReadSingle())

		if postConditionType.Check() == false {
			return fmt.Errorf("post condition %d is not valid", index)
		}

		switch postConditionType {
		case constant.PostConditionTypeSTX:
			addressHash, err := GetSignerFromAuthorization(transaction.Authorization)
			if err != nil {
				return fmt.Errorf("Could not get signer from authorization: %v", err)
			}


			hashMode, err := GetSignerHashModeFromAuthorization(transaction.Authorization)
			if err != nil {
				return fmt.Errorf("Could not get signer hash mode from transaction authorization: %v", err)
			}

			addressVersion := HashModeToAddressVersion(hashMode, transaction.Version)

			originAddress := address.NewAddressFromPublicKeyHash(addressHash[:], addressVersion)

			postConditionPrincipal, err := DeserializePostConditionPrincipal(&reader, originAddress)
			if err != nil {
				return err
			}

			fungibleConditionCode := FungibleConditionCode(reader.ReadSingle())
			if !fungibleConditionCode.Check() {
				return errors.New("Fungible condition code is invalid")
			}

			amount := binary.BigEndian.Uint64(reader.Read(8))

			postCondition := PostConditionSTX {
				postConditionPrincipal: postConditionPrincipal,
				fungibleConditionCode: fungibleConditionCode,
				amount: amount,
			}

			postConditions = append(postConditions, PostCondition(postCondition))
		case constant.PostConditionTypeFT:
			addressHash, err := GetSignerFromAuthorization(transaction.Authorization)
			if err != nil {
				return fmt.Errorf("Could not get signer from transaction authorization: %v", err)
			}

			hashMode, err := GetSignerHashModeFromAuthorization(transaction.Authorization)
			if err != nil {
				return fmt.Errorf("Could not get signer hash mode from transaction authorization: %v", err)
			}

			addressVersion := HashModeToAddressVersion(hashMode, transaction.Version)

			originAddress := address.NewAddressFromPublicKeyHash(addressHash[:], addressVersion)
	
			principalAddress, err := DeserializePostConditionPrincipal(&reader, originAddress)
			if err != nil {
				return fmt.Errorf("Could not deserialize post condition principal: %v", err)
			}

			assetInfo, err := DeserializeAssetInfo(&reader)
			if err != nil {
				return fmt.Errorf("Could not deserialize asset info: %v", err)
			}

			fungibleConditionCode := FungibleConditionCode(reader.ReadSingle())
			fungibleConditionCode.Check()

			amount := binary.BigEndian.Uint64(reader.Read(8))

			postCondition := PostConditionFungible {
				postConditionPrincipal: principalAddress,
				assetInfo: assetInfo,
				fungibleConditionCode: fungibleConditionCode,
				amount: amount,
			}

			postConditions = append(postConditions, postCondition)
		case constant.PostConditionTypeNFT:
			panic("non fungible token post conditions are not yet implemented")
		}
	}

	transaction.PostConditions = postConditions

	payloadType := PayloadType(reader.ReadSingle())

	if payloadType.Check() == false {
		return fmt.Errorf("payload type is invalid, it cannot be %v", byte(payloadType))
	}

	switch payloadType {
	case PayloadTypeTokenTransfer:
		var payload PayloadTokenTransfer

		clarityType := clarity.ClarityType(reader.ReadSingle())

		if clarityType.Check() == false {
			return errors.New("address is invalid")
		}

		payload.Address, err = clarity.DecodePrincipal(clarityType, &reader)

		if err != nil {
			return err
		}

		payload.Amount = int(binary.BigEndian.Uint64(reader.Read(8)))
		payload.Memo = string(reader.Read(bytes.IndexByte(reader.Value, 0)))

		transaction.Payload = payload

	case PayloadTypeSmartContract:
		var payload PayloadSmartContract

		name := clarity.Value{
			PrefixLength: 1,
		}

		err = name.Unmarshal(reader.Value, false)

		if err != nil {
			return errors.New("invalid contract name")
		}

		reader.Read(name.Length(false))

		payload.Name = string(name.Content)

		body := clarity.Value{
			PrefixLength: 4,
		}

		err = body.Unmarshal(reader.Value, false)

		if err != nil {
			return errors.New("invalid contract body")
		}

		reader.Read(body.Length(false))

		payload.Body = string(body.Content)

		transaction.Payload = payload

	case PayloadTypeContractCall:
		var payload PayloadContractCall

		payload.Address, err = clarity.DecodePrincipal(clarity.ClarityTypePrincipalContract, &reader)

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

		reader.Read(function.Length(false))

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
	var buffer bytes.Buffer

	writer := binstruct.NewWriter(&buffer, binary.BigEndian, false)

	writer.WriteUint8(uint8(transaction.Version))
	writer.Write(transaction.ChainID[:])

	switch authorization := any(transaction.Authorization).(type) {
	case StandardAuthorization:
		switch condtion := any(authorization).(type) {
		case SingleSignatureSpendingCondition:
			writer.WriteUint8(uint8(0x04))
			writer.WriteUint8(uint8(condtion.HashMode))

			writer.Write(condtion.Signer[:])

			writer.WriteUint64(condtion.Nonce)
			writer.WriteUint64(condtion.Fee)

			switch condtion.HashMode {
			case constant.HashModeP2PKH, constant.HashModeP2WPKH:
				writer.WriteUint8(uint8(condtion.KeyEncoding))
				writer.Write(condtion.Signature[:])

			case constant.HashModeP2SH, constant.HashModeP2WSH:
				panic("TODO: implment HashModeP2SH and HashModeP2WSH_P2SH")
			}

		default:
			panic("TODO: implment sponsored authorization")
		}

	default:
		panic("TODO: implment sponsored authorization")
	}

	writer.WriteUint8(uint8(transaction.AnchorMode))

	writer.WriteUint8(uint8(transaction.PostConditionMode))
	writer.WriteUint32(uint32(len(transaction.PostConditions)))

	//TODO: handle strict post-conditions

	switch any(transaction.Payload).(type) {
	case PayloadTokenTransfer:
		writer.WriteUint8(uint8(PayloadTypeTokenTransfer))

		payload := any(transaction.Payload).(PayloadTokenTransfer)

		if payload.Address.Contract == "" {
			writer.WriteUint8(uint8(clarity.ClarityTypePrincipalStandard))
		} else {
			writer.WriteUint8(uint8(clarity.ClarityTypePrincipalContract))
		}

		err := clarity.EncodePrincipal(payload.Address, writer)

		if err != nil {
			return []byte{}, err
		}

		writer.WriteUint64(uint64(payload.Amount))

		// TODO: figure out this padding
		length := writer.Len() + constant.MemoLengthBytes + 2

		writer.Write([]byte(payload.Memo))

		for writer.Len() < length {
			writer.WriteByte(0x00)
		}

	case PayloadSmartContract:
		writer.WriteUint8(uint8(PayloadTypeSmartContract))

		payload := any(transaction.Payload).(PayloadSmartContract)

		name := clarity.Value{
			Content:      []byte(payload.Name),
			PrefixLength: 1,
		}

		raw, err := name.Marshal(false)

		if err != nil {
			return []byte{}, err
		}

		writer.Write(raw)

		body := clarity.Value{
			Content:      []byte(payload.Body),
			PrefixLength: 4,
		}

		raw, err = body.Marshal(false)

		if err != nil {
			return []byte{}, err
		}

		writer.Write(raw)

	case PayloadContractCall:
		writer.WriteUint8(uint8(PayloadTypeContractCall))

		payload := any(transaction.Payload).(PayloadContractCall)

		err := clarity.EncodePrincipal(payload.Address, writer)

		if err != nil {
			return []byte{}, err
		}

		function := clarity.Value{
			Content:      []byte(payload.Function),
			PrefixLength: 1,
		}

		raw, err := function.Marshal(false)

		if err != nil {
			return []byte{}, err
		}

		writer.Write(raw)

		raw, err = payload.Arguments.Marshal(true)

		if err != nil {
			return []byte{}, err
		}

		writer.Write(raw)

	default:
		panic("TODO: handle other payloads")

	}

	raw := buffer.Bytes()

	encoded := make([]byte, hex.EncodedLen(len(raw)))
	hex.Encode(encoded, raw)

	return encoded, nil
}
