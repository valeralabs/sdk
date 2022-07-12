package transaction

import (
	"github.com/valeralabs/sdk/address"
	"github.com/valeralabs/sdk/constant"
)

type SpendingCondition interface {
	GetKeyEncoding() constant.PublicKeyEncoding
	GetSignature() [65]byte
	SetSignature([65]byte, [20]byte, constant.PublicKeyEncoding) SpendingCondition
	GetHashMode() address.HashMode
	GetSigner() [20]byte
	GetNonce() uint64
	GetFee() uint64
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

func (condition SingleSignatureSpendingCondition) SetSignature(signature [65]byte, signer [20]byte, encoding constant.PublicKeyEncoding) SpendingCondition {
	condition.Signer = signer
	condition.Signature = signature
	condition.KeyEncoding = encoding

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

func (condition MultipleSignatureSpendingCondition) SetSignature(signature [65]byte, signer [20]byte, publicKeyEncoding constant.PublicKeyEncoding) SpendingCondition {
	panic("TODO: Implement multiple signature spending condition adding signatures")
}
