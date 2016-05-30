package main

import (
	"fmt"
	"reflect"

	"github.com/Shopify/go-lua"
	"github.com/dimiro1/experiments/go-lua/helpers"
)

func main() {
	l := lua.NewState()
	lua.OpenLibraries(l)
	helpers.Open(l)

	/*
	 * offer = { id = "6bd07872-2730-4140-b5d3-250fb7f59d60" }
	 */
	l.NewTable()
	l.PushString("id")
	l.PushString("6bd07872-2730-4140-b5d3-250fb7f59d60")
	stackDump(l)
	l.SetTable(-3)
	l.SetGlobal("offer")
	stackDump(l)

	if err := lua.DoFile(l, "config.lua"); err != nil {
		panic(err)
	}

	// Calling function hello
	l.Global("hello")
	l.PushString("Claudemiro")
	stackDump(l)

	if err := l.ProtectedCall(1, 1, 0); err != nil {
		fmt.Println(err, reflect.TypeOf(err))
	}

	fmt.Println(l.ToBoolean(1))
	l.Pop(1)
	stackDump(l)
}

func stackDump(l *lua.State) {
	n := l.Top()

	fmt.Println("-------------------- Stack Dump -------------------")
	fmt.Printf("Total in stack: %d\n", n)

	for i := 1; i <= n; i++ {
		fmt.Printf("%d: ", i)

		switch l.TypeOf(i) {
		case lua.TypeString:
			s, _ := l.ToString(i)
			fmt.Printf("string %s", s)
		case lua.TypeBoolean:
			fmt.Printf("boolean %t", l.ToBoolean(i))
		case lua.TypeNumber:
			f, _ := l.ToNumber(i)
			fmt.Printf("number %f", f)
		default:
			fmt.Printf("%s", lua.TypeNameOf(l, i))
		}

		fmt.Println()
	}
	fmt.Println("--------------- Stack Dump Finished ---------------")
}
