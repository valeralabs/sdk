package transaction

import "github.com/valeralabs/sdk/address"

type PayloadType byte

var (
	PayloadTypeTokenTransfer = PayloadType(0x00)
	PayloadTypeContractCall  = PayloadType(0x02)
	PayloadTypeSmartContract = PayloadType(0x01)
	PayloadTypePoison        = PayloadType(0x03)
	PayloadTypeCoinbase      = PayloadType(0x04)
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
