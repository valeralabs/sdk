package transaction

import "github.com/valeralabs/sdk/constant"

type Authorization any
type SpendingCondition any

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
