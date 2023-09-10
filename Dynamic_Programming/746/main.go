package main

import "math"

// 746. 使用最小花费爬楼梯
func main() {

}

// dp[i] = min(dp[i-1]+cost[i-1],dp[i-2]+cost[i-2])
func minCostClimbingStairs(cost []int) int {
	a, b, c := 0, 0, 0
	for i := 2; i <= len(cost); i++ {
		c = int(math.Min(float64(a+cost[i-2]), float64(b+cost[i-1])))
		a, b = b, c
	}
	return c
}
