package arrays

import "fmt"

func SquareOfSortedArray(nums []int) []int {
	squaredNums := make([]int, 0, len(nums))

	for _, num := range nums {
		squaredNums = append(squaredNums, num*num)
	}
	fmt.Println("Squared nums: ", squaredNums)

	var i = 0

	for i <= len(squaredNums)-1 {
		fmt.Println("Iteration number:", i+1)

		var w = 1

		var curr = squaredNums[i]

		fmt.Println("Array status at this step:", squaredNums)
		fmt.Println("Current value", curr, "at index", i, "is compared with", squaredNums[w], "at index:", w)
		for curr > squaredNums[w] {
			fmt.Println("The value", curr, "is greater than", squaredNums[w], "at index", w)
			w++
		}

		if curr < squaredNums[w] {
			w--
		}

		fmt.Println("The current value", curr, "will replace the value", squaredNums[w], "at index:", w)

		var temp = squaredNums[w]
		squaredNums[w] = curr
		squaredNums[i] = temp
		i++

		fmt.Println("Array status at this step:", squaredNums)
		fmt.Println()
	}

	fmt.Print("Sorted array:", squaredNums)
	return squaredNums
}
