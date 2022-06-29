package lengthprefixed

import (
	"bytes"
	"reflect"
	"testing"
)

func TestString(test *testing.T) {
	var err error
	var plain []byte

	marshaled := NewString([]byte("Hello World"))

	test.Run("marshal string", func(test *testing.T) {
		plain, err = marshaled.Marshal()

		if err != nil {
			test.Fatal(err)
		}

		if bytes.Compare(plain, []byte{0, 0, 0, 11, 72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}) != 0 {
			test.Fatalf("failed to marshal string got %+v", plain)
		}

		test.Logf("marshaled %x", plain)
	})

	test.Run("unmarshal string", func(test *testing.T) {
		var unmarshaled String

		err = unmarshaled.Unmarshal(plain)

		if err != nil {
			test.Fatal(err)
		}

		if reflect.DeepEqual(marshaled, unmarshaled) == false {
			test.Fatalf("failed to unmarshal string got %+v", unmarshaled)
		}

		test.Logf("unmarshaled %#+v\n", unmarshaled)
	})
}
