package main

import (
	"fmt"
	"time"
)

func main() {
	var i int
	go fmt.Println(i)
	i++
	go func() { fmt.Println(i) }()
	i++
	go func(x int) { fmt.Println(x) }(i)
	i++
	time.Sleep(time.Second)
}
