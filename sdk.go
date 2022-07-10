// Provides a mobile-safe wrapper.
package ValeraSDK

//go:generate go vet
//go:generate go run golang.org/x/mobile/cmd/gomobile bind --target=macos .

import (
	"errors"
	"fmt"
	"strings"

	"github.com/valeralabs/sdk/address"
	"github.com/valeralabs/sdk/constant"
	"github.com/valeralabs/sdk/encoding/c32"
	"github.com/valeralabs/sdk/transaction"
	"github.com/valeralabs/sdk/wallet"
)

// Wrapper around [wallet.Wallet].
type Wallet struct {
	value *wallet.Wallet
}

// Wrapper around [wallet.Account].
type Account struct {
	value *wallet.Account
}

// Wrapper around [address.Address].
type Principal struct {
	value *address.Address
}

// Wrapper around [transaction.Payload].
type Payload struct {
	value *transaction.Payload
}

// Wrapper around [transaction.StacksTransaction].
type StacksTransaction struct {
	Version    int
	ChainID    int
	AnchorMode int

	Payload *Payload

	value *transaction.StacksTransaction
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

	return &Account{&account}, nil
}

// Get underlying Wallet as a String.
func (wallet *Wallet) Raw() string {
	return fmt.Sprintf("%#+v\n", *wallet.value)
}

func (account *Account) Principal() (*Principal, error) {
	version := address.AddressVersion{
		HashMode: address.HashModeP2PKH,
		Network:  address.NetworkMainnet,
	}

	principal, err := address.NewAddressSingleSignature((*account.value).PrivateKey.PublicKey(), version)

	if err != nil {
		return &Principal{}, err
	}

	return &Principal{&principal}, nil
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

	principal.Hash, version, err = c32.ChecksumDecode(raw)

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

func bindStacksTransaction(from transaction.StacksTransaction) *StacksTransaction {
	return &StacksTransaction{
		Version:    int(from.Version),
		ChainID:    int(from.ChainID),
		AnchorMode: int(from.AnchorMode),

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

	return bindStacksTransaction(cursor), nil
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

// Create a token transfer.
// `account`: source of funds and signing.
// `recipient`: the destination of the funds can be a standard or contract principal.
// `amount`: total uSTX sent.
// `memo`: optional arbitrary info.
func NewTokenTransfer(account *Account, recipient *Principal, amount int, memo string) (*StacksTransaction, error) {
	if account == nil {
		return &StacksTransaction{}, errors.New("account is nil")
	}

	if recipient == nil {
		return &StacksTransaction{}, errors.New("recipient is nil")
	}

	//TODO: post-conditions
	result := transaction.StacksTransaction{
		Version: constant.TransactionVersionMainnet,
		ChainID: constant.ChainIDMainnet,
		Payload: transaction.PayloadTokenTransfer{
			Address: *recipient.value,
			Amount:  amount,
			Memo:    memo,
		},
		AnchorMode: constant.AnchorModeAny,
	}

	return bindStacksTransaction(result), nil
}

// Create a contract call.
// `account`:  signing.
// `contract`: the principal of the contract.
// `function`: the function being called.
func NewContractCall(account *Account, contract *Principal, function string) (*StacksTransaction, error) {
	if account == nil {
		return &StacksTransaction{}, errors.New("account is nil")
	}

	if contract == nil {
		return &StacksTransaction{}, errors.New("contract is nil")
	}

	if (*contract.value).Contract == "" {
		return &StacksTransaction{}, errors.New("contract is a standard principal not a contract principal")
	}

	//TODO: post-conditions
	result := transaction.StacksTransaction{
		Version: constant.TransactionVersionMainnet,
		ChainID: constant.ChainIDMainnet,
		Payload: transaction.PayloadContractCall{
			Address:  *contract.value,
			Function: function,
		},
		AnchorMode: constant.AnchorModeAny,
	}

	return bindStacksTransaction(result), nil
}

// Create a new contract.
// `account`: creator of the contract.
// `name`: the name of the contract.
// `body`: the contract source code.
func NewSmartContract(account *Account, name string, body string) (*StacksTransaction, error) {
	if account == nil {
		return &StacksTransaction{}, errors.New("account is nil")
	}

	//TODO: post-conditions
	result := transaction.StacksTransaction{
		Version: constant.TransactionVersionMainnet,
		ChainID: constant.ChainIDMainnet,
		Payload: transaction.PayloadSmartContract{
			Name: name,
			Body: body,
		},
		AnchorMode: constant.AnchorModeAny,
	}

	return bindStacksTransaction(result), nil
}
