package main

import "fmt"

func main() {
	arraysAndSlices()
	fmt.Println("\n-----------\n")
	maps()
	fmt.Println("\n-----------\n")
	person()
}
func dump(name string, slice []int) {
	fmt.Println(name, " = ", slice, " len = ", len(slice), " cap = ", cap(slice))
}

func arraysAndSlices() {

	var a1 []int = []int{0, 1, 2, 3, 4}
	dump("a1", a1)

	a2 := []int{10, 11, 12, 13, 14}
	dump("a2", a2)

	a3 := append(a1, 5)
	dump("a1", a1)
	dump("a3", a3)

	a4 := make([]int, 5, 10)
	dump("a4", a4)
	copy(a4, a2)
	dump("a4", a4)

	dump("a4[1:3]", a4[1:3])
	dump("a4[1:]", a4[1:])
	dump("a4[:3]", a4[:3])

	a5 := make([]int, 20)
	dump("a5", a5)
}
func mapDump(name string, data map[string]int) {
	fmt.Println(name, " = ", data, " len = ", len(data))
}

func maps() {

	ints := make(map[string]int)
	mapDump("ints", ints)

	ints["1"] = 1
	ints["2"] = 2
	ints["3"] = 3
	ints["4"] = 4
	mapDump("ints", ints)

	delete(ints, "4")
	mapDump("ints", ints)

	fmt.Println("ints[\"666\"] = ", ints["666"])

	res, ok := ints["666"]
	fmt.Println("res = ", res, ", ok = ", ok)

	res, ok = ints["1"]
	fmt.Println("res = ", res, ", ok = ", ok)

	_, ok = ints["999"]
	fmt.Println(ok)
}

type Person struct {
	name  string
	email string
	age   uint8
}

func person() {
	mrbean := Person{"Mr Bean", "bean@outlook.com", 59}
	fmt.Println(mrbean)
	fmt.Println(mrbean.email)
}
