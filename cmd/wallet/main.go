package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"valera.co/vdk/address"
	"valera.co/vdk/wallet"
)

var HelpMessage = strings.Join([]string{
	"wallet (flags) [method] [arguments...]",
	"",
	"help\tdisplay this message",
	"list\tlist the stacks accounts attached to a private key",
	"",
}, "\r\n")

var HelpListMessage = strings.Join([]string{
	"wallet list",
	"",
	"--mnemonic\t[seed phrase]",
	"--total \t[total addresses]",
	"--password\t[seed phrase password] (note: always leave blank if using hiro wallet)",
	"",
}, "\r\n")

var mnemonic = flag.String("mnemonic", "", "")
var password = flag.String("password", "", "")
var total = flag.Int("total", 0, "")

func init() {
	flag.Usage = func() {
		fmt.Println(HelpMessage)
	}

	flag.Parse()
}

func main() {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Printf("uh-oh, error occured \"%v\" please file an issue at valera.co/vdk\n", err)
		}
	}()

	switch flag.Arg(0) {
	case "list":
		if *mnemonic == "" {
			fmt.Println("mnemonic is required" + "\n\n" + HelpListMessage)
			os.Exit(0)
		}

		if *total == 0 {
			*total = 10
		}

		seed, err := wallet.NewSeed(*mnemonic, *password)

		if err != nil {
			panic(err)
		}

		manager, err := wallet.NewWallet(seed)

		if err != nil {
			panic(err)
		}

		version := address.AddressVersion{
			HashMode: address.HashModeP2PKH,
			Network:  address.NetworkMainnet,
		}

		for index := 0; index < *total; index++ {
			account, err := manager.Account(index)

			if err != nil {
				panic(err)
			}

			principal, err := address.NewAddressSingleSignature(account.PrivateKey.PublicKey(), version)

			if err != nil {
				panic(err)
			}

			c32, err := principal.C32()

			if err != nil {
				panic(err)
			}

			b58, err := principal.B58()

			if err != nil {
				panic(err)
			}

			fmt.Printf("c32 %s\t-   b58 %s\n", c32, b58)
		}

	default:
		fmt.Println(HelpMessage)
	}
}

// go run main.go --mnemonic="famous power eternal inquiry garbage sample news burger east little attitude genuine trumpet tube aunt fold run famous mercy mesh coral garlic razor aerobic" --password="password" --total 5 list
