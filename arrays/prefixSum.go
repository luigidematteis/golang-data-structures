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
