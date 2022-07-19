package transaction

import (
	"fmt"
	"testing"
)

func TestPayloadSmartContract(test *testing.T) {
	payload := PayloadSmartContract{"hello", ";; hello world"}

	marshaled, err := payload.Marshal()

	if err != nil {
		test.Fatalf("failed to encode payload %v", err)
	}

	if fmt.Sprintf("%x", marshaled) != "010568656c6c6f0000000e3b3b2068656c6c6f20776f726c64" {
		test.Fatalf("failed to encode payload, invalid result")
	}

	test.Logf("payload %x\n", marshaled)
}
