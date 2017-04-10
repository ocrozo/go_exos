package main

import "fmt"

func main() {

	c := make(chan int, 16)
	for i := 0; i < 8; i++ {
		go func() {
			c <- i
			c <- i * 10
		}()
	}

	for i := 0; i < 17; i++ {
		fmt.Println(<-c)
	}
}
