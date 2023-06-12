package arrays

import "fmt"

func SortNumsArray(nums []int) []int {

	fmt.Println("Original array:", nums)

	squaredNums := make([]int, 0, len(nums)-1)

	for _, num := range nums {
		squaredNums = append(squaredNums, num*num)

		if len(nums) == 1 {
			fmt.Println("The array contains only 1 member:", squaredNums)
			return squaredNums
		}
	}

	fmt.Println("Squared array:", squaredNums)

	var i = 0
	sortedArray := make([]int, len(nums))

	for i <= len(squaredNums)-1 {
		var currNumberIndex = 0
		var currNumber = squaredNums[i]

		for j, num := range squaredNums {
			fmt.Println("Current value", num, "at index", j, "is compared with", currNumber)
			if currNumber > num {
				currNumberIndex++
			}
		}

		for _, num := range sortedArray {
			if currNumber != 0 && currNumber == num {
				currNumberIndex++
			}
		}

		fmt.Println("The value", currNumber, "will be inserted at index:", currNumberIndex)
		sortedArray[currNumberIndex] = currNumber

		fmt.Println("Array at this step:", sortedArray)
		i++
	}

	fmt.Println("Sorted array:", sortedArray)
	return sortedArray
}
