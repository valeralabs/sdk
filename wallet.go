package VDK

import (
	"fmt"

	"valera.co/vdk/address"
	"valera.co/vdk/api"
	"valera.co/vdk/wallet"
)

// Wrapper around [wallet.Wallet].
type Wallet struct {
	value *wallet.Wallet
}

// Wrapper around [wallet.Account].
type Account struct {
	Principal *Principal

	value *wallet.Account
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
