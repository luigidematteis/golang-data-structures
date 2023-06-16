package arrays

import "fmt"

func FindSumOfTwoNumbersEqualToTargetGUA(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, ok := seen[complement]; ok {
			fmt.Println("A pair of values that sum up to the target was found.")
			fmt.Println("The target is:", target)
			fmt.Println("The values are:", nums[j], nums[i])
			fmt.Println("Indexes are:", j, i)
			return []int{j, i}
		}
		seen[num] = i
	}

	return nil
}
