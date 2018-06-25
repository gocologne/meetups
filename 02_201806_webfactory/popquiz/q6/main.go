package main

import (
	"fmt"
)

func fn(m map[int]int) {
	m = map[int]int{42: -6}
}

func main() {
	var m map[int]int
	fn(m)
	fmt.Println(len(m))
}
