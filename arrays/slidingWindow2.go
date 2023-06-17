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

/*
	Given an array of positive integers nums and a positive integer target, return the minimal length of a
	subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.
	Example 1:
	Input: target = 7, nums = [2,3,1,2,4,3]
	Output: 2
	Explanation: The subarray [4,3] has the minimal length under the problem constraint.
*/
func MinSubarrayLenght(nums []int, target int) int {
	fmt.Println()
	fmt.Println("Array:", nums)
	fmt.Println("Target:", target)
	l := 0
	totalSum := 0
	sum := 0
	w := 0
	for i, num := range nums {
		totalSum += num
		fmt.Println("Iteration:", i, "with number:", num)
		sum += num
		fmt.Println("Sum:", sum)
		if i == len(nums)-1 && w == 0 && sum == target {
			fmt.Println("Return i", i+1)
			i++
			return i
		}
		if sum > target {
			sum = num
			l = i - 1
			for sum < target && l > 0 {
				sum += nums[l]
				l--
				fmt.Println("Sum:", sum)
			}
			if w == 0 || i-l < w {
				if i == 1 {
					i++
				}
				w = i - l
			}
			fmt.Println("Min subarray length found at this moment:", w)
			fmt.Println("Sum calculated at this moment:", sum)
		}
	}
	fmt.Println("Sum calculated at this moment as result is", sum, "and target is", target)
	if totalSum < target {
		fmt.Println("Return zero")
		return 0
	}
	fmt.Println("Return w that is:", w)
	return w
}
