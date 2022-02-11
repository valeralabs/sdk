package c32check

import (
	"fmt"
	"testing"
)

func TestEncode(test *testing.T) {
	raw, err := Encode(MainnetP2PKH, []byte("HelloWorld"))

	if err != nil {
		test.Fatalf("failed to encode: %v", err)
	}

	plain := fmt.Sprintf("%X", raw)

	if plain != "55D3FD4E102BF6994CBB8BFF0A47A07A5056C790A3D7AB24A42CA6799E939ADC" {
		test.Fatalf("failed (silently) to encode: got %v", err)
	}

	test.Log("got encoded", plain)
}
