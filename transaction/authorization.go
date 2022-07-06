package transaction

import (
	"github.com/valeralabs/sdk/constant"
	"fmt"
)

type Authorization any

func GetSignerFromAuthorization(auth Authorization) ([20]byte, error) {
	switch auth := any(auth).(type) {
	case StandardAuthorization:
		return GetSignerFromStandardAuth(auth)
	case SponsoredAuthorization:
		return GetSignerFromSponsoredAuth(auth)
	}

	return [20]byte{}, fmt.Errorf("Authorization must be standard authorization or sponsored authorization")
}

func GetSignerHashModeFromAuthorization(auth Authorization) (constant.HashMode, error) {
	switch auth := any(auth).(type) {
	case StandardAuthorization:
		hashMode, err := GetSignerHashModeFromStandardAuth(auth)
		if err != nil {
			return constant.HashMode(0), fmt.Errorf("Could not get signer hash mode from standard auth: %v", err)
		}

		return hashMode, nil
	case SponsoredAuthorization:
		hashMode, err := GetSignerHashModeFromSponsoredAuth(auth)
		if err != nil {
			return constant.HashMode(0), fmt.Errorf("Could not get signer hash mode from standard auth: %v", err)
		}

		return hashMode, nil
	}

	return constant.HashMode(0), fmt.Errorf("Authorization must be StandardAuthorization or SponsoredAuthorization")
}

type SpendingCondition any

type StandardAuthorization SpendingCondition

func GetSignerFromStandardAuth(standardAuth StandardAuthorization) ([20]byte, error) {
	switch standardAuth := any(standardAuth).(type) {
	case SingleSignatureSpendingCondition:
		return standardAuth.Signer, nil
	case MultipleSignatureSpendingCondition:
		return standardAuth.Signer, nil
	}

	return [20]byte{}, fmt.Errorf("Standard Authorization must be single signature spending condition or multiple signature spending condition")
}

func GetSignerHashModeFromStandardAuth(standardAuth StandardAuthorization) (constant.HashMode, error) {
	switch standardAuth := any(standardAuth).(type) {
	case SingleSignatureSpendingCondition:
		return standardAuth.HashMode, nil
	case MultipleSignatureSpendingCondition:
		// TODO: Change to proper hash mode once MultipleSignatureSpendingCondition is implemented
		return constant.HashModeP2SH, nil
	}

	return constant.HashMode(0), fmt.Errorf("Standard Authorization must be single signature spending condition or multiple signature spending condition")
}

type SponsoredAuthorization SpendingCondition

func GetSignerFromSponsoredAuth(sponsoredAuth SponsoredAuthorization) ([20]byte, error) {
	switch sponsoredAuth := any(sponsoredAuth).(type) {
	case SingleSignatureSpendingCondition:
		return sponsoredAuth.Signer, nil
	case MultipleSignatureSpendingCondition:
		return sponsoredAuth.Signer, nil
	}

	return [20]byte{}, fmt.Errorf("Sponsored Authorization must be single signature spending condition or multiple signature spending condition")
}

func GetSignerHashModeFromSponsoredAuth(sponsoredAuth SponsoredAuthorization) (constant.HashMode, error) {
	switch sponsoredAuth := any(sponsoredAuth).(type) {
	case SingleSignatureSpendingCondition:
		return sponsoredAuth .HashMode, nil
	case MultipleSignatureSpendingCondition:
		// TODO: Change to proper hash mode once MultipleSignatureSpendingCondition is implemented
		return constant.HashModeP2SH, nil
	}

	return constant.HashMode(0), fmt.Errorf("Standard Authorization must be single signature spending condition or multiple signature spending condition")
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
type TransactionFieldContents any
