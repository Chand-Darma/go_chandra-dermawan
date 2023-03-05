package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	n := len(jumps)
	dp := make([]int, n)
	dp[0] = 0

	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt32
		for j := 0; j < i; j++ {
			if j+jumps[j] >= i {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[j]+1)))
			}
		}
	}

	return dp[n-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 30
}
