package main

import (
	"fmt"
	"strconv"
)

func binaryArray(n int) []string {
	ans := make([]string, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = strconv.FormatInt(int64(i), 2)
	}
	return ans
}

func main() {
	fmt.Println(binaryArray(2)) // [0 1 10]
	fmt.Println(binaryArray(3)) // [0 1 10 11]
}
