package main

import (
	"encoding/base32"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"

	jenny "github.com/valeralabs/jenny/jen"
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

func (path Path) ConvertRust(argument *jenny.Statement) (*jenny.Statement, bool) {
	value := argument

	switch path {
	case "u8":
		value.ID("uint")

	case "u16":
		value.ID("uint")

	case "u32":
		value.ID("uint")

	case "usize":
		value.ID("int")

	case "i8":
		value.ID("int")

	case "i16":
		value.ID("int")

	case "i32":
		value.ID("int")

	case "isize":
		value.ID("int")

	default:
		return value.Qualified("unsafe", "Pointer"), false
	}

	return value, true
}

func (path Path) ConvertBlock(name string) *jenny.Statement {
	value := jenny.ID("_" + name).Op(":=")

	switch path {
	case "u8":
		value.Qualified("github.com/linden/ffi", "CreateU8")

	case "u16":
		value.Qualified("github.com/linden/ffi", "CreateU16")

	case "u32":
		value.Qualified("github.com/linden/ffi", "CreateU32")

	case "usize":
		value.Qualified("github.com/linden/ffi", "CreateUSize")

	case "i8":
		value.Qualified("github.com/linden/ffi", "CreateI8")

	case "i16":
		value.Qualified("github.com/linden/ffi", "CreateI16")

	case "i32":
		value.Qualified("github.com/linden/ffi", "CreateI32")

	case "isize":
		value.Qualified("github.com/linden/ffi", "CreateISize")

	default:
		panic("can't handle unknown types")
	}

	value.Call(jenny.ID(name))

	return value
}

func (path Path) ConvertGo() *jenny.Statement {
	var value *jenny.Statement

	switch path {
	case "u8":
		value = jenny.Qualified("github.com/linden/ffi", "CreateUIntFromU8")

	case "u16":
		value = jenny.Qualified("github.com/linden/ffi", "CreateUIntFromU16")

	case "u32":
		value = jenny.Qualified("github.com/linden/ffi", "CreateUIntFromU32")

	case "usize":
		value = jenny.Qualified("github.com/linden/ffi", "CreateUIntFromUSize")

	case "i8":
		value = jenny.Qualified("github.com/linden/ffi", "CreateIntFromI8")

	case "i16":
		value = jenny.Qualified("github.com/linden/ffi", "CreateIntFromI16")

	case "i32":
		value = jenny.Qualified("github.com/linden/ffi", "CreateIntFromI32")

	case "isize":
		value = jenny.Qualified("github.com/linden/ffi", "CreateIntFromISize")

	default:
		panic("can't handle unknown types")
	}

	value.Call(jenny.ID("_return"))

	return jenny.Return(value)
}

func unbox(path Path) Path {
	cursor := string(path)

	if len(cursor) > len("Box<>") {
		cursor = cursor[len("Box<"):]
		cursor = cursor[:(len(cursor) - len(">"))]
	}

	return Path(cursor)
}

func main() {
	command := exec.Command("nm", "-j", "../libblockstack_lib.dylib")
	output, err := command.Output()

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(output), "\n")

	structures := map[string]*jenny.Statement{}
	functions := []*jenny.Statement{}

	var header string

	file := jenny.NewFile("core")

	file.PackageComment("DO NOT EDIT. Code generated by github.com/valeralabs/core/generate. //")

	file.CgoPreamble(strings.Join([]string{
		"#cgo LDFLAGS: -L. -lblockstack_lib",
		"#include \"header.h\"",
	}, "\n"))

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

		wrapper := jenny.Func()

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

			if _, ok := structures[structure]; !ok {
				structures[structure] = jenny.Type().ID(structure).Structure(
					jenny.ID("pointer").Qualified("unsafe", "Pointer"),
				)
			}

			if hasSelf {
				wrapper.Parameters(
					jenny.ID("self").Op("*").ID(structure),
				)
			} else {
				wrapper.Parameters(
					jenny.ID("_").Op("*").ID(structure),
				)
			}
		}

		wrapper.ID(CamelCase(name, true))

		parameters := []jenny.Code{}

		if hasSelf {
			parameters = append(parameters, jenny.ID("self.pointer"))
		}

		arguments := []jenny.Code{}
		blocks := []jenny.Code{}

		for _name, _path := range function.Arguments {
			name := CamelCase(_name, false)
			path := _path.(Path)

			//TODO (Linden): fix mike `Result` handling
			if path == Path("Box") {
				continue
			}

			path = unbox(path)

			argument, isKnown := path.ConvertRust(jenny.ID(name))

			arguments = append(arguments, argument)

			parameter := jenny.ID(name)

			if isKnown {
				parameter = jenny.ID("_" + name)

				blocks = append(blocks, path.ConvertBlock(name))
			}

			parameters = append(parameters, parameter)
		}

		returnPath := function.ReturnType.(Path)
		returnPath = unbox(returnPath)

		returnValue, isKnown := returnPath.ConvertRust(jenny.ID(""))

		if isKnown {
			blocks = append(blocks, jenny.ID("_return").Op(":=").Qualified("C", decoded.Unmangled).Call(parameters...))
			blocks = append(blocks, returnPath.ConvertGo())
		} else {
			blocks = append(blocks, jenny.Return(jenny.Qualified("C", decoded.Unmangled).Call(parameters...)))
		}

		wrapper.Parameters(arguments...)
		wrapper.Parameters(returnValue)
		wrapper.Block(blocks...)

		functions = append(functions, wrapper)
	}

	for _, structure := range structures {
		file.Add(structure)
	}

	for _, function := range functions {
		file.Add(function)
	}

	err = ioutil.WriteFile("../header.h", []byte(header), 0644)

	if err != nil {
		panic(err)
	}

	err = file.Save("../core.gen.go")

	if err != nil {
		panic(err)
	}
}
