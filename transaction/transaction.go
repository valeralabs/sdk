package transaction

import (
	"bytes"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"unicode"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/linden/binstruct"
	"github.com/linden/bite"
	"valera.co/vdk/address"
	"valera.co/vdk/constant"
	"valera.co/vdk/encoding/clarity"
	"valera.co/vdk/wallet/keys"
)

type StacksTransaction struct {
	Version       constant.TransactionVersion
	ChainID       constant.ChainID
	Payload       Payload
	Authorization Authorization
	AnchorMode    constant.AnchorMode
	// TODO: add post-conditions
	PostConditionMode PostConditionMode
	PostConditions    []PostCondition
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

	transaction.ChainID = constant.ChainID(binary.BigEndian.Uint32(reader.Read(4)))

	if transaction.ChainID.Check() == false {
		return errors.New("transaction chainID is invalid")
	}

	authorizationType := reader.ReadSingle()

	switch authorizationType {
	case 0x04:
		var condition SingleSignatureSpendingCondition

		condition.HashMode = address.HashMode(reader.ReadSingle())

		if condition.HashMode.Check() == false {
			return errors.New("authorization must be signed with either P2PKH, P2SH, P2WPKH-P2SH or P2WSH-P2SH")
		}

		condition.Signer = *(*[20]byte)(reader.Read(20))
		condition.Nonce = binary.BigEndian.Uint64(reader.Read(8))
		condition.Fee = binary.BigEndian.Uint64(reader.Read(8))

		switch condition.HashMode {
		case address.HashModeP2PKH, address.HashModeP2WPKH:
			condition.KeyEncoding = constant.PublicKeyEncoding(reader.ReadSingle())

			if condition.KeyEncoding.Check() == false {
				return errors.New("public key encoding must be compressed or uncompressed")
			}

			condition.Signature = *(*[65]byte)(reader.Read(65))

		case address.HashModeP2SH, address.HashModeP2WSH:
			panic("TODO: implement HashModeP2SH and HashModeP2WSH_P2SH")
		}

		transaction.Authorization = StandardAuthorization{&condition}

	case 0x05:
		// sponsored authorization
		// two spending conditions.

		panic("TODO: implement sponsored authorization")
	}

	transaction.AnchorMode = constant.AnchorMode(reader.ReadSingle())

	if transaction.AnchorMode.Check() == false {
		return errors.New("anchor mode is invalid")
	}

	transaction.PostConditionMode = PostConditionMode(reader.ReadSingle())

	if transaction.PostConditionMode.Check() == false {
		return errors.New("post condition mode is invalid")
	}

	postConditionCount := binary.BigEndian.Uint32(reader.Read(4))
	postConditions := []PostCondition{}

	for index := uint32(0); index < postConditionCount; index++ {
		postConditionType := PostConditionType(reader.ReadSingle())

		if postConditionType.Check() == false {
			return fmt.Errorf("post condition %d is not valid", index)
		}

		switch postConditionType {
		case PostConditionTypeSTX:
			hash := transaction.Authorization.GetCondition().GetSigner()
			mode := transaction.Authorization.GetCondition().GetHashMode()

			network := address.NetworkMainnet

			if transaction.Version == constant.TransactionVersionTestnet {
				network = address.NetworkTestnet
			}

			version := address.AddressVersion{mode, network}

			origin := address.Address{
				Version: version,
				Hash:    hash[:],
			}

			principal, err := DecodePostConditionPrincipal(&reader, origin)

			if err != nil {
				return fmt.Errorf("could not decode post condition principal: %v", err)
			}

			condition := FungibleConditionCode(reader.ReadSingle())

			if condition.Check() == false {
				return errors.New("fungible condition code is invalid")
			}

			amount := binary.BigEndian.Uint64(reader.Read(8))

			postConditions = append(postConditions, PostConditionSTX{
				Principal: principal,
				Condition: condition,
				Amount:    amount,
			})

		case PostConditionTypeFT:
			hash := transaction.Authorization.GetCondition().GetSigner()
			mode := transaction.Authorization.GetCondition().GetHashMode()

			network := address.NetworkMainnet

			if transaction.Version == constant.TransactionVersionTestnet {
				network = address.NetworkTestnet
			}

			version := address.AddressVersion{mode, network}

			origin := address.Address{
				Version: version,
				Hash:    hash[:],
			}

			principal, err := DecodePostConditionPrincipal(&reader, origin)

			if err != nil {
				return fmt.Errorf("could not deserialize post condition principal: %v", err)
			}

			asset, err := DecodeAsset(&reader)

			if err != nil {
				return fmt.Errorf("could not deserialize asset info: %v", err)
			}

			condition := FungibleConditionCode(reader.ReadSingle())

			if condition.Check() == false {
				return fmt.Errorf("fungible condition code is invalid: %v", err)
			}

			amount := binary.BigEndian.Uint64(reader.Read(8))

			postCondition := PostConditionFT{condition, principal, amount, asset}

			postConditions = append(postConditions, postCondition)

		case PostConditionTypeNFT:
			hash := transaction.Authorization.GetCondition().GetSigner()
			mode := transaction.Authorization.GetCondition().GetHashMode()

			network := address.NetworkMainnet

			if transaction.Version == constant.TransactionVersionTestnet {
				network = address.NetworkTestnet
			}

			version := address.AddressVersion{mode, network}

			origin := address.Address{
				Version: version,
				Hash:    hash[:],
			}

			principal, err := DecodePostConditionPrincipal(&reader, origin)

			if err != nil {
				return fmt.Errorf("failed to deserialize post condition principal: %v", err)
			}

			asset, err := DecodeAsset(&reader)

			if err != nil {
				return fmt.Errorf("failed to deserialize asset: %v", err)
			}

			var value clarity.Value

			err = value.Unmarshal(reader.Value, true)

			if err != nil {
				return fmt.Errorf("could not unmarshal asset value: %v", err)
			}

			reader.Read(value.Length(true))

			condition := NonFungibleConditionCode(reader.ReadSingle())

			if condition.Check() == false {
				return fmt.Errorf("non-fungible condition code is invalid: %v", err)
			}

			postConditions = append(postConditions, PostConditionNFT{condition, principal, value, asset})
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
	writer.WriteUint32(uint32(transaction.ChainID))

	switch authorization := any(transaction.Authorization).(type) {
	case StandardAuthorization:
		switch condition := any(authorization.GetCondition()).(type) {
		case *SingleSignatureSpendingCondition:
			writer.WriteUint8(uint8(0x04))
			writer.WriteUint8(uint8(condition.HashMode))

			writer.Write(condition.Signer[:])

			writer.WriteUint64(condition.Nonce)
			writer.WriteUint64(condition.Fee)

			switch condition.HashMode {
			case address.HashModeP2PKH, address.HashModeP2WPKH:
				writer.WriteUint8(uint8(condition.KeyEncoding))
				writer.Write(condition.Signature[:])

			case address.HashModeP2SH, address.HashModeP2WSH:
				panic("TODO: implement HashModeP2SH and HashModeP2WSH_P2SH")
			}

		default:
			panic("TODO: implement multiple spending condition")
		}

	default:
		panic("TODO: implement sponsored authorization")
	}

	writer.WriteUint8(uint8(transaction.AnchorMode))

	writer.WriteUint8(uint8(transaction.PostConditionMode))
	writer.WriteUint32(uint32(len(transaction.PostConditions)))

	for _, postCondition := range transaction.PostConditions {
		switch any(postCondition).(type) {
		case PostConditionSTX:
			writer.WriteUint8(uint8(PostConditionTypeSTX))

			postCondition := any(postCondition).(PostConditionSTX)

			postCondition.Principal.Encode(writer)

			writer.WriteUint8(uint8(postCondition.Condition))

			writer.WriteUint64(postCondition.Amount)

		case PostConditionFT:
			writer.WriteUint8(uint8(PostConditionTypeFT))

			postCondition := any(postCondition).(PostConditionFT)

			postCondition.Principal.Encode(writer)

			postCondition.Asset.Encode(writer)

			writer.WriteUint8(uint8(postCondition.Condition))

			writer.WriteUint64(postCondition.Amount)

		case PostConditionNFT:
			writer.WriteUint8(uint8(PostConditionTypeNFT))

			postCondition := any(postCondition).(PostConditionNFT)

			postCondition.Principal.Encode(writer)

			postCondition.Asset.Encode(writer)

			valueBytes, err := postCondition.Value.Marshal(true)

			if err != nil {
				return []byte{}, fmt.Errorf("could not marshal post condition value: %v", err)
			}

			writer.Write(valueBytes)

			writer.WriteUint8(uint8(postCondition.Condition))
		}
	}

	payload, err := transaction.Payload.Marshal()

	if err != nil {
		return []byte{}, err
	}

	writer.Write(payload)

	raw := buffer.Bytes()

	encoded := make([]byte, hex.EncodedLen(len(raw)))
	hex.Encode(encoded, raw)

	return encoded, nil
}

func (transaction *StacksTransaction) Sign(private keys.PrivateKey) error {
	// set the signer
	public := private.PublicKey().Serialize()

	hash160 := btcutil.Hash160(public)

	condition := transaction.Authorization.GetCondition()
	condition.SetSigner(*(*[20]byte)(hash160))
	transaction.Authorization = transaction.Authorization.SetCondition(condition)

	// get the empty transaction
	empty := *transaction
	empty.Authorization = empty.Authorization.SetCondition(empty.Authorization.GetCondition().Clear())

	hexed, err := empty.Marshal()

	if err != nil {
		return err
	}

	decoded := make([]byte, hex.DecodedLen(len(hexed)))
	hex.Decode(decoded, hexed)

	hash := sha512.Sum512_256(decoded)

	// order presign
	presign := new(bytes.Buffer)

	presign.Write(hash[:])
	presign.WriteByte(0x04)
	binary.Write(presign, binary.BigEndian, condition.GetFee())
	binary.Write(presign, binary.BigEndian, condition.GetNonce())

	if presign.Len() != 49 {
		return errors.New("presign is invalid")
	}

	presigned := sha512.Sum512_256(presign.Bytes())

	signature, err := private.SignRecoverable(presigned[:])

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	if err != nil {
		return err
	}

	condition.SetSignature(*(*[65]byte)(signature))

	transaction.Authorization.SetCondition(condition)

	return nil
}
