package clarity

import (
	"bytes"
	"reflect"
	"testing"
)

func TestValue(test *testing.T) {
	var err error
	var plain []byte

	marshaled := NewValue([]byte("Hello World"))

	test.Run("marshal value", func(test *testing.T) {
		plain, err = marshaled.Marshal(false)

		if err != nil {
			test.Fatal(err)
		}

		if bytes.Compare(plain, []byte{0, 0, 0, 11, 72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}) != 0 {
			test.Fatalf("failed to marshal value got %+v", plain)
		}

		test.Logf("marshaled %x", plain)
	})

	test.Run("unmarshal value", func(test *testing.T) {
		var unmarshaled Value

		err = unmarshaled.Unmarshal(plain, false)

		if err != nil {
			test.Fatal(err)
		}

		if reflect.DeepEqual(marshaled, unmarshaled) == false {
			test.Fatalf("failed to unmarshal value got %+v", unmarshaled)
		}

		test.Logf("unmarshaled %#+v\n", unmarshaled)
	})
}

func TestTypedValue(test *testing.T) {
	var err error
	var plain []byte

	marshaled := NewValue([]byte("Hello World"))
	marshaled.Type = ClarityTypeStringASCII

	test.Run("marshal typed value", func(test *testing.T) {
		plain, err = marshaled.Marshal(true)

		if err != nil {
			test.Fatal(err)
		}

		if bytes.Compare(plain, []byte{13, 0, 0, 0, 11, 72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}) != 0 {
			test.Fatalf("failed to marshal typed value got %+v", plain)
		}

		test.Logf("marshaled %x", plain)
	})

	test.Run("unmarshal value", func(test *testing.T) {
		var unmarshaled Value

		err = unmarshaled.Unmarshal(plain, true)

		if err != nil {
			test.Fatal(err)
		}

		if reflect.DeepEqual(marshaled, unmarshaled) == false {
			test.Fatalf("failed to unmarshal value got %+v", unmarshaled)
		}

		test.Logf("unmarshaled %#+v\n", unmarshaled)
	})
}

func TestList(test *testing.T) {
	var err error
	var plain []byte

	marshaled := NewList([][]byte{[]byte("Hello"), []byte("World")})

	test.Run("marshal list", func(test *testing.T) {
		plain, err = marshaled.Marshal(false)

		if err != nil {
			test.Fatal(err)
		}

		if bytes.Compare(plain, []byte{0, 0, 0, 2, 0, 0, 0, 5, 72, 101, 108, 108, 111, 0, 0, 0, 5, 87, 111, 114, 108, 100}) != 0 {
			test.Fatalf("failed to marshal list got %+v", plain)
		}

		test.Logf("marshaled %v", plain)
	})

	test.Run("unmarshal list", func(test *testing.T) {
		var unmarshaled List

		err = unmarshaled.Unmarshal(plain, false)

		if err != nil {
			test.Fatal(err)
		}

		if reflect.DeepEqual(marshaled, unmarshaled) == false {
			test.Fatalf("failed to unmarshal list got %v", unmarshaled)
		}

		test.Logf("unmarshaled %#+v\n", unmarshaled)
	})
}
