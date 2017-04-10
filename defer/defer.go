package main

import "fmt"

func test(b bool) {
	defer fmt.Println("[defer]")
	fmt.Println("> start")
	if b {
		fmt.Println("  true")
	} else {
		fmt.Println("  false")
	}
	fmt.Println("< end")
}

func main() {
	test(true)
	test(false)
}
