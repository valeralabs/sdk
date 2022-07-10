// Provides a mobile-safe wrapper.
package ValeraSDK

//go:generate go vet
//go:generate go run golang.org/x/mobile/cmd/gomobile bind --target=macos .

import (
	"fmt"

	"github.com/valeralabs/sdk/address"
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

// Wrapper around [transaction.StacksTransaction].
type StacksTransaction struct {
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
func (wallet *Wallet) DeriveAccount(index int) (*Account, error) {
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

func (account *Account) deriveAddress() (address.Address, error) {
	cursor := *account.value

	version := address.AddressVersion{
		HashMode: address.HashModeP2PKH,
		Network:  address.NetworkMainnet,
	}

	principal, err := address.NewAddressSingleSignature(cursor.PrivateKey.PublicKey(), version)

	if err != nil {
		return address.Address{}, err
	}

	return principal, nil
}

// Derive a Stacks address (C32) from a [Account].
func (account *Account) DeriveStacksAddress() (string, error) {
	principal, err := account.deriveAddress()

	if err != nil {
		return "", err
	}

	c32, err := principal.C32()

	if err != nil {
		return "", err
	}

	return c32, nil
}

// Derive a bitcoin address (B58) from a [Account].
func (account *Account) DeriveBitcoinAddress() (string, error) {
	principal, err := account.deriveAddress()

	if err != nil {
		return "", err
	}

	b58, err := principal.B58()

	if err != nil {
		return "", err
	}

	return b58, nil
}

// Get underlying Account as a String.
func (account *Account) Raw() string {
	return fmt.Sprintf("%#+v\n", *account.value)
}

// Parse a Stacks Transaction (hex encoded).
func ParseStacksTransaction(raw string) (*StacksTransaction, error) {
	var cursor transaction.StacksTransaction

	err := cursor.Unmarshal([]byte(raw))

	if err != nil {
		return &StacksTransaction{}, err
	}

	return &StacksTransaction{&cursor}, nil
}

// Get the underlying Stacks transaction as a String.
func (cursor *StacksTransaction) Raw() string {
	return fmt.Sprintf("%#+v\n", *cursor.value)
}
