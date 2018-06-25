package main

func main() {
	type P *int
	type Q *int
	var p P = new(int)
	*p += 8
	var x *int = p
	var q Q = x
	*q++
	println(*p)
}
