package c32

import (
	"testing"
)

func TestEncode(test *testing.T) {
	encoded, err := Encode([]byte("abcdef"))

	if err != nil {
		test.Fatalf("failed to encode got %v", err)
	}

	if string(encoded) != "AQKFF" {
		test.Fatalf("expected \"AQKFF\" got \"%s\"", encoded)
	}

	test.Log("got encoded", string(encoded))
}

func TestChecksum(test *testing.T) {
	checksum := Checksum([]byte("0123456789abcdef"))

	if string(checksum) != "137ad663" {
		test.Fatalf("expected \"137ad663\" got \"%s\"", checksum)
	}

	test.Log("got checksum", string(checksum))
}

func TestCheckEncode(test *testing.T) {
	checksum, err := ChecksumEncode([]byte("a956ed79819901b1b2c7b3ec045081f749c588ed"), 22)

	if err != nil {
		test.Fatalf("failed to encode got %v", err)
	}

	if string(checksum) != "P2MNDVBSG6CG3CDJRYSYR12GG7VMKHC8XMVVHEH6" {
		test.Fatalf("expected \"P2MNDVBSG6CG3CDJRYSYR12GG7VMKHC8XMVVHEH6\" got \"%s\"", checksum)
	}

	test.Log("got checksum", string(checksum))
}
