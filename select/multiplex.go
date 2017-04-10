package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func FibFun() func() uint64 {
	a, b := uint64(0), uint64(1)
	return func() uint64 {
		a, b = b, a+b
		return a
	}
}

func main() {

	fib := FibFun()
	fibs := make(chan uint64, 10)
	computeFibs := make(chan int, 10)
	go func() {
		for {
			<-computeFibs
			fibs <- fib()
		}
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, os.Kill)

	running := true
	for running {
		select {
		case next := <-fibs:
			fmt.Println("> ", next)
		case <-time.Tick(1 * time.Second):
			fmt.Println("[tick]")
			for i := 0; i < 5; i++ {
				computeFibs <- 1
			}
		case <-exit:
			fmt.Println("~ stop")
			running = false
		}
	}

	fmt.Println("Bye!")
}
