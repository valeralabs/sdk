package constant

const MaxStringLength = 128
const DefaultPrefixLength = 4

type TransactionVersion byte

const (
	TransactionVersionMainnet TransactionVersion = 0x00
	TransactionVersionTestnet TransactionVersion = 0x80
)

func (version TransactionVersion) Check() bool {
	if version == TransactionVersionMainnet || version == TransactionVersionTestnet {
		return true
	}

	return false
}

type ChainID [4]byte

var (
	// stacks 2.0 was 0x00000000
	// stacks 2.05 started a soft fork using 0x00000001
	ChainIDMainnet = ChainID([4]byte{00, 00, 00, 01})
	ChainIDTestnet = ChainID([4]byte{80, 00, 00, 00})
)

func (chainID ChainID) Check() bool {
	return chainID == ChainIDMainnet || chainID == ChainIDTestnet
}

type AnchorMode byte

const (
	AnchorModeOnChainOnly AnchorMode = iota + 1
	AnchorModeOffChainOnly
	AnchorModeAny
)

func (mode AnchorMode) Check() bool {
	return mode >= AnchorModeOnChainOnly && mode <= AnchorModeAny
}

type PublicKeyEncoding byte

const (
	PubKeyEncodingCompressed PublicKeyEncoding = iota
	PubKeyEncodingUncompressed
)

func (encoding PublicKeyEncoding) Check() bool {
	return encoding == PubKeyEncodingCompressed || encoding == PubKeyEncodingUncompressed
}

type AddressVersion int

const (
	AddressVersionMainnetSingleSignature AddressVersion = iota
	AddressVersionTestnetSingleSignature
	AddressVersionMainnetMultipleSignature
	AddressVersionTestnetMultipleSignature
)

func (version AddressVersion) Check() bool {
	return version >= AddressVersionMainnetSingleSignature && version <= AddressVersionTestnetMultipleSignature
}

type HashMode byte

const (
	HashModeP2PKH HashMode = iota
	HashModeP2SH
	HashModeP2WPKH
	HashModeP2WSH
)

func (mode HashMode) Check() bool {
	return mode >= HashModeP2PKH && mode <= HashModeP2WSH
}

type PostConditionMode byte

const (
	PostConditionModeLoose PostConditionMode = iota + 1
	PostConditionModeStrict
)

func (mode PostConditionMode) Check() bool {
	return mode == PostConditionModeLoose || mode == PostConditionModeStrict
}

type PostConditionType byte

const (
	PostConditionTypeSTX PostConditionType = iota
	PostConditionTypeFT
	PostConditionTypeNFT
)

func (_type PostConditionType) Check() bool {
	return _type >= PostConditionTypeSTX && _type <= PostConditionTypeNFT
}

type PostCondition byte

func (mode PostCondition) Check() bool {
	return true
}

type ClarityType byte

const (
	ClarityTypeInt ClarityType = iota
	ClarityTypeUInt
	ClarityTypeBuffer
	ClarityTypeBoolTrue
	ClarityTypeBoolFalse
	ClarityTypePrincipalStandard
	ClarityTypePrincipalContract
	ClarityTypeResponseOk
	ClarityTypeResponseErr
	ClarityTypeOptionalNone
	ClarityTypeOptionalSome
	ClarityTypeList
	ClarityTypeTuple
	ClarityTypeStringASCII
	ClarityTypeStringUTF8
)

func (_type ClarityType) Check() bool {
	return _type >= ClarityTypeInt && _type <= ClarityTypeStringUTF8
}
