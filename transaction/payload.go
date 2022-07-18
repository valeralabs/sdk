package transaction

import (
	"valera.co/vdk/address"
	"valera.co/vdk/encoding/clarity"
)

//go:generate go run golang.org/x/tools/cmd/stringer@latest -type=PayloadType -output=payload_string.go

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

type PayloadContractCall struct {
	Address   address.Address
	Function  string
	Arguments clarity.List
}
