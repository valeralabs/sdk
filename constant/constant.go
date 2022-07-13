package constant

//go:generate go run golang.org/x/tools/cmd/stringer@latest -type=ChainID,TransactionVersion,AnchorMode,PublicKeyEncoding -output=constant_string.go

const MaxStringLength = 128
const DefaultPrefixLength = 4
const MemoLengthBytes = 32

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

type ChainID int

const (
	// stacks 2.0 was 0x00000000
	// stacks 2.05 started a soft fork using 0x00000001
	ChainIDMainnet ChainID = 0x00000001
	ChainIDTestnet ChainID = 0x80000000
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
	PublicKeyEncodingCompressed PublicKeyEncoding = iota
	PublicKeyEncodingUncompressed
)

func (encoding PublicKeyEncoding) Check() bool {
	return encoding == PublicKeyEncodingCompressed || encoding == PublicKeyEncodingUncompressed
}
