package transaction

import "github.com/valeralabs/sdk/address"

type PayloadType byte

const (
	PayloadTypeTokenTransfer PayloadType = iota
	PayloadTypeContractCall
	PayloadTypeSmartContract
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
