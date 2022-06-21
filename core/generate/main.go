package main

import (
	"encoding/base32"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"

	"github.com/vmihailenco/msgpack/v5"
)

//TODO (Linden): use generics
type Array struct {
	Length int
	Body   any
}

type Path string

type Pointer any

type Function struct {
	Arguments  map[string]any
	ReturnType any
}

type Symbol struct {
	Unmangled string
	Path      []string
	Body      any
}

var _ msgpack.CustomDecoder = (*Function)(nil)

var path string

var VERSION = 1
var decoder = base32.StdEncoding.WithPadding(base32.NoPadding)
var mapping = map[string]string{
	"id":       "ID",
	"addrtype": "addressType",
	"b58":      "B58",
	"pkg":      "package",
	"pubkey":   "PublicKey",
	"p2pkh":    "P2PKH",
	"p2sh":     "P2SH",
	"tx":       "TX",
}

func parseType(raw any) any {
	switch value := raw.(type) {
	case map[string]any:
		array, isArray := value["Array"]

		if isArray == true {
			nested := array.([]any)

			return Array{
				Length: int(nested[0].(int8)),
				Body:   nested[1],
			}
		}

		path, isPath := value["Path"]

		if isPath == true {
			nested := path.([]any)

			return Path(nested[0].(string))
		}

		pointer, isPointer := value["Pointer"]

		if isPointer == true {
			nested := pointer.([]any)

			return Pointer(nested[0])
		}

		panic("TODO: handle non-(array/path) types")
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

func decode(plain string) (Symbol, error) {
	var err error

	prefix := fmt.Sprintf("_MIKE_%d_", VERSION)

	if strings.HasPrefix(plain, prefix) == false {
		return Symbol{}, fmt.Errorf("symbol lacks the mike prefix")
	}

	unmangled := plain[1:]
	plain = plain[len(prefix):]

	var total int

	for index := 0; index < len(plain); index++ {
		if plain[index] == '_' {
			total, err = strconv.Atoi(plain[:index])

			if err != nil {
				return Symbol{}, err
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

	last := len(path) - 1

	path[last] = strings.TrimPrefix(path[last], "mike_fn_")
	path[last] = strings.TrimPrefix(path[last], "mike_new_")

	plain = plain[1:]

	decoded, err := decoder.DecodeString(plain)

	if err != nil {
		return Symbol{}, err
	}

	var function Function

	err = msgpack.Unmarshal(decoded, &function)

	if err != nil {
		return Symbol{}, err
	}

	return Symbol{
		Unmangled: unmangled,
		Path:      path,
		Body:      function,
	}, nil
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

func CamelCase(name string, function bool) string {
	var result string

	for index, cursor := range strings.Split(name, "_") {
		if function == false && index == 0 {
			result += cursor

			continue
		}

		for before, after := range mapping {
			if cursor == before {
				cursor = after
			}
		}

		cursor = strings.Title(cursor)

		result += cursor
	}

	return result
}

func main() {
	command := exec.Command("nm", "-j", "../libblockstack_lib.dylib")
	output, err := command.Output()

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(output), "\n")

	created := map[string]bool{}

	var header string

	module := "package core\n\n"
	module += "//#cgo LDFLAGS: -L. -lblockstack_lib\n"
	module += "//#include \"header.h\"\n"
	module += "import \"C\"\n"
	module += "import \"unsafe\"\n\n"

	for _, line := range lines {
		decoded, err := decode(line)

		if err != nil || decoded.Path[0] != "blockstack_lib" {
			continue
		}

		function := decoded.Body.(Function)

		header += "void"

		if function.ReturnType != "Nothing" {
			header += "*"
		}

		header += " "
		header += decoded.Unmangled
		header += "("

		header += strings.Repeat("void*, ", len(function.Arguments))

		if header[len(header)-2] == ',' {
			header = header[:len(header)-2]
		}

		header += ");\n"

		name := decoded.Path[len(decoded.Path)-1]

		hasSelf := false

		if _, ok := function.Arguments["wrapped_self"]; ok {
			delete(function.Arguments, "wrapped_self")

			hasSelf = true
		}

		if strings.Contains(name, "__") == true {
			split := strings.Split(name, "__")

			structure := split[0]
			name = split[1]

			if _, ok := created[structure]; !ok {
				module += "type " + structure + " unsafe.Pointer\n\n"
				created[structure] = true
			}

			fmt.Println(function.Arguments)

			if hasSelf {
				module += "func (self *"
			} else {
				module += "func (_ "
			}

			module += structure + ") " + CamelCase(name, true)
		} else {
			module += "func " + CamelCase(name, true)
		}

		module += "("

		var arguments []string

		for name, _ := range function.Arguments {
			arguments = append(arguments, CamelCase(name, false)+" unsafe.Pointer")
		}

		module += strings.Join(arguments, ", ")
		module += ") {\n"
		module += "\tC." + decoded.Unmangled + "("

		arguments = []string{}

		if hasSelf {
			arguments = append(arguments, "self")
		}

		for name, _ := range function.Arguments {
			arguments = append(arguments, CamelCase(name, false))
		}

		module += strings.Join(arguments, ", ")
		module += ")\n"
		module += "}\n\n"
	}

	err = ioutil.WriteFile("../header.h", []byte(header), 0644)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("../core.gen.go", []byte(module), 0644)

	if err != nil {
		panic(err)
	}
}
