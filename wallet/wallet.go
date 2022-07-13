package wallet

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/tyler-smith/go-bip39"

	"github.com/valeralabs/sdk/wallet/keys"
)

const HardenedOffset = 0x80000000

var (
	DataDerivationPath          = []uint32{888 + HardenedOffset, 0 + HardenedOffset}
	AccountDerivationPath       = []uint32{44 + HardenedOffset, 5757 + HardenedOffset, 0 + HardenedOffset, 0}
	ConfigurationDerivationPath = []uint32{44, 5757 + HardenedOffset, 0 + HardenedOffset, 1}
)

type Wallet struct {
	Root *hdkeychain.ExtendedKey

	Salt          []byte
	Configuration []byte
}

type Account struct {
	Root *hdkeychain.ExtendedKey

	Keychain *hdkeychain.ExtendedKey

	AppPrivateKey  *btcec.PrivateKey
	DataPrivateKey []byte

	PrivateKey *keys.PrivateKey
}

func Derive(Root *hdkeychain.ExtendedKey, path []uint32) (*hdkeychain.ExtendedKey, error) {
	var err error

	key := Root

	for _, cursor := range path {
		key, err = key.Derive(cursor)

		if err != nil {
			return nil, err
		}
	}

	return key, nil
}

func (account *Account) getDataPrivateKey() ([]byte, error) {
	private, err := account.Keychain.ECPrivKey()

	if err != nil {
		return []byte{}, err
	}

	serialized := private.Serialize()

	hexed := make([]byte, hex.EncodedLen(len(serialized)))
	hex.Encode(hexed, serialized)

	return hexed, nil
}

func (account *Account) getAppPrivateKey() (*btcec.PrivateKey, error) {
	key, err := Derive(account.Keychain, []uint32{HardenedOffset})

	if err != nil {
		return nil, err
	}

	private, err := key.ECPrivKey()

	if err != nil {
		return nil, err
	}

	return private, nil
}

func (account *Account) getPrivateKey() (*keys.PrivateKey, error) {
	private, err := account.Root.ECPrivKey()

	if err != nil {
		return nil, err
	}

	return &keys.PrivateKey{private, true}, nil
}

func (wallet *Wallet) getSalt() ([]byte, error) {
	key, err := Derive(wallet.Root, DataDerivationPath)

	if err != nil {
		return []byte{}, err
	}

	public, err := key.ECPubKey()

	if err != nil {
		return []byte{}, err
	}

	serialized := public.SerializeCompressed()

	length := hex.EncodedLen(len(serialized))
	first := make([]byte, length)

	actual := hex.Encode(first, serialized)

	if actual != length {
		return []byte{}, errors.New("failed to hex encode (1st round)")
	}

	hashed := sha256.New()
	hashed.Write(first)

	hash := hashed.Sum(nil)

	length = hex.EncodedLen(len(hash))
	second := make([]byte, length)

	actual = hex.Encode(second, hash)

	if actual != length {
		return []byte{}, errors.New("failed to hex encode (2nd round)")
	}

	return second, nil
}

func (wallet *Wallet) getConfiguration() ([]byte, error) {
	key, err := Derive(wallet.Root, ConfigurationDerivationPath)

	if err != nil {
		return []byte{}, err
	}

	private, err := key.ECPrivKey()

	if err != nil {
		return []byte{}, err
	}

	serialized := private.Serialize()

	length := hex.EncodedLen(len(serialized))
	hexed := make([]byte, length)

	actual := hex.Encode(hexed, serialized)

	if length != actual {
		return []byte{}, errors.New("failed to hex encode")
	}

	return hexed, nil
}

func (wallet *Wallet) Account(index int) (Account, error) {
	var err error
	var account Account

	account.Root, err = Derive(wallet.Root, append(AccountDerivationPath, uint32(index)))

	if err != nil {
		return Account{}, err
	}

	account.Keychain, err = Derive(wallet.Root, append(DataDerivationPath, uint32(index+HardenedOffset)))

	if err != nil {
		return Account{}, err
	}

	account.DataPrivateKey, err = account.getDataPrivateKey()

	if err != nil {
		return Account{}, err
	}

	account.AppPrivateKey, err = account.getAppPrivateKey()

	if err != nil {
		return Account{}, err
	}

	account.PrivateKey, err = account.getPrivateKey()

	if err != nil {
		return Account{}, err
	}

	return account, nil
}

func NewWallet(seed []byte) (Wallet, error) {
	Root, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)

	if err != nil {
		return Wallet{}, err
	}

	wallet := Wallet{
		Root: Root,
	}

	wallet.Salt, err = wallet.getSalt()

	if err != nil {
		return Wallet{}, err
	}

	wallet.Configuration, err = wallet.getConfiguration()

	if err != nil {
		return Wallet{}, err
	}

	return wallet, nil
}

func NewSeed(phrase string, password string) ([]byte, error) {
	if fits((32 * len(strings.Split(phrase, " ")) / 3)) == false {
		return []byte{}, errors.New("phrase size is invalid")
	}

	return bip39.NewSeed(phrase, password), nil
}

func NewPhrase(length int) (string, error) {
	size := 32 * length / 3
	entropy, err := bip39.NewEntropy(size)

	if err != nil {
		return "", err
	}

	phrase, err := bip39.NewMnemonic(entropy)

	if err != nil {
		return "", err
	}

	if len(strings.Split(phrase, " ")) != length {
		return "", errors.New("phrase did not match required length")
	}

	return phrase, nil
}

func fits(size int) bool {
	return (size%32) == 0 && size >= 128 && size <= 256
}
