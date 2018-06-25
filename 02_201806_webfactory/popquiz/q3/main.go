package main

func ch(ch chan int) {
	ch <- -<-ch
}

func main() {}
