package transaction

import (
	"valera.co/vdk/address"
	"valera.co/vdk/constant"
)

type SpendingCondition interface {
	GetKeyEncoding() constant.PublicKeyEncoding
	GetSignature() [65]byte
	SetSignature([65]byte)
	GetHashMode() address.HashMode
	GetSigner() [20]byte
	SetSigner([20]byte)
	SetNonce(uint64)
	GetNonce() uint64
	SetFee(uint64)
	GetFee() uint64
	Clear() SpendingCondition
}

//TODO(Linden): simplify with a pointer?
type Authorization interface {
	SetCondition(SpendingCondition) Authorization
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

func (authorization StandardAuthorization) SetCondition(condition SpendingCondition) Authorization {
	authorization.Condition = condition
	return authorization
}

func (authorization SponsoredAuthorization) GetCondition() SpendingCondition {
	return authorization.Condition
}

func (authorization SponsoredAuthorization) SetCondition(condition SpendingCondition) Authorization {
	authorization.Condition = condition
	return authorization
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

func (condition *SingleSignatureSpendingCondition) SetSigner(signer [20]byte) {
	condition.Signer = signer
}

func (condition SingleSignatureSpendingCondition) GetFee() uint64 {
	return condition.Fee
}

func (condition SingleSignatureSpendingCondition) GetNonce() uint64 {
	return condition.Nonce
}

func (condition *SingleSignatureSpendingCondition) SetFee(fee uint64) {
	condition.Fee = fee
}

func (condition *SingleSignatureSpendingCondition) SetNonce(nonce uint64) {
	condition.Nonce = nonce
}

func (condition SingleSignatureSpendingCondition) Clear() SpendingCondition {
	condition.Fee = 0
	condition.Nonce = 0
	condition.Signature = [65]byte{}

	return &condition
}

func (condition *SingleSignatureSpendingCondition) SetSignature(signature [65]byte) {
	condition.Signature = signature
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

func (condition *MultipleSignatureSpendingCondition) SetFee(fee uint64) {
	condition.Fee = fee
}

func (condition *MultipleSignatureSpendingCondition) SetNonce(nonce uint64) {
	condition.Nonce = nonce
}

func (condition MultipleSignatureSpendingCondition) Clear() SpendingCondition {
	panic("TODO: implement multiple signatures")
}

func (condition MultipleSignatureSpendingCondition) SetSignature(signature [65]byte) SpendingCondition {
	panic("TODO: implement multiple signature spending condition adding signatures")
}
