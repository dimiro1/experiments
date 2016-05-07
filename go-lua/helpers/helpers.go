package helpers

import (
	"log"

	"github.com/Shopify/go-lua"
)

// Open register a new helpers library
func Open(l *lua.State) {
	lua.Require(l, "helpers", func(l *lua.State) int {
		lua.NewLibrary(l, registerLibrary)
		return 1
	}, false)

	l.Pop(1)
}

var registerLibrary = []lua.RegistryFunction{
	{"registerService", registerService},
}

func registerService(l *lua.State) int {
	path := lua.CheckString(l, 1)

	log.Println(path)

	return 1
}
