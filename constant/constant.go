package constant

type TransactionVersion byte

const (
	TransactionVersionMainnet = TransactionVersion(byte(0x00))
	TransactionVersionTestnet = TransactionVersion(byte(0x80))
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

var (
	AnchorModeOnChainOnly  = AnchorMode(byte(0x01))
	AnchorModeOffChainOnly = AnchorMode(byte(0x02))
	AnchorModeAny          = AnchorMode(byte(0x03))
)

func (mode AnchorMode) Check() bool {
	return mode == AnchorModeOnChainOnly || mode == AnchorModeOffChainOnly || mode == AnchorModeAny
}

type PublicKeyEncoding byte

var (
	PubKeyEncodingCompressed   = PublicKeyEncoding(byte(0x00))
	PubKeyEncodingUncompressed = PublicKeyEncoding(byte(0x01))
)

func (encoding PublicKeyEncoding) Check() bool {
	return encoding == PubKeyEncodingCompressed || encoding == PubKeyEncodingUncompressed
}

const MaxStringLength = 128
const DefaultPrefixLength = 4

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

var (
	HashModeP2PKH  = HashMode(byte(0x00))
	HashModeP2SH   = HashMode(byte(0x01))
	HashModeP2WPKH = HashMode(byte(0x02))
	HashModeP2WSH  = HashMode(byte(0x03))
)

func (mode HashMode) Check() bool {
	return mode >= HashModeP2PKH && mode <= HashModeP2WSH
}

type PostConditionMode byte

var (
	PostConditionModeLoose  = PostConditionMode(byte(0x01))
	PostConditionModeStrict = PostConditionMode(byte(0x02))
)

func (mode PostConditionMode) Check() bool {
	return mode == PostConditionModeLoose || mode == PostConditionModeStrict
}

type PostConditionType byte

var (
	PostConditionTypeSTX = PostConditionType(byte(0x00))
	PostConditionTypeFT  = PostConditionType(byte(0x01))
	PostConditionTypeNFT = PostConditionType(byte(0x02))
)

func (_type PostConditionType) Check() bool {
	return _type >= PostConditionTypeSTX && _type <= PostConditionTypeNFT
}

type PostCondition byte

func (mode PostCondition) Check() bool {
	return true
}

type ClarityType byte

var (
	ClarityTypeInt               = ClarityType(byte(0x00))
	ClarityTypeUInt              = ClarityType(byte(0x01))
	ClarityTypeBuffer            = ClarityType(byte(0x02))
	ClarityTypeBoolTrue          = ClarityType(byte(0x03))
	ClarityTypeBoolFalse         = ClarityType(byte(0x04))
	ClarityTypePrincipalStandard = ClarityType(byte(0x05))
	ClarityTypePrincipalContract = ClarityType(byte(0x06))
	ClarityTypeResponseOk        = ClarityType(byte(0x07))
	ClarityTypeResponseErr       = ClarityType(byte(0x08))
	ClarityTypeOptionalNone      = ClarityType(byte(0x09))
	ClarityTypeOptionalSome      = ClarityType(byte(0x0a))
	ClarityTypeList              = ClarityType(byte(0x0b))
	ClarityTypeTuple             = ClarityType(byte(0x0c))
	ClarityTypeStringASCII       = ClarityType(byte(0x0d))
	ClarityTypeStringUTF8        = ClarityType(byte(0x0e))
)

func (_type ClarityType) Check() bool {
	return _type >= ClarityTypeInt && _type <= ClarityTypeStringUTF8
}
