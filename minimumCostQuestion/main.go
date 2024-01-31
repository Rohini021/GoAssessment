package main

import (
	"fmt"
)

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n <= 1 {
		return 0
	}

	dp := make([]int, n)

	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < n; i++ {
		dp[i] = cost[i] + min(dp[i-1], dp[i-2])
	}

	return min(dp[n-1], dp[n-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	cost := []int{10, 15, 20} // tried with {1, 2, 20, 5, 7} output was 7
	result := minCostClimbingStairs(cost)
	fmt.Println("Output:", result)
}
