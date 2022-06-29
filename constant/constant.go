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
