package arrays

import "fmt"

func TwoSum(nums []int, target int) {
	fmt.Println("Hello from the function TwoSum")

	sumExists, _, _ := findSumOfTwoNumbersEqualToTarget(nums, target)

	if !sumExists {
		fmt.Println("A pair of values that sum up to the target was not found.")
	}
}

func findSumOfTwoNumbersEqualToTarget(nums []int, target int) (bool, int, int) {
	left := 0
	right := len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]

		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			fmt.Println("A pair of values that sum up to the target was found.")
			fmt.Println("The target is:", target)
			fmt.Println("The values are:", nums[left], nums[right])
			fmt.Println("Indexes are:", left, right)
			return true, left, right
		}
	}

	return false, 0, 0
}
