package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	ch <- 100
	value := <-ch
	fmt.Println(ch)
	close(ch)
	fmt.Println(value)
}
