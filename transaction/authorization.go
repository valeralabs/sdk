package transaction

import (
	"github.com/valeralabs/sdk/constant"
)

type Authorization interface {
	GetFee() int
	GetSigner() [20]byte
	GetSignerHashMode() constant.HashMode
}

type StandardAuthorization[T SpendingCondition] struct {
	SpendingCondition T
}

type SponsoredAuthorization[T SpendingCondition] struct {
	SpendingCondition T
}

type SpendingCondition interface {
	SingleSignatureSpendingCondition | MultipleSignatureSpendingCondition
	GetSigner() [20]byte
	GetSignerHashMode() constant.HashMode
}

type SingleSignatureSpendingCondition struct {
	HashMode    constant.HashMode
	Signer      [20]byte
	Nonce       uint64
	Fee         uint64
	KeyEncoding constant.PublicKeyEncoding
	Signature   [65]byte
}

func (condition SingleSignatureSpendingCondition) GetSigner() [20]byte {
	return condition.Signer
}

func (condition SingleSignatureSpendingCondition) GetSignerHashMode() constant.HashMode {
	return condition.HashMode
}

type MultipleSignatureSpendingCondition struct {
	Signer             [20]byte
	Nonce              uint64
	Fee                uint64
	Fields             []AuthorizationField
	signaturesRequired int
}

func (condition MultipleSignatureSpendingCondition) GetSigner() [20]byte {
	return condition.Signer
}

func (condition MultipleSignatureSpendingCondition) GetSignerHashMode() constant.HashMode {
	// TODO Handle P2WSH
	return constant.HashModeP2SH
}

type AuthorizationField struct {
	KeyEncoding constant.PublicKeyEncoding
	Contents    TransactionFieldContents
}

type TransactionFieldContents interface{}

func (authorization StandardAuthorization[SpendingCondition]) GetFee() int {
	return 5000
}

func (authorization StandardAuthorization[SpendingCondition]) GetSigner() [20]byte {
	return authorization.SpendingCondition.GetSigner()
}

func (authorization StandardAuthorization[SpendingCondition]) GetSignerHashMode() constant.HashMode {
	return authorization.SpendingCondition.GetSignerHashMode()
}

func (authorization SponsoredAuthorization[SpendingCondition]) GetFee() int {
	return 5000
}

func (authorization SponsoredAuthorization[SpendingCondition]) GetSigner() [20]byte {
	return authorization.SpendingCondition.GetSigner()
}

func (authorization SponsoredAuthorization[SpendingCondition]) GetSignerHashMode() constant.HashMode {
	return authorization.SpendingCondition.GetSignerHashMode()
}
