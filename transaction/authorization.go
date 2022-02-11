package transaction

import "github.com/syvita/racks/constant"

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
	Signer      string
	Nonce       int
	Fee         int
	KeyEncoding constant.PublicKeyEncoding
	Signature   string
}

//TODO: check hashMode == MultiSigHashMode
type MultiSigSpendingCondition struct {
	Signer             string
	Nonce              int
	Fee                int
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
