package arrays

import "fmt"

/*
	Sorting function with time complexity of O(n log n) and space complexity of O(n)
*/
func SortedSquares(nums []int) []int {
	fmt.Println("Original array:", nums)
	squaredNums := make([]int, 0, len(nums)-1)

	for _, num := range nums {
		squaredNums = append(squaredNums, num*num)

		if len(nums) == 1 {
			return squaredNums
		}
	}

	var i = 0
	sortedArray := make([]int, len(nums))

	for i <= len(squaredNums)-1 {
		var currNumberIndex = 0
		var currNumber = squaredNums[i]

		for _, num := range squaredNums {
			if currNumber > num {
				currNumberIndex++
			}
		}

		for currNumber != 0 && currNumber == sortedArray[currNumberIndex] {
			currNumberIndex++
		}

		sortedArray[currNumberIndex] = currNumber

		i++
	}

	fmt.Println("Sorted array:", sortedArray)
	return sortedArray
}
