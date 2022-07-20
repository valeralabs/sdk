package VDK

import (
	"encoding/hex"
	"errors"
	"fmt"

	"valera.co/vdk/api"
	"valera.co/vdk/constant"
	"valera.co/vdk/encoding/clarity"
	"valera.co/vdk/transaction"
)

// Wrapper around [transaction.Payload].
type Payload struct {
	value *transaction.Payload
}

// Wrapper around [transaction.Payload].
type PostCondition struct {
	value *[]transaction.PostCondition
}

// Wrapper around [transaction.StacksTransaction].
type StacksTransaction struct {
	Version    string
	ChainID    string
	AnchorMode string

	Payload *Payload

	value *transaction.StacksTransaction
}

func create(payload transaction.Payload) transaction.StacksTransaction {
	result := transaction.StacksTransaction{
		Version:           constant.TransactionVersionMainnet,
		ChainID:           constant.ChainIDMainnet,
		Payload:           payload,
		AnchorMode:        constant.AnchorModeAny,
		Authorization:     transaction.StandardAuthorization{&transaction.SingleSignatureSpendingCondition{}},
		PostConditionMode: transaction.PostConditionModeLoose,
	}

	return result
}

func bind(from transaction.StacksTransaction) *StacksTransaction {
	return &StacksTransaction{
		Version:    from.Version.String(),
		ChainID:    from.ChainID.String(),
		AnchorMode: from.AnchorMode.String(),

		Payload: &Payload{&from.Payload},
		value:   &from,
	}
}

// Get the underlying Stacks transaction as a String.
func (cursor *StacksTransaction) Raw() string {
	return fmt.Sprintf("%#+v\n", *cursor.value)
}

// Encode a Stacks Transaction.
func (cursor *StacksTransaction) Encode() (string, error) {
	encoded, err := (*cursor.value).Marshal()
	return string(encoded), err
}

// Estimate the fee of a Stacks Transaction.
func (cursor *StacksTransaction) EstimateFee() (int, error) {
	marshaled, err := (*cursor.value).Payload.Marshal()

	if err != nil {
		return 0, err
	}

	hexed := make([]byte, hex.EncodedLen(len(marshaled)))
	hex.Encode(hexed, marshaled)

	fee, err := api.EstimateFee(hexed)

	if err != nil {
		return 5000, nil
	}

	return fee, nil
}

// Set the nonce manually.
func (stacks *StacksTransaction) SetNonce(nonce int) {
	condition := (*stacks.value).Authorization.GetCondition()
	condition.SetNonce(uint64(nonce))
}

// Set the fee manually.
func (stacks *StacksTransaction) SetFee(fee int) {
	condition := (*stacks.value).Authorization.GetCondition()
	condition.SetFee(uint64(fee))
}

// Sign a Stacks Transaction.
// `account`: the account used to sign the transaction.
func (stacks *StacksTransaction) Sign(account *Account) error {
	if account == nil {
		return errors.New("account is nil")
	}

	if stacks == nil {
		return errors.New("stacks is nil")
	}

	if (*stacks.value).Authorization == nil {
		(*stacks.value).Authorization = transaction.StandardAuthorization{
			Condition: &transaction.SingleSignatureSpendingCondition{},
		}
	}

	condition := (*stacks.value).Authorization.GetCondition()

	if condition.GetNonce() == 0 {
		nonce, err := account.NextNonce()

		if err != nil {
			return err
		}

		condition.SetNonce(uint64(nonce))
	}

	if condition.GetFee() == 0 {
		fee, err := stacks.EstimateFee()

		if err != nil {
			return err
		}

		condition.SetFee(uint64(fee))
	}

	(*stacks.value).Authorization = (*stacks.value).Authorization.SetCondition(condition)

	err := (*stacks.value).Sign(*account.value.PrivateKey)

	if err != nil {
		return err
	}

	return nil
}

// Broadcast a Stacks Transaction to the Stacks network.
func (stacks *StacksTransaction) Broadcast() error {
	if stacks == nil {
		return errors.New("transaction is nil")
	}

	raw, err := (*stacks.value).Marshal()

	if err != nil {
		return err
	}

	decoded := make([]byte, hex.DecodedLen(len(raw)))
	hex.Decode(decoded, raw)

	err = api.BroadcastTransaction(decoded)

	if err != nil {
		return err
	}

	return nil
}

func (stacks *StacksTransaction) SetCondition(conditions *PostCondition, strict bool) {
	(*stacks.value).PostConditions = *conditions.value

	if strict == true {
		(*stacks.value).PostConditionMode = transaction.PostConditionModeStrict
	} else {
		(*stacks.value).PostConditionMode = transaction.PostConditionModeLoose
	}
}

// Parse a Stacks Transaction (hex encoded).
func NewStacksTransaction(raw string) (*StacksTransaction, error) {
	var cursor transaction.StacksTransaction

	err := cursor.Unmarshal([]byte(raw))

	if err != nil {
		return &StacksTransaction{}, err
	}

	return bind(cursor), nil
}

