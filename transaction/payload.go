package transaction

import (
	"github.com/valeralabs/sdk/address"
	"github.com/valeralabs/sdk/encoding/clarity"
)

type PayloadType byte

const (
	PayloadTypeTokenTransfer PayloadType = iota
	PayloadTypeSmartContract
	PayloadTypeContractCall
	PayloadTypePoison
	PayloadTypeCoinbase
)

func (payload PayloadType) Check() bool {
	return payload >= PayloadTypeTokenTransfer && payload <= PayloadTypeCoinbase
}

type Payload any

type PayloadTokenTransfer struct {
	Address address.Address
	Amount  int
	Memo    string
}

type PayloadSmartContract struct {
	Name string
	Body string
}

type ContractCallTransfer struct {
	Address   address.Address
	Function  string
	Arguments clarity.List
}
