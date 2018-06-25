package main

import "fmt"

func main() {
	src := []int{2, 2, 2}
	dst := []int{}
	fmt.Println(copy(dst, src))
}
