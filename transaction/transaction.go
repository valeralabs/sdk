package transaction

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"unicode"

	"github.com/linden/bite"
	"github.com/valeralabs/sdk/address/c32"
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

	postConditions := []PostCondition{}

	//TODO (Linden): post conditions
	for index := uint32(0); index < postConditionCount; index++ {
		postConditionType := constant.PostConditionType(reader.ReadSingle())

		if postConditionType.Check() == false {
			return fmt.Errorf("post condition %d is not valid", index)
		}

		switch postConditionType {
		case constant.PostConditionTypeSTX:
			// TODO: make it actually work
			addressHash := transaction.Authorization.GetSigner()
			hashMode := transaction.Authorization.GetSignerHashMode()

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

		case constant.PostConditionTypeNFT:
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

		payload.Address, err = clarity.ParsePrincipal(clarityType, &reader)

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

func WriteWithLength[T ~[]byte | ~byte | ~int](buffer *bytes.Buffer, content T, length int) {
	before := buffer.Len()

	padding := true

	switch any(content).(type) {
	case []byte:
		buffer.Write(any(content).([]byte))

	case byte:
		buffer.WriteByte(any(content).(byte))

	case int:
		padding = false

		temporary := new(bytes.Buffer)

		binary.Write(temporary, binary.BigEndian, int8(any(content).(int)))

		for buffer.Len()+temporary.Len() < (before + length) {
			buffer.WriteByte(0)
		}

		buffer.Write(temporary.Bytes())
	}

	if padding == true {
		for buffer.Len() < (before + length) {
			buffer.WriteByte(0)
		}
	}
}

func (transaction *StacksTransaction) Marshal() ([]byte, error) {
	buffer := new(bytes.Buffer)

	WriteWithLength(buffer, transaction.Version, 1)

	chainID := ([4]byte)(transaction.ChainID)
	WriteWithLength(buffer, chainID[:], 4)

	switch any(transaction.Authorization).(type) {
	case StandardAuthorization[SingleSignatureSpendingCondition]:
		condtion := any(transaction.Authorization).(StandardAuthorization[SingleSignatureSpendingCondition]).SpendingCondition

		WriteWithLength(buffer, byte(0x04), 1)
		WriteWithLength(buffer, condtion.HashMode, 1)

		signer := (([20]byte)(condtion.Signer))
		WriteWithLength(buffer, signer[:], 20)

		binary.Write(buffer, binary.BigEndian, condtion.Nonce)
		binary.Write(buffer, binary.BigEndian, condtion.Fee)

		switch condtion.HashMode {
		case constant.HashModeP2PKH, constant.HashModeP2WPKH:
			WriteWithLength(buffer, condtion.KeyEncoding, 1)

			signature := (([65]byte)(condtion.Signature))
			WriteWithLength(buffer, signature[:], 65)

		case constant.HashModeP2SH, constant.HashModeP2WSH:
			panic("TODO: implment HashModeP2SH and HashModeP2WSH_P2SH")
		}

	default:
		panic("TODO: implment sponsored authorization")
	}

	WriteWithLength(buffer, byte(transaction.AnchorMode), 1)

	WriteWithLength(buffer, byte(transaction.PostConditionMode), 1)
	WriteWithLength(buffer, len(transaction.PostConditions), 4)

	//TODO: handle strict post-conditions

	switch any(transaction.Payload).(type) {
	case PayloadTokenTransfer:
		WriteWithLength(buffer, byte(PayloadTypeTokenTransfer), 1)

		payload := any(transaction.Payload).(PayloadTokenTransfer)

		if payload.Address.Contract == "" {
			WriteWithLength(buffer, byte(clarity.ClarityTypePrincipalStandard), 1)
		} else {
			WriteWithLength(buffer, byte(clarity.ClarityTypePrincipalContract), 1)
		}

		version, err := c32.ConvertVersion(payload.Address.Version)

		if err != nil {
			return []byte{}, err
		}

		WriteWithLength(buffer, byte(version), 1)
		WriteWithLength(buffer, payload.Address.Hash, 20)

		if payload.Address.Contract != "" {
			name := clarity.Value{
				Content:      []byte(payload.Address.Contract),
				PrefixLength: 1,
			}

			raw, err := name.Marshal(false)

			if err != nil {
				return []byte{}, err
			}

			WriteWithLength(buffer, raw, 0)
		}

		WriteWithLength(buffer, payload.Amount, 8)
		WriteWithLength(buffer, []byte(payload.Memo), constant.MemoLengthBytes)

		// TODO: figure out this padding
		WriteWithLength(buffer, []byte{0x00, 0x00}, 2)

	case PayloadSmartContract:
		WriteWithLength(buffer, byte(PayloadTypeSmartContract), 1)

		payload := any(transaction.Payload).(PayloadSmartContract)

		name := clarity.Value{
			Content:      []byte(payload.Name),
			PrefixLength: 1,
		}

		raw, err := name.Marshal(false)

		if err != nil {
			return []byte{}, err
		}

		WriteWithLength(buffer, raw, 0)

		body := clarity.Value{
			Content:      []byte(payload.Body),
			PrefixLength: 4,
		}

		raw, err = body.Marshal(false)

		if err != nil {
			return []byte{}, err
		}

		WriteWithLength(buffer, raw, 0)

	default:
		panic("TODO: handle other payloads")

	}

	raw := buffer.Bytes()

	encoded := make([]byte, hex.EncodedLen(len(raw)))
	hex.Encode(encoded, raw)

	return encoded, nil
}

