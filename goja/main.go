package main

import (
	"fmt"

	"github.com/dop251/goja"
)

type person struct {
	Name string
	Age  int
}

func newPerson(name string, age int) person { return person{name, age} }

func main() {
	vm := goja.New()
	vm.Set("person", vm.ToValue(person{Name: "Claudemiro", Age: 28}))
	vm.Set("println", fmt.Println)
	vm.Set("printf", fmt.Printf)
	vm.Set("newPerson", newPerson)

	if _, err := vm.RunString(`println(newPerson("Claudemiro", 28))`); err != nil {
		panic(err)
	}

	if _, err := vm.RunString(`println(Object.keys(person))`); err != nil {
		panic(err)
	}
}
