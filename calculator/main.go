// See: https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions
package main

import (
	"fmt"
	"math"
)

type Operation func(float64) float64

type Calculator struct {
	acc float64
}

func (c Calculator) Display() {
	fmt.Println(c.acc)
}

func (c *Calculator) Do(ops ...Operation) float64 {
	for _, op := range ops {
		c.acc = op(c.acc)
	}

	return c.acc
}

func Sum(n float64) Operation {
	return func(acc float64) float64 {
		return acc + n
	}
}

func Mul(n float64) Operation {
	return func(acc float64) float64 {
		return acc * n
	}
}

func main() {
	c := Calculator{}
	c.Do(Sum(10), Sum(5))
	c.Do(Sum(20))
	c.Do(Mul(2))
	c.Do(math.Sqrt)

	c.Display()
}
