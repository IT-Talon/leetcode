package main

import (
	"fmt"
)

func main() {
	res := uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}})
	fmt.Println(res)

}

// dp[i][j] = dp[i][j-1]+dp[i-1][j]  dp[0][j] = 1  dp[i][0] = 1
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == 1 {
		return 0
	}
	dp := make([][]int, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		dp[i] = make([]int, len(obstacleGrid[0]))
		for j := 0; j < len(obstacleGrid[0]); j++ {

			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				if i == 0 && j == 0 {
					dp[i][j] = 1
				} else if i == 0 {
					dp[i][j] = dp[i][j-1]
				} else if j == 0 {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i-1][j] + dp[i][j-1]
				}
			}
		}
	}

	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
