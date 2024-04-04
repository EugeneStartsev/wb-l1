package main

import "fmt"

type target interface {
	operation()
}

type adapter struct {
	a, b int
}

type concreteAdapter struct {
	*adapter
}

func (c *concreteAdapter) operation() {
	c.multiplication()
}

func newAdapter(ad *adapter) target {
	return &concreteAdapter{ad}
}

func (ad *adapter) multiplication() {
	fmt.Println(ad.a * ad.b)
}

func main() {
	ad := newAdapter(&adapter{10, 20})
	ad.operation()
}