// Create a token transfer.
// `recipient`: the destination of the funds can be a standard or contract principal.
// `amount`: total uSTX sent.
// `memo`: optional arbitrary info.
// `conditions`: the post-conditions.
// `strict`: has to follow post-condtions.
func NewTokenTransfer(recipient *Principal, amount int, memo string) (*StacksTransaction, error) {
	if recipient == nil {
		return &StacksTransaction{}, errors.New("recipient is nil")
	}

	payload := transaction.PayloadTokenTransfer{
		Address: *recipient.value,
		Amount:  amount,
		Memo:    memo,
	}

	return bind(create(payload)), nil
}

// Create a contract call.
// `contract`: the principal of the contract.
// `function`: the function being called.
// `arguments`: arguments as a `ClarityList` or nil.
// `conditions`: the post-conditions.
// `strict`: has to follow post-condtions.
func NewContractCall(contract *Principal, function string, arguments *ClarityList) (*StacksTransaction, error) {
	if contract == nil {
		return &StacksTransaction{}, errors.New("contract is nil")
	}

	if (*contract.value).Contract == "" {
		return &StacksTransaction{}, errors.New("contract is a standard principal not a contract principal")
	}

	if arguments == nil {
		*arguments.value = clarity.List{}
	}

	payload := transaction.PayloadContractCall{
		Address:   *contract.value,
		Function:  function,
		Arguments: *arguments.value,
	}

	return bind(create(payload)), nil
}

// Create a new contract.
// `account`: creator of the contract.
// `name`: the name of the contract.
// `body`: the contract source code.
// `conditions`: the post-conditions.
// `strict`: has to follow post-condtions.
func NewSmartContract(name string, body string) (*StacksTransaction, error) {
	payload := transaction.PayloadSmartContract{
		Name: name,
		Body: body,
	}

	return bind(create(payload)), nil
}

func conditionPrincipal(from *Principal) transaction.PostConditionPrincipal {
	var result transaction.PostConditionPrincipal

	result.Type = transaction.PostConditionPrincipalTypeOrigin

	if from != nil {
		value := *from.value

		if value.Contract != "" {
			result.Type = transaction.PostConditionPrincipalTypeContract
		} else {
			result.Type = transaction.PostConditionPrincipalTypeStandard
		}

		result.Address = value
	}

	return result
}

// Add a STX PostCondition.
// `code`:
// 		1 is ==
// 		2 is >
// 		3 is >=
// 		4 is <
// 		5 is <=
// `amount`: total uSTX
// `principal`: the destination (or if nil origin is assumed)
func (condition *PostCondition) AddSTX(code int, amount int, principal *Principal) error {
	if code < 1 || code > 5 {
		return errors.New("code is invalid")
	}

	if amount < 0 {
		return errors.New("amount must be > 0")
	}

	*condition.value = append(*condition.value, transaction.PostConditionSTX{
		Condition: transaction.FungibleConditionCode(code),
		Principal: conditionPrincipal(principal),
		Amount:    uint64(amount),
	})

	return nil
}

// Add a FT PostCondition.
// `code`:
// 		1 is ==
// 		2 is >
// 		3 is >=
// 		4 is <
// 		5 is <=
// `amount`: total uSTX
// `principal`: the destination (or if nil origin is assumed)
// `asset`: the fungible token
func (condition *PostCondition) AddFT(code int, amount int, principal *Principal, asset *Asset) error {
	if code < 1 || code > 5 {
		return errors.New("code is invalid")
	}

	if amount < 0 {
		return errors.New("amount must be > 0")
	}

	if asset == nil {
		return errors.New("asset is nil")
	}

	if principal == nil {
		return errors.New("origin principal is not allowed")
	}

	*condition.value = append(*condition.value, transaction.PostConditionFT{
		Condition: transaction.FungibleConditionCode(code),
		Principal: conditionPrincipal(principal),
		Amount:    uint64(amount),
		Asset: transaction.Asset{
			Address: *asset.Contract.value,
			Name:    asset.Name,
		},
	})

	return nil
}

// Add a NFT PostCondition.
// `send`: sending (true) or no sending (false)
// `principal`: the destination (or if nil origin is assumed)
// `asset`: the non-fungible token
// `value`: the identifying clarity value
func (condition *PostCondition) AddNFT(send bool, principal *Principal, asset *Asset, value *ClarityValue) error {
	if asset == nil {
		return errors.New("asset is nil")
	}

	if principal == nil {
		return errors.New("origin principal is not allowed")
	}

	var code transaction.NonFungibleConditionCode

	if send == true {
		code = transaction.NonFungibleConditionCodeSent
	} else {
		code = transaction.NonFungibleConditionCodeNotSent
	}

	*condition.value = append(*condition.value, transaction.PostConditionNFT{
		Condition: code,
		Principal: conditionPrincipal(principal),
		Value:     *value.value,
		Asset: transaction.Asset{
			Address: *asset.Contract.value,
			Name:    asset.Name,
		},
	})

	return nil
}

func (condition *PostCondition) Raw() string {
	return fmt.Sprintf("%#+v\n", *condition.value)
}

func NewPostCondition() *PostCondition {
	return &PostCondition{&[]transaction.PostCondition{}}
}
