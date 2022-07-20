// Provides a mobile-safe wrapper.
package VDK

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"valera.co/vdk/address"
	"valera.co/vdk/api"
	"valera.co/vdk/constant"
	"valera.co/vdk/encoding/c32"
	"valera.co/vdk/encoding/clarity"
	"valera.co/vdk/transaction"
	"valera.co/vdk/wallet"
)

// Wrapper around [wallet.Wallet].
type Wallet struct {
	value *wallet.Wallet
}

// Wrapper around [address.Address].
type Principal struct {
	value *address.Address
}

// Wrapper around [wallet.Account].
type Account struct {
	Principal *Principal

	value *wallet.Account
}

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

// Wrapper around [encoding/clarity.Value].
type ClarityValue struct {
	value *clarity.Value
}

// Wrapper around [encoding/clarity.List].
type ClarityList struct {
	value *clarity.List
}

type Asset struct {
	Contract *Principal
	Name     string
}

// Derive a wallet from your mnemonic seed phrase with the option of using a password.
// This implements [BIP39](https://github.com/bitcoin/bips/blob/master/bip-0039/bip-0039-wordlists.md).
// Note: Hiro wallet doesn't use a password only for encryption.
func NewWalletFromPhrase(phrase string, password string) (*Wallet, error) {
	seed, err := wallet.NewSeed(phrase, password)

	if err != nil {
		return &Wallet{}, err
	}

	created, err := wallet.NewWallet(seed)

	if err != nil {
		return &Wallet{}, err
	}

	return &Wallet{&created}, nil
}

// Generate a new mnemonic seed phrase for use with [NewWalletFromMnemonic].
// We recommoned a length of 24 for maximum security and compatibility.
func NewPhrase(length int) (string, error) {
	if length <= 0 {
		length = 24
	}

	return wallet.NewPhrase(length)
}

// Derive a Stacks account from a [Wallet], index starting at 0.
func (wallet *Wallet) Account(index int) (*Account, error) {
	cursor := *wallet.value

	account, err := cursor.Account(index)

	if err != nil {
		return &Account{}, err
	}

	version := address.AddressVersion{
		HashMode: address.HashModeP2PKH,
		Network:  address.NetworkMainnet,
	}

	principal, err := address.NewAddressSingleSignature(account.PrivateKey.PublicKey(), version)

	if err != nil {
		return &Account{}, err
	}

	return &Account{
		Principal: &Principal{&principal},

		value: &account,
	}, nil
}

// Get underlying Wallet as a String.
func (wallet *Wallet) Raw() string {
	return fmt.Sprintf("%#+v\n", *wallet.value)
}

// Derive a Stacks address (C32) from a [Principal].
func (principal *Principal) Stacks() (string, error) {
	c32, err := (*principal.value).C32()

	if err != nil {
		return "", err
	}

	return c32, nil
}

// Derive a bitcoin address (B58) from a [Principal].
func (principal *Principal) Bitcoin() (string, error) {
	b58, err := (*principal.value).B58()

	if err != nil {
		return "", err
	}

	return b58, nil
}

// Create a principal from a Stacks (c32) address.
func NewPrincipal(value string) (*Principal, error) {
	var principal address.Address
	var err error

	split := strings.Split(value, ".")
	value = split[0]

	if len(split) > 1 {
		principal.Contract = split[1]
	}

	raw := []byte(value)
	raw = raw[1:]

	var version byte

	hexed, version, err := c32.ChecksumDecode(raw)

	hash := make([]byte, hex.DecodedLen(len(hexed)))
	hex.Decode(hash, hexed)

	principal.Hash = hash

	if err != nil {
		return &Principal{}, err
	}

	principal.Version, err = address.ReverseStacksVersion(version)

	if err != nil {
		return &Principal{}, err
	}

	return &Principal{&principal}, nil
}

// Get underlying Account as a String.
func (account *Account) Raw() string {
	return fmt.Sprintf("%#+v\n", *account.value)
}

// Get the *possible* next nonce.
func (account *Account) NextNonce() (int, error) {
	principal, err := account.Principal.Stacks()

	if err != nil {
		return 0, err
	}

	return api.NextNonce(principal)
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

// Parse a Stacks Transaction (hex encoded).
func ParseStacksTransaction(raw string) (*StacksTransaction, error) {
	var cursor transaction.StacksTransaction

	err := cursor.Unmarshal([]byte(raw))

	if err != nil {
		return &StacksTransaction{}, err
	}

	return bind(cursor), nil
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
// `account`: the account used to sign the transaction.
func (stacks *StacksTransaction) Broadcast(account *Account) error {
	if account == nil {
		return errors.New("account is nil")
	}

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
// `asset`: name of the token
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

func (condition *PostCondition) Raw() string {
	return fmt.Sprintf("%#+v\n", *condition.value)
}

func NewPostCondition() *PostCondition {
	return &PostCondition{&[]transaction.PostCondition{}}
}

func (value *ClarityValue) Raw() string {
	return fmt.Sprintf("%#+v\n", *value.value)
}

func (value *ClarityValue) SetPrefix(prefix int) {
	(*value.value).PrefixLength = prefix
}

// `base`:
// 	  0  is Int
// 	  1  is UInt
// 	  2  is Buffer
// 	  3  is BoolTrue
// 	  4  is BoolFalse
// 	  5  is PrincipalStandard
// 	  6  is PrincipalContract
// 	  7  is ResponseOk
// 	  8  is ResponseErr
// 	  9  is OptionalNone
// 	  10 is OptionalSome
// 	  11 is List
// 	  12 is Tuple
// 	  13 is StringASCII
// 	  14 is StringUTF8
// `content`: value
func NewClarityValue(base int, content string) (*ClarityValue, error) {
	typed := clarity.ClarityType(base)

	if typed.Check() == false {
		return &ClarityValue{}, errors.New("base is not a clarity type")
	}

	return &ClarityValue{&clarity.Value{
		Type:         typed,
		Content:      []byte(content),
		PrefixLength: constant.DefaultPrefixLength,
	}}, nil
}

func (list *ClarityList) SetPrefix(prefix int) {
	(*list.value).PrefixLength = prefix
}

func (list *ClarityList) SetSubPrefix(prefix int) {
	(*list.value).SubPrefixLength = prefix
}

func (list *ClarityList) Add(value *ClarityValue) {
	list.value.Content = append(list.value.Content, *value.value)
}

func (list *ClarityList) Raw() string {
	return fmt.Sprintf("%#+v\n", *list.value)
}

func NewClarityList() *ClarityList {
	return &ClarityList{&clarity.List{
		PrefixLength:    constant.DefaultPrefixLength,
		SubPrefixLength: constant.DefaultPrefixLength,
	}}
}

func NewAsset(contract *Principal, name string) (*Asset, error) {
	if contract == nil {
		return &Asset{}, errors.New("contract is nil")
	}

	if (*contract.value).Contract == "" {
		return &Asset{}, errors.New("contract is a standard principal")
	}

	if name == "" {
		return &Asset{}, errors.New("asset name is invalid")
	}

	return &Asset{
		Contract: contract,
		Name:     name,
	}, nil
}
