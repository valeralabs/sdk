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
	HashMode    constant.HashMode
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

func (authorization StandardAuthorization) GetFee() int {
	return 5000
}

func (authorization SponsoredAuthorization) GetFee() int {
	return 5000
}
