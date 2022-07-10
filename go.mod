module github.com/valeralabs/sdk

go 1.18

replace golang.org/x/mobile => github.com/linden/mobile v0.0.0-20220709224438-7b6c75ef7684

require (
	github.com/btcsuite/btcd v0.23.1
	github.com/btcsuite/btcd/btcec/v2 v2.1.3
	github.com/btcsuite/btcd/btcutil v1.1.1
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1
	github.com/linden/binstruct v1.3.3-0.20220705022851-05a2106f4fb3
	github.com/linden/bite v0.0.0-20220704231724-6e6bb999fee4
	github.com/tyler-smith/go-bip39 v1.1.0
)

require (
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.1 // indirect
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f // indirect
	github.com/decred/dcrd/crypto/blake256 v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/mobile v0.0.0-20220518205345-8578da9835fd // indirect
	golang.org/x/mod v0.4.2 // indirect
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	golang.org/x/tools v0.1.8-0.20211022200916-316ba0b74098 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)
