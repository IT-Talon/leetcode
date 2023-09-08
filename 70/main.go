package main

import "fmt"

// 70. 爬楼梯
func main() {
	fmt.Println(climbStairs(4))
}

// 1   1
// 2   2
// 3   3
// 4   5
// n   (n-1)+(n-2)
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b, c := 1, 2, 0
	for i := 2; i < n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}
