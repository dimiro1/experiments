package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Option func(*User)

func WithName(name string) Option {
	return func(u *User) {
		u.Name = name
	}
}

func WithAge(age int) Option {
	return func(u *User) {
		u.Age = age
	}
}

func NewUser(options ...Option) User {
	u := User{}

	for _, option := range options {
		option(&u)
	}

	return u
}

func main() {
	u := NewUser(
		WithName("Claudemiro"),
		WithAge(29),
	)
	fmt.Printf("%#v\n", u)
}
