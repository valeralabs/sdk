package VDK

import (
	"errors"
	"strings"

	"valera.co/vdk/api"
)

type BalanceKeys struct {
	FT  []string
	NFT []string
}

type Balance struct {
	keys *BalanceKeys

	value *api.AccountBalance
}

func splitUp(from string) (*Asset, error) {
	split := strings.Split(from, "::")

	principal, err := NewPrincipal(split[0])

	if err != nil {
		return &Asset{}, errors.New("principal not valid")
	}

	asset, err := NewAsset(principal, split[1])

	if err != nil {
		return &Asset{}, errors.New("asset not valid")
	}

	return asset, nil
}

func (balance *Balance) STX() int {
	return (*balance.value).STX.Balance
}

//TODO(Linden): remove once golang/go#13445 is fixed.
func (balance *Balance) FTAt(at int) (*Asset, int, error) {
	index := 0

	for _, raw := range (*balance.keys).FT {
		if index == at {
			total := (*balance.value).FT[raw]

			asset, err := splitUp(raw)

			if err != nil {
				return &Asset{}, 0, err
			}

			return asset, total.Balance, nil
		}

		index++
	}

	return &Asset{}, 0, errors.New("not found")
}

func (balance *Balance) FTLength() int {
	return len((*balance.value).FT)
}

func (balance *Balance) NFTAt(at int) (*Asset, int, error) {
	index := 0

	for _, raw := range (*balance.keys).NFT {
		if index == at {
			total := (*balance.value).NFT[raw]

			asset, err := splitUp(raw)

			if err != nil {
				return &Asset{}, 0, err
			}

			return asset, total.Balance, nil
		}

		index++
	}

	return &Asset{}, 0, errors.New("not found")
}

func (balance *Balance) NFTLength() int {
	return len((*balance.value).NFT)
}

// Retrieve the balance of a user.
func NewBalance(principal *Principal) (*Balance, error) {
	stacks, err := principal.Stacks()

	if err != nil {
		return &Balance{}, err
	}

	balance, err := api.GetBalance(stacks)

	if err != nil {
		return &Balance{}, err
	}

	var keys BalanceKeys

	for key, _ := range balance.FT {
		keys.FT = append(keys.FT, key)
	}

	for key, _ := range balance.NFT {
		keys.NFT = append(keys.NFT, key)
	}

	return &Balance{value: &balance, keys: &keys}, nil
}
