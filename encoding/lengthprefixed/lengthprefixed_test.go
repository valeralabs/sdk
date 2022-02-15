package lengthprefixed

import (
	"bytes"
	"testing"
)

func TestStringMarshal(test *testing.T) {
	LPString := CreateString([]byte("Hello World"))

	plain, err := LPString.Marshal()

	if err != nil {
		test.Fatal(err)
	}

	if bytes.Compare(plain, []byte{
		48, 48, 48, 48, 48, 48, 48, 49, 49, 55, 50, 49, 48, 49,
		49, 48, 56, 49, 48, 56, 49, 49, 49, 51, 50, 56, 55, 49,
		49, 49, 49, 49, 52, 49, 48, 56, 49, 48, 48,
	}) != 0 {
		test.Fatalf("failed to encode LP string")
	}

	test.Log("encoded", string(plain))
}
