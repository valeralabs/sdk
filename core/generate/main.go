package main

import (
	"encoding/base32"
	"fmt"
	"strconv"
	"strings"

	"github.com/vmihailenco/msgpack/v5"
)

//TODO (Linden): use generics
type Array struct {
	Length string
	Body   any
}

type Path string

type Function struct {
	Arguments  map[string]any
	ReturnType any
}

type Symbol struct {
	Path []string
	Body any
}

var _ msgpack.CustomDecoder = (*Function)(nil)

var VERSION = 1
var decoder = base32.StdEncoding.WithPadding(base32.NoPadding)

func parseType(raw any) any {
	switch value := raw.(type) {
	case string:
		return Path(value)
	case map[string]any:
		array, isArray := value["Array"]

		if isArray == true {
			nested := array.([]any)

			return Array{
				Length: nested[0].(string),
				Body:   nested[1],
			}
		}

		panic("TODO: handle non-array types")
	default:
		panic("file a bug report, you shouldn't be here")
	}
}

func (function *Function) DecodeMsgpack(unpacker *msgpack.Decoder) error {
	function.Arguments = make(map[string]any)

	raw, err := unpacker.DecodeUntypedMap()

	if err != nil {
		panic(err)
	}

	body, isFunction := raw["Function"]

	if isFunction == false {
		return fmt.Errorf("not a function, got %+v", raw)
	}

	base := body.([]any)
	arguments := base[0].([]any)

	for _, cursor := range arguments {
		argument := cursor.([]any)

		name := argument[0].(string)
		value := parseType(argument[1])

		function.Arguments[name] = value
	}

	function.ReturnType = parseType(base[1])

	return nil
}

func decode(plain string) Symbol {
	var err error

	prefix := fmt.Sprintf("MIKE_%d_", VERSION)

	if strings.HasPrefix(plain, prefix) == false {
		panic("symbol lacks the mike prefix")
	}

	plain = plain[len(prefix):]

	var total int

	for index := 0; index < len(plain); index++ {
		if plain[index] == '_' {
			total, err = strconv.Atoi(plain[:index])

			if err != nil {
				panic(err)
			}

			plain = plain[index:]

			break
		}
	}

	plain = plain[1:]

	var path []string
	var cursor string

	for index := 0; index < total; index++ {
		plain, cursor = component(plain)

		path = append(path, cursor)
	}

	plain = plain[1:]

	decoded, err := decoder.DecodeString(plain)

	if err != nil {
		panic(err)
	}

	var function Function

	err = msgpack.Unmarshal(decoded, &function)

	if err != nil {
		panic(err)
	}

	return Symbol{
		Path: path,
		Body: function,
	}
}

func component(plain string) (string, string) {
	var length int
	var err error

	for index := 0; index < len(plain); index++ {
		if 'a' <= plain[index] && plain[index] <= 'z' || 'A' <= plain[index] && plain[index] <= 'Z' || plain[index] == '_' {
			length, err = strconv.Atoi(plain[:index])

			if err != nil {
				panic(err)
			}

			plain = plain[index:]

			break
		}
	}

	return plain[length:], plain[:length]
}

func main() {
	fmt.Printf("example: %+v\n", decode("MIKE_1_4_4mike12capture_test7example4demo_QGUEM5LOMN2GS33OSKIZFJ3FPBQW24DMMWA2KQLSOJQXTEVCGMZKIUDBORUKIUDBORUA"))
}
