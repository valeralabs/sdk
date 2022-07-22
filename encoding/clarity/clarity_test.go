package clarity

import (
	"bytes"
	"reflect"
	"testing"

	"valera.co/vdk/address"
)

func TestValue(test *testing.T) {
	var err error
	var plain []byte

	marshaled := NewValue([]byte("Hello World"), ClarityTypeStringUTF8)

	test.Run("marshal value", func(test *testing.T) {
		plain, err = marshaled.Marshal(false)

		if err != nil {
			test.Fatal(err)
		}

		if bytes.Compare(plain, []byte{0, 0, 0, 11, 72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}) != 0 {
			test.Fatalf("failed to marshal value got %+v", plain)
		}

		test.Logf("marshaled %v", plain)
	})

	test.Run("unmarshal value", func(test *testing.T) {
		unmarshaled := Value{
			Type: ClarityTypeStringUTF8,
		}

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

	marshaled := NewValue([]byte("Hello World"), ClarityTypeStringUTF8)

	test.Run("marshal typed value", func(test *testing.T) {
		plain, err = marshaled.Marshal(true)

		if err != nil {
			test.Fatal(err)
		}

		if bytes.Compare(plain, []byte{14, 0, 0, 0, 11, 72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}) != 0 {
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

	marshaled := NewList([][]byte{[]byte("Hello"), []byte("World")}, ClarityTypeStringUTF8)

	test.Logf("here %+v\n", marshaled)

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
		unmarshaled := List{
			Type: ClarityTypeStringUTF8,
		}

		err = unmarshaled.Unmarshal(plain, false)

		if err != nil {
			test.Fatal(err)
		}

		if reflect.DeepEqual(marshaled, unmarshaled) == false {
			test.Fatalf("failed to unmarshal list got %+v", unmarshaled)
		}

		test.Logf("unmarshaled %#+v\n", unmarshaled)
	})
}

func TestTypes(test *testing.T) {
	test.Run("marshal int (/uint)", func(test *testing.T) {
		digit := NewValue(100, ClarityTypeInt)

		marshaled, err := digit.Marshal(true)

		if err != nil {
			test.Logf("failed to marshal %v\n", err)
		}

		if bytes.Compare(marshaled, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 100}) != 0 {
			test.Fatalf("got invalid encoding %v\n", marshaled)
		}

		test.Logf("marshaled %v\n", marshaled)
	})

	test.Run("marshal bool & none", func(test *testing.T) {
		bool := NewValue(nil, ClarityTypeBoolTrue)

		marshaled, err := bool.Marshal(true)

		if err != nil {
			test.Logf("failed to marshal %v\n", err)
		}

		if bytes.Compare(marshaled, []byte{3}) != 0 {
			test.Logf("got invalid encoding %x\n", marshaled)
		}

		test.Logf("marshaled %v\n", marshaled)
	})

	test.Run("marshal principal", func(test *testing.T) {
		version := address.AddressVersion{
			HashMode: address.HashModeP2PKH,
			Network:  address.NetworkMainnet,
		}

		value := address.Address{
			Version:  version,
			Hash:     []byte{21, 195, 27, 140, 28, 17, 197, 21, 226, 68, 183, 88, 6, 186, 196, 141, 19, 153, 199, 117},
			Contract: "",
		}

		principal := NewValue(value, ClarityTypePrincipalStandard)

		marshaled, err := principal.Marshal(true)

		if err != nil {
			test.Logf("failed to marshal %v\n", err)
		}

		if bytes.Compare(marshaled, []byte{5, 22, 21, 195, 27, 140, 28, 17, 197, 21, 226, 68, 183, 88, 6, 186, 196, 141, 19, 153, 199, 117}) != 0 {
			test.Logf("got invalid encoding %x\n", marshaled)
		}

		test.Logf("marshaled %v\n", marshaled)
	})
}
