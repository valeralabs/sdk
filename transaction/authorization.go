package transaction

import "github.com/valeralabs/sdk/constant"

type SpendingCondition interface {
	GetHashMode() constant.HashMode
	GetSigner() [20]byte
}

type Authorization SpendingCondition

type StandardAuthorization SpendingCondition
type SponsoredAuthorization SpendingCondition

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

type TransactionFieldContents any

func (condition SingleSignatureSpendingCondition) GetHashMode() constant.HashMode {
	return condition.HashMode
}

func (condition SingleSignatureSpendingCondition) GetSigner() [20]byte {
	return condition.Signer
}

func (condition MultipleSignatureSpendingCondition) GetHashMode() constant.HashMode {
	return constant.HashMode(constant.HashModeP2SH)
}

func (condition MultipleSignatureSpendingCondition) GetSigner() [20]byte {
	return condition.Signer
}
