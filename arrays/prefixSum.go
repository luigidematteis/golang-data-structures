package arrays

import "fmt"

/*
	Given an integer array nums, an array queries where queries[i] = [x, y] and an integer limit, AnswerQueries return
	a boolean array that represents the answer to each query. A query is true if the sum of the subarray from x to y is
	less than limit, or false otherwise.
	For example, given nums = [1, 6, 3, 2, 7, 2], queries = [[0, 3], [2, 5], [2, 4]], and limit = 13, the answer is
	[true, false, true]. For each query, the subarray sums are [12, 14, 12].
*/
func AnswerQueries(nums []int, queries [][]int, limit int) []bool {
	fmt.Println()
	fmt.Println("Hello from AnswerQueries")
	// For example, given nums = [1, 6, 3, 2, 7, 2], queries = [[0, 3], [2, 5], [2, 4]], and limit = 13,
	// the answer is [true, false, true]

	// create slices for prefix sum array
	prefix := make([]int, len(nums))

	// assign the first element of nums to prefix[0]
	prefix[0] = nums[0]

	// process the prefix array
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	fmt.Println("Processed prefix array:", prefix)

	// create a bool array for the answers
	ans := make([]bool, len(queries))
	for i := 0; i < len(queries); i++ {
		// declare variables that represent the given indexes of each query
		x, y := queries[i][0], queries[i][1]
		// using the formula prefix[j] - prefix[i] + nums[i] to calculate the
		// sum of the subarray from the index x to y
		// we didn't use the formula prefix[j] - prefix[i - 1] for not dealing with case x = 0
		curr := prefix[y] - prefix[x] + nums[x]
		// calculate the answer
		ans[i] = curr < limit
	}
	fmt.Println("Result:", ans)
	return ans
}

/*
	Given an integer array nums, find the number of ways to split the array into two parts so that the first section
	has a sum greater than or equal to the sum of the second section. The second section should have at least one
	number.
*/
func WaysToSplitArray(nums []int) int {
	fmt.Println()
	fmt.Println("Hello from WaysToSplitArray")
	// declare variable for iterations
	n := len(nums)

	// create slice for prefix array
	prefix := make([]int, len(nums))

	// assign the value of nums[0] as first element of the prefix array
	prefix[0] = nums[0]

	// process the prefix array
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	// declare variable for the result
	ans := 0

	// calculate how many subarray that respect the constraint are present
	for i := 0; i < n-1; i++ {
		// the sum of the left section is retrieved by taking the i element from the prefix array
		sumOfLeftSection := prefix[i]
		// the sum of the right section is calculated by taking the last element of the prefix array
		// that correspond to the sum of all the elements in the nums array and subtracting the prefix[i]
		// that correspond to the sum of all the elements until the index i
		sumOfRightSection := prefix[n-1] - prefix[i]
		// if the sum of left section is greater than the right one, the answer will be updated
		if sumOfLeftSection >= sumOfRightSection {
			ans++
		}
	}
	fmt.Println("Result:", ans)
	return ans
}

// WaysToSplitArray2 is the improved version of WaysToSplitArray with space complexity of O(1)
func WaysToSplitArray2(nums []int) int {
	fmt.Println()
	fmt.Println("Hello from WaysToSplitArray2")
	ans, leftSection, total := 0, 0, 0
	for _, num := range nums {
		total += num
	}
	for i := 0; i < len(nums)-1; i++ {
		leftSection += nums[i]
		rightSection := total - leftSection
		if leftSection >= rightSection {
			ans++
		}
	}
	fmt.Println("Result:", ans)
	return ans
}

/*
	Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
	Return the running sum of nums.
*/
func RunningSum(nums []int) []int {
	fmt.Println()
	fmt.Println("Hello from RunningSum")
	prefix := make([]int, len(nums))
	prefix[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}
	fmt.Println("Result:", prefix)
	return prefix
}
