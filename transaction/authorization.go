package transaction

import "github.com/valeralabs/sdk/constant"

type Authorization interface {
	GetFee() int
}

type StandardAuthorization[T SpendingCondition] struct {
	SpendingCondition T
}

type SponsoredAuthorization[T SpendingCondition] struct {
	SpendingCondition T
}

type SpendingCondition interface {
	SingleSignatureSpendingCondition | MultipleSignatureSpendingCondition
}

type SingleSignatureSpendingCondition struct {
	HashMode    constant.HashMode
	Signer      [20]byte
	Nonce       uint64
	Fee         uint64
	KeyEncoding constant.PublicKeyEncoding
	Signature   [65]byte
}

type MultipleSignatureSpendingCondition struct {
	Signer             [20]byte
	Nonce              uint64
	Fee                uint64
	Fields             []AuthorizationField
	signaturesRequired int
}

type AuthorizationField struct {
	KeyEncoding constant.PublicKeyEncoding
	Contents    TransactionFieldContents
}

type TransactionFieldContents interface{}

func (authorization StandardAuthorization[SpendingCondition]) GetFee() int {
	return 5000
}

func (authorization SponsoredAuthorization[SpendingCondition]) GetFee() int {
	return 5000
}
