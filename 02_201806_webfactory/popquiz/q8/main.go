package main

import "fmt"

type A int

func (a A) A() string { return "a" }

type B struct{ A }

func main() {
	var b B
	b.A = 1
	fmt.Printf("%T", b.A())
}
