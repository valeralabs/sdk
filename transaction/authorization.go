package transaction

import "github.com/valeralabs/sdk/constant"

type SpendingCondition interface {
	GetKeyEncoding() constant.PublicKeyEncoding
	GetSignature() [65]byte
	GetHashMode() constant.HashMode
	GetSigner() [20]byte
}

type Authorization interface {
	GetCondition() SpendingCondition
}

type StandardAuthorization struct {
	Condition SpendingCondition
}

type SponsoredAuthorization struct {
	Condition   SpendingCondition
	Sponsorship SpendingCondition
}

type SingleSignatureSpendingCondition struct {
	KeyEncoding constant.PublicKeyEncoding
	Signature   [65]byte
	HashMode    constant.HashMode
	Signer      [20]byte
	Nonce       uint64
	Fee         uint64
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

func (authorization StandardAuthorization) GetCondition() SpendingCondition {
	return authorization.Condition
}

func (authorization SponsoredAuthorization) GetCondition() SpendingCondition {
	return authorization.Condition
}

func (condition SingleSignatureSpendingCondition) GetKeyEncoding() constant.PublicKeyEncoding {
	return condition.KeyEncoding
}

func (condition SingleSignatureSpendingCondition) GetSignature() [65]byte {
	return condition.Signature
}

func (condition SingleSignatureSpendingCondition) GetHashMode() constant.HashMode {
	return condition.HashMode
}

func (condition SingleSignatureSpendingCondition) GetSigner() [20]byte {
	return condition.Signer
}

//TODO(MooseMan): implement multiple-signature
func (condition MultipleSignatureSpendingCondition) GetKeyEncoding() constant.PublicKeyEncoding {
	return constant.PublicKeyEncoding(0)
}

func (condition MultipleSignatureSpendingCondition) GetSignature() [65]byte {
	return [65]byte{}
}

func (condition MultipleSignatureSpendingCondition) GetHashMode() constant.HashMode {
	return constant.HashMode(constant.HashModeP2SH)
}

func (condition MultipleSignatureSpendingCondition) GetSigner() [20]byte {
	return condition.Signer
}
