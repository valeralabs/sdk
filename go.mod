module github.com/valeralabs/sdk

go 1.18

replace golang.org/x/mobile => github.com/linden/mobile v0.0.0-20220709224438-7b6c75ef7684

require (
	github.com/btcsuite/btcd v0.23.1
	github.com/btcsuite/btcd/btcec/v2 v2.2.0
	github.com/btcsuite/btcd/btcutil v1.1.1
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1
	github.com/linden/binstruct v1.3.3-0.20220705022851-05a2106f4fb3
	github.com/linden/bite v0.0.0-20220704231724-6e6bb999fee4
	github.com/tidwall/gjson v1.14.1
	github.com/tyler-smith/go-bip39 v1.1.0
)

require (
	github.com/linden/fetch v0.0.0-20220607012258-070a21681224
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
)

require (
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.1 // indirect
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f // indirect
	github.com/decred/dcrd/crypto/blake256 v1.0.0 // indirect
	github.com/ethereum/go-ethereum v1.10.20
	golang.org/x/crypto v0.0.0-20220513210258-46612604a0f9 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
)
