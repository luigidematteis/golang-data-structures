package arrays

import "fmt"

/*
	You are given a 0-indexed array nums of n integers, and an integer k.
	The k-radius average for a subarray of nums centered at some index i with the radius k is the average of all
	elements in nums between the indices i - k and i + k (inclusive). If there are less than k elements before or
	after the index i, then the k-radius average is -1.
	Build and return an array avgs of length n where avgs[i] is the k-radius average for the subarray centered
	at index i.
	The average of x elements is the sum of the x elements divided by x, using integer division. The integer division
	truncates toward zero, which means losing its fractional part.
	Example 1:
	Input: nums = [7,4,3,9,1,8,5,2,6], k = 3
	Output: [-1,-1,-1,5,4,4,-1,-1,-1]
	Explanation:
	- avg[0], avg[1], and avg[2] are -1 because there are less than k elements before each index.
	- The sum of the subarray centered at index 3 with radius 3 is: 7 + 4 + 3 + 9 + 1 + 8 + 5 = 37.
	  Using integer division, avg[3] = 37 / 7 = 5.
	- For the subarray centered at index 4, avg[4] = (4 + 3 + 9 + 1 + 8 + 5 + 2) / 7 = 4.
	- For the subarray centered at index 5, avg[5] = (3 + 9 + 1 + 8 + 5 + 2 + 6) / 7 = 4.
	- avg[6], avg[7], and avg[8] are -1 because there are less than k elements after each index.
*/
func GetAverage(nums []int, k int) []int {
	fmt.Println("Array:", nums)
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}
	fmt.Println("Prefix Array:", prefix)

	kRadius := make([]int, n)

	for i := 0; i < n; i++ {
		if i >= k && i < n-k {
			temp := prefix[i+k]
			if i-k > 0 {
				temp -= prefix[i-k-1]
			}
			fmt.Println()
			fmt.Println("We should find the prefix at current index + 3 that is:", prefix[i+3])
			fmt.Println("And divide it by the number of elements in the subarray, calculated as k * 2 + 1 =", k*2+1)
			fmt.Println("The result will be obtained by (prefix[i+3])/(k*2+1):", (prefix[i+3])/(k*2+1))
			kRadius[i] = temp / (k*2 + 1)
		} else {
			kRadius[i] = -1
		}
	}

	fmt.Println(kRadius)

	return kRadius
}
