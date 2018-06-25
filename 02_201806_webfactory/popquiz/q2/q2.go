package main

func main() {
	print(len(map[interface{}]int{
		new(int):      1,
		new(int):      2,
		new(struct{}): 3,
		new(struct{}): 4,
	}))
}
