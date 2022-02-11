package constant

type TransactionVersion byte

const (
	TransactionVersionMainnet = TransactionVersion(byte(00))
	TransactionVersionTestnet = TransactionVersion(byte(80))
)

func (version TransactionVersion) Check() bool {
	if version == TransactionVersionMainnet || version == TransactionVersionTestnet {
		return true
	}

	return false
}

type ChainID [4]byte

var (
	ChainIDMainnet = ChainID([4]byte{00, 00, 00, 01})
	ChainIDTestnet = ChainID([4]byte{80, 00, 00, 00})
)

func (chainID ChainID) Check() bool {
	if chainID == ChainIDMainnet || chainID == ChainIDTestnet {
		return true
	}

	return false
}

type PostConditionMode byte

var (
	PostConditionModeAllow = PostConditionMode(byte(0x01))
	PostConditionModeDeny  = PostConditionMode(byte(0x02))
)

func (postConditionMode PostConditionMode) Check() bool {
	if postConditionMode == PostConditionModeAllow || postConditionMode == PostConditionModeDeny {
		return true
	}

	return false
}

//TODO (Linden): figure out what can be inside this
type PostCondition struct{}

type AnchorMode byte

var (
	AnchorModeOnChainOnly  = AnchorMode(byte(0x01))
	AnchorModeOffChainOnly = AnchorMode(byte(0x02))
	AnchorModeAny          = AnchorMode(byte(0x03))
)

func (mode AnchorMode) Check() bool {
	if mode == AnchorModeOnChainOnly || mode == AnchorModeOffChainOnly || mode == AnchorModeAny {
		return true
	}

	return false
}

type PublicKeyEncoding byte

var (
	PubKeyEncodingCompressed   = PublicKeyEncoding(byte(0x00))
	PubKeyEncodingUncompressed = PublicKeyEncoding(byte(0x01))
)

func (encoding PublicKeyEncoding) Check() bool {
	if encoding == PubKeyEncodingCompressed || encoding == PubKeyEncodingUncompressed {
		return true
	}

	return false
}
