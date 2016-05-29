package main

import (
	"github.com/Shopify/go-lua"
	"github.com/dimiro1/experiments/go-lua/helpers"
)

func main() {
	l := lua.NewState()
	lua.OpenLibraries(l)
	helpers.Open(l)

	l.NewTable()
	l.PushString("id")
	l.PushString("6bd07872-2730-4140-b5d3-250fb7f59d60")
	l.SetTable(-3)
	l.SetGlobal("offer")

	if err := lua.DoFile(l, "config.lua"); err != nil {
		panic(err)
	}
}
