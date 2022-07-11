package transaction

import (
	"github.com/valeralabs/sdk/constant"
	"github.com/valeralabs/sdk/address"
)

type SpendingCondition interface {
	GetKeyEncoding() constant.PublicKeyEncoding
	GetSignature() [65]byte
	GetHashMode() address.HashMode
	GetSigner() [20]byte
	GetNonce() uint64
	GetFee() uint64
	WithAddedSignature(signature [65]byte, publicKeyEncoding constant.PublicKeyEncoding) SpendingCondition
	Cleared() SpendingCondition
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
	HashMode    address.HashMode
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

func (condition SingleSignatureSpendingCondition) GetHashMode() address.HashMode {
	return condition.HashMode
}

func (condition SingleSignatureSpendingCondition) GetSigner() [20]byte {
	return condition.Signer
}

func (condition SingleSignatureSpendingCondition) GetFee() uint64 {
	return condition.Fee
}

func (condition SingleSignatureSpendingCondition) GetNonce() uint64 {
	return condition.Nonce
}

func (condition SingleSignatureSpendingCondition) WithAddedSignature(signature [65]byte, publicKeyEncoding constant.PublicKeyEncoding) SpendingCondition {
	condition.Signature = [65]byte{}
	condition.KeyEncoding = publicKeyEncoding
	return condition
}

func (condition SingleSignatureSpendingCondition) Cleared() SpendingCondition {
	condition.Fee = 0
	condition.Nonce = 0
	condition.Signature = [65]byte{}
	return condition
}

//TODO(MooseMan): implement multiple-signature
func (condition MultipleSignatureSpendingCondition) GetKeyEncoding() constant.PublicKeyEncoding {
	return constant.PublicKeyEncoding(0)
}

func (condition MultipleSignatureSpendingCondition) GetSignature() [65]byte {
	return [65]byte{}
}

func (condition MultipleSignatureSpendingCondition) GetHashMode() address.HashMode {
	return address.HashMode(address.HashModeP2SH)
}

func (condition MultipleSignatureSpendingCondition) GetSigner() [20]byte {
	return condition.Signer
}

func (condition MultipleSignatureSpendingCondition) GetFee() uint64 {
	return condition.Fee
}

func (condition MultipleSignatureSpendingCondition) GetNonce() uint64 {
	return condition.Nonce
}

func (condition MultipleSignatureSpendingCondition) WithAddedSignature(signature [65]byte, publicKeyEncoding constant.PublicKeyEncoding) MultipleSignatureSpendingCondition {
	panic("TODO: Implement multiple signature spending condition adding signatures")
}

func (condition MultipleSignatureSpendingCondition) Cleared() MultipleSignatureSpendingCondition {
	condition.Fee = 0
	condition.Nonce = 0

	panic("TODO: Implement multiple signature spending condition clearing")
}
