package arrays

import "fmt"

func FindSumOfTwoNumbersEqualToTargetGivenAnUnsortedArray(nums []int, target int) []int {

	var startLeftIterator = 0
	var startRightIterator = len(nums) - 1

	var endLeftIterator = len(nums) - 1
	var endRightIterator = 0

	for startLeftIterator < endLeftIterator {
		fmt.Println("Start iteration from left to right")
		sum := nums[startLeftIterator] + nums[endLeftIterator]

		if sum > target {
			endLeftIterator--
		} else if sum < target {
			startLeftIterator++
		}

		fmt.Println(sum, target)
		if sum == target {
			fmt.Println("Iterator from left to right:")
			fmt.Println("A pair of values that sum up to the target was found.")
			fmt.Println("The target is:", target)
			fmt.Println("The values are:", nums[startLeftIterator], nums[endLeftIterator])
			fmt.Println("Indexes are:", startLeftIterator, endLeftIterator)
			return []int{startLeftIterator, endLeftIterator}
		}
	}

	for startRightIterator > endRightIterator {
		fmt.Println("Start iteration from right to left")
		sum := nums[startRightIterator] + nums[endRightIterator]

		if sum > target {
			endRightIterator++
		} else if sum < target {
			startRightIterator--
		}

		fmt.Println(sum, target)
		if sum == target {
			fmt.Println("Iterator from right to left:")
			fmt.Println("A pair of values that sum up to the target was found.")
			fmt.Println("The target is:", target)
			fmt.Println("The values are:", nums[startRightIterator], nums[endRightIterator])
			fmt.Println("Indexes are:", startRightIterator, endRightIterator)
			return []int{startRightIterator, endRightIterator}
		}
	}

	return nil
}
