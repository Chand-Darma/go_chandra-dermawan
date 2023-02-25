package main

import "fmt"

func PairSum(arr []int, target int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println("Explanation: The numbers at index 1 and 3 add up to 6: 2+4=6")
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11)) // [0, 2]
	fmt.Println("Explanation: The numbers at index 0 and 2 add up to 11: 2+9=11")
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12)) // [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10)) // [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))  // [0, 1]
}
