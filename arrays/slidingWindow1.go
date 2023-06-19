package arrays

import (
	"fmt"
)

/*
	FindLengthOfSubarrayWithSumLessThanK find the length of the longest subarray with a sum less than or equal to k
*/
func FindLengthOfSubarrayWithSumLessThanK(nums []int, k int) int {
	fmt.Println()
	fmt.Println("Hello from Sliding Window: FindLengthOfSubarrayWithSumLessThanK")

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

		result = Max(result, i-left+1)
	}

	fmt.Println("Result:", result)
	return result
}

/*
	FindNumOfSubarrayWithProductLessThanK find the number of subarray with product less than K
*/
func FindNumOfSubarrayWithProductLessThanK(nums []int, k int) int {
	fmt.Println()
	fmt.Println("Hello from Sliding Window: FindNumOfSubarrayWithProductLessThanK")

	if k < 1 {
		return 0
	}

	var left = 0
	var result = 0
	var currentProduct = 1

	for i, num := range nums {
		currentProduct *= num

		for currentProduct > k {
			currentProduct /= num
			left++
		}

		result = i - left + 1
	}

	fmt.Println("Result:", result)
	return result
}

/*
	MaximumAverageSubarray find a contiguous subarray whose length is equal to k
	that has the maximum average value and return this value
*/
func MaximumAverageSubarray(nums []int, k int) float64 {
	fmt.Println()
	fmt.Println("Hello from Sliding Window: MaximumAverageSubarray")

	if k < 1 {
		return 0
	}

	if len(nums) == 1 {
		return float64(nums[0])
	}

	var left = 0
	var currentValue = 0
	var windowLenght = 0
	maxAverage := float64(0)
	result := float64(0)

	for i, num := range nums {

		currentValue += num
		windowLenght = i - left + 1
		fmt.Println("Dividing current value", currentValue, "by", windowLenght)
		var currentAverage = float64(currentValue) / float64(windowLenght)
		fmt.Println("Calculated current average:", currentAverage)

		for windowLenght > k {
			fmt.Println("Window lenght was greater than k")
			currentValue -= nums[left]
			left++
			windowLenght = i - left + 1
			currentAverage = float64(currentValue) / float64(windowLenght)
		}

		fmt.Println("Window lenght:", windowLenght, "K:", k)
		if windowLenght == k {
			if left == 0 && maxAverage == 0 {
				maxAverage = currentAverage
			}
			maxAverage = maxFloat64(maxAverage, currentAverage)
			fmt.Println("Max average calculated:", maxAverage)
			result = maxAverage
		}
	}

	println("Result:", result)
	return result
}

func Maxaveragesubarray2(nums []int, k int) float64 {
	var left = 0
	var currentTotal int
	var maxTotal int
	for i, num := range nums {
		if i-left+1 > k {
			currentTotal -= nums[left]
			left++
		}
		currentTotal += num
		if currentTotal > maxTotal || left == 0 {
			maxTotal = currentTotal
		}
	}
	fmt.Println("Result:", float64(maxTotal)/float64(k))
	return float64(maxTotal) / float64(k)
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxFloat64(x float64, y float64) float64 {
	if x > y {
		return x
	}
	return y
}
