package main

import (
	"fmt"
	"start/mymod"
)

func main() {

	var three int = mymod.Sum(1, 2)
	fmt.Println("1 + 2 = ", three)

	var trois = mymod.AlsoSum(1, 2)
	fmt.Println("1 + 2 = ", trois)

	one := mymod.Identity(1)
	fmt.Println("1 == ", one)

	a, b := mymod.Swap("a", "b")
	fmt.Println("a <-> b = ", a, b)

	a, b = mymod.AlsoSwap("a", "b")
	fmt.Println("a <-> b = ", a, b)

	adder := mymod.Adder(10)
	fmt.Println("5 +=10 = ", adder(5))
}
