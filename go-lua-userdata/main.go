package main

import (
	"github.com/Shopify/go-lua"
	"github.com/dimiro1/experiments/go-lua-userdata/array"
)

func main() {
	l := lua.NewState()
	lua.OpenLibraries(l)
	array.Open(l)

	if err := lua.DoFile(l, "script.lua"); err != nil {
		panic(err)
	}
}
