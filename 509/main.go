package main

import "fmt"

// 509. 斐波那契数
func main() {
	fmt.Println(fib(4))
}

// 0 1 1 2 3 5 8
func fib(n int) int {
	if n <= 1 {
		return n
	}
	a, b, c := 0, 1, 0
	for i := 2; i <= n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}
