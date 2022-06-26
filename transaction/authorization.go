package transaction

import "github.com/valeralabs/racks/constant"

type Authorization interface {
	GetFee() int
}

type StandardAuthorization struct {
	SpendingCondition SpendingCondition
}

type SponsoredAuthorization struct {
	SpendingCondition SpendingCondition
}

type SpendingCondition interface{}

//TODO: check hashMode == SingleSigHashMode
type SingleSigSpendingCondition struct {
	HashMode    HashMode
	Signer      [20]byte
	Nonce       uint64
	Fee         uint64
	KeyEncoding constant.PublicKeyEncoding
	Signature   [65]byte
}

//TODO: check hashMode == MultiSigHashMode
type MultiSigSpendingCondition struct {
	Signer             [20]byte
	Nonce              uint64
	Fee                uint64
	Fields             []AuthorizationField
	signaturesRequired int
}

//TODO: check type == StacksMessageType.TransactionAuthField
type AuthorizationField struct {
	KeyEncoding constant.PublicKeyEncoding
	Contents    TransactionFieldContents
}

type TransactionFieldContents interface{}

type HashMode byte

var (
	HashModeP2PKH       = HashMode(byte(0x00))
	HashModeP2SH        = HashMode(byte(0x01))
	HashModeP2WPKH_P2SH = HashMode(byte(0x02))
	HashModeP2WSH_P2SH  = HashMode(byte(0x03))
)

func (mode HashMode) Check() bool {
	if mode == HashModeP2PKH || mode == HashModeP2SH || mode == HashModeP2WPKH_P2SH || mode == HashModeP2WSH_P2SH {
		return true
	}

	return false
}

func (authorization StandardAuthorization) GetFee() int {
	return 5000
}

func (authorization SponsoredAuthorization) GetFee() int {
	return 5000
}
