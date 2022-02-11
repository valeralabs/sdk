// Copyright 2013 Richard Lehane. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package crock32 implements Douglas Crockford's Base32 encoding.
//
// Crock32 is useful for "expressing numbers in a form that can be conveniently and accurately transmitted between humans and computer systems."
// See http://www.crockford.com/wrmg/base32.html for details.
// Note: crock32 differs from Crockford in its use of lower-case letters when encoding (decode works for both cases). To change, use: crock32.SetDigits("0123456789ABCDEFGHJKMNPQRSTVWXYZ")
//
// Example:
//   i, _ := crock32.Decode("a1j3")
//   s := crock32.Encode(i)
//   fmt.Println(s)
package crock32

import "errors"

const cutoff uint64 = (1<<64-1)/32 + 1
const maxuint = 13

var digits = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

// Decode converts a string matching Douglas Crockford's character set (case insensitive) into an unsigned 64-bit integer.
func Decode(from []byte) (uint64, error) {
	var decoded uint64

	for _, char := range from {
		var to byte

		switch {
		case char == 'O', char == 'o':
			to = '0'
		case char == 'L', char == 'l', char == 'I', char == 'i':
			to = '1'
		case '0' <= char && char <= '9':
			to = char - '0'
		case 'a' <= char && char <= 'h':
			to = char - 'a' + 10
		case 'A' <= char && char <= 'H':
			to = char - 'A' + 10
		case 'j' <= char && char <= 'k':
			to = char - 'a' + 9
		case 'J' <= char && char <= 'K':
			to = char - 'A' + 9
		case 'm' <= char && char <= 'n':
			to = char - 'a' + 8
		case 'M' <= char && char <= 'N':
			to = char - 'A' + 8
		case 'p' <= char && char <= 't':
			to = char - 'a' + 7
		case 'P' <= char && char <= 'T':
			to = char - 'A' + 7
		case 'v' <= char && char <= 'z':
			to = char - 'a' + 6
		case 'V' <= char && char <= 'Z':
			to = char - 'A' + 6
		default:
			return 0, errors.New("crock32.Decode: invalid character " + string(char))
		}

		if decoded >= cutoff {
			return 0, errors.New("crock32.Decode: overflowing uint64")
		}

		decoded = decoded*32 + uint64(to)
	}

	return decoded, nil
}

// Encode converts a uint64 into a Crockford base32 encoded string
func Encode(n uint64) string {
	var a [maxuint]byte

	i := maxuint

	for n >= 32 {
		i--

		a[i] = digits[n%32]

		n /= 32
	}
	i--

	a[i] = digits[n]

	return string(a[i:])
}
