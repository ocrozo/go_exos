package mymod

func Identity(n int) int {
	return n
}

func Sum(x int, y int) int {
	return x + y
}

func AlsoSum(x, y int) int {
	return x + y
}

func Swap(x, y string) (string, string) {
	return y, x
}

func AlsoSwap(x, y string) (a, b string) {
	a = y
	b = x
	return
}

type intFunc func(int) int

func Adder(x int) intFunc {
	return func(y int) int {
		return x + y
	}
}
