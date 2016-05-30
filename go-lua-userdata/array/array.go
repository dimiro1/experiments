// https://www.lua.org/pil/28.1.html

package array

import (
	"fmt"

	"github.com/Shopify/go-lua"
)

var arraylibFunctions = []lua.RegistryFunction{
	{"new", newarray},
}

var arraylibMethods = []lua.RegistryFunction{
	{"__tostring", array2string},
	{"__len", getsize},
	{"__newindex", setarray},
	{"__index", getarray},
}

// Open open the array lib
func Open(l *lua.State) {
	lua.NewMetaTable(l, "go-lua.array")
	lua.SetFunctions(l, arraylibMethods, 0)

	array := func(l *lua.State) int {
		lua.NewLibrary(l, arraylibFunctions)
		return 1
	}

	lua.Require(l, "array", array, false)
	l.Pop(1)
}

func checkarray(l *lua.State) []float64 {
	arr := lua.CheckUserData(l, 1, "go-lua.array")
	lua.ArgumentCheck(l, arr != nil, 1, "'array' expected")
	return arr.([]float64)
}

func array2string(l *lua.State) int {
	arr := checkarray(l)
	l.PushString(fmt.Sprintf("%v", arr))
	return 1
}

func newarray(l *lua.State) int {
	size := lua.CheckInteger(l, 1)
	arr := make([]float64, size, size)

	l.PushUserData(arr)
	lua.SetMetaTableNamed(l, "go-lua.array")

	return 1
}

func getarray(l *lua.State) int {
	arr := checkarray(l)
	index := lua.CheckInteger(l, 2)

	l.PushNumber(arr[index-1])

	return 1
}

func setarray(l *lua.State) int {
	arr := checkarray(l)
	index := lua.CheckInteger(l, 2)
	value := lua.CheckNumber(l, 3)

	lua.ArgumentCheck(l, arr != nil, 1, "array expected")
	lua.ArgumentCheck(l, 1 <= index && index <= cap(arr), 2, "index out of range")

	arr[index-1] = value

	return 0
}

func getsize(l *lua.State) int {
	arr := checkarray(l)
	l.PushInteger(len(arr))

	return 1
}
