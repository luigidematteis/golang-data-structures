package arrays

import "fmt"

/*
	MaxConsecutiveOnes return the maximum number of consecutive 1's in the array,
	given a binary array nums and an integer k, which indicate that you can flip at most k 0's.
*/
func MaxConsecutiveOnes(nums []int, k int) int {
	var left = 0
	var zeroCount = 0
	var result = 0

	for i, num := range nums {
		if num == 0 {
			zeroCount += 1
		}
		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount -= 1
			}
			left++
		}
		if i-left+1 > result {
			result = i - left + 1
		}
	}
	fmt.Println("Result:", result)
	return result
}
