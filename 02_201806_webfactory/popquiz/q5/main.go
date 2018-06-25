package main

func main() {
again:
	for i := 0; i < 10; i++ {
		if i > 6 {
			break again
		}
	}
}
