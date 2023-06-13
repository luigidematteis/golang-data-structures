package arrays

import (
	"fmt"
)

/*
	FindLength find the length of the longest subarray with a sum less than or equal to k
*/
func FindLength(nums []int, k int) int {
	fmt.Println()
	fmt.Println("Hello from Sliding Window")

	// 3, 1, 2, 7, 4, 2, 1, 1, 5
	var left = 0
	var currentValue = 0
	var result = 0

	for i, num := range nums {
		currentValue += num

		for currentValue > k {
			currentValue -= nums[left]
			left++
		}

		result = max(result, i-left+1)

	}

	fmt.Println("Result:", result)
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
