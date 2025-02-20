package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err == nil {
			fmt.Println("expected panic")
		}
	}()
	ch := make(chan int, 10)
	close(ch)
	ch <- 1

}
