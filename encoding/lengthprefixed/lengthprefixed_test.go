package lengthprefixed

import (
	"testing"
)

func TestStringMarshal(test *testing.T) {
	LPString := CreateString([]byte("Hello World"))

	plain, err := LPString.Marshal()

	if err != nil {
		test.Fatal(err)
	}

	if string(plain) != "11721011081081113287111114108100" {
		test.Fatalf("failed to encode LP string")
	}

	test.Log("encoded", string(plain))
}
