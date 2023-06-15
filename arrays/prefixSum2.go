package arrays

import (
	"fmt"
	"math"
)

/*
	Given an array of integers nums, you start with an initial positive value startValue.
	In each iteration, you calculate the step by step sum of startValue plus elements in nums (from left to right).
	Return the minimum positive value of startValue such that the step by step sum is never less than 1.
	Example:
	Input: nums = [-3,2,-3,4,2]
	Output: 5
	Explanation: If you choose startValue = 4, in the third iteration your step by step sum is less than 1.
	step by step sum
	startValue = 4 | startValue = 5 | nums
	  (4 -3 ) = 1  | (5 -3 ) = 2    |  -3
	  (1 +2 ) = 3  | (2 +2 ) = 4    |   2
	  (3 -3 ) = 0  | (4 -3 ) = 1    |  -3
	  (0 +4 ) = 4  | (1 +4 ) = 5    |   4
	  (4 +2 ) = 6  | (5 +2 ) = 7    |   2
*/
func MinStartValue(nums []int) int {
	fmt.Println()
	fmt.Println("Hello from MinStartValue")
	fmt.Println("Nums:", nums)
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = nums[0]
	minPrefix := prefix[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + nums[i]
		if prefix[i] < minPrefix {
			minPrefix = prefix[i]
		}
	}
	fmt.Println("Min prefix:", minPrefix)
	fmt.Println(prefix)
	if minPrefix > 0 {
		fmt.Println("Result:", minPrefix)
		return 1
	} else {
		fmt.Println("Result:", int(math.Abs(float64(minPrefix))+1))
		return 1 - minPrefix
	}
}
