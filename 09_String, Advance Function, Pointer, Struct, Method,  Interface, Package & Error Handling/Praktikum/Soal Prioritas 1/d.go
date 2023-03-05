package main

import "fmt"

func findMaxMin(numbers []int) (int, int) {
	max := numbers[0]
	min := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return max, min
}

func main() {
	var nums [6]int
	for i := 0; i < 6; i++ {
		fmt.Printf("Enter number #%d: ", i+1)
		fmt.Scan(&nums[i])
	}
	max, min := findMaxMin(nums[:])
	fmt.Printf("%d is the maximum number\n", max)
	fmt.Printf("%d is the minimum number\n", min)
}
