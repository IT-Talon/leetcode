package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 7))
}

// 62. 不同路径
func uniquePaths(m int, n int) int {
	// dp[m-1][n-1] = dp[m-2][n-1]+dp[m-1][n-2]  dp[0][n]=1  dp[m][0]=1
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j > 0 && i > 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
