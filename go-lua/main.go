package main

import (
	"github.com/Shopify/go-lua"
	"github.com/dimiro1/experiments/go-lua/helpers"
)

func main() {
	l := lua.NewState()
	lua.OpenLibraries(l)
	helpers.Open(l)

	if err := lua.DoFile(l, "config.lua"); err != nil {
		panic(err)
	}
}
