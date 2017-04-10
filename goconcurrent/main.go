package main

import "fmt"

func main() {

	stop := make(chan int)

	go func() {
		fmt.Println("I am in a goroutine")
		stop <- 0
	}()

	stopper := <-stop
	fmt.Println("Bye-bye: ", stopper)
}
