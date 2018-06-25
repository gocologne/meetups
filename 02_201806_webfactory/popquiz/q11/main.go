package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	var i int
	for i = range x {
		x = nil
	}
	fmt.Println(i)
}
