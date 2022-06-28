package transaction

type Payload byte

var (
	PayloadTokenTransfer = Payload(0x00)
	PayloadContractCall  = Payload(0x02)
	PayloadSmartContract = Payload(0x01)
	PayloadPoison        = Payload(0x03)
	PayloadCoinbase      = Payload(0x04)
)

func (payload Payload) Check() bool {
	if payload == PayloadTokenTransfer || payload == PayloadContractCall || payload == PayloadSmartContract || payload == PayloadPoison || payload == PayloadCoinbase {
		return true
	}

	return false
}
