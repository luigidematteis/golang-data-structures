package arrays

import "fmt"

func TwoPointers(word string) bool {
	fmt.Println("Two Pointers example")
	var left int = 0
	var right int = len(word) - 1

	// print the char indexes in the given string
	fmt.Println("Char indexes in the word:", word)
	for i := range word {
		fmt.Println(i)
	}

	fmt.Println("First and last indexes:")
	// print the first and the last index
	fmt.Println(left, right)

	// check if the given word is palindrome or not
	for left < right {
		if word[left] != word[right] {
			return false
		}
		left++
		right--
	}

	return true
}

/*
	Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
	Note that you must do this in-place without making a copy of the array.
	Example 1:
	Input: nums = [0,1,0,3,12]
	Output: [1,3,12,0,0]
*/
func MoveZeros(nums []int) {
	left := 0
	for i, num := range nums {
		if num == 0 || len(nums) == 1 {
			continue
		} else {
			if nums[left] == 0 && nums[i] != 0 {
				nums[left] = num
				nums[i] = 0
			}
			left++
		}
	}
	fmt.Println(nums)
}

func MoveZeros2(nums []int) {
	w := 0
	r := len(nums) - 1
	for i := 0; i <= r; i++ {
		if nums[i] != 0 {
			nums[w] = nums[i]
			w++
		}
	}
	for j := r; j >= w; j-- {
		nums[j] = 0
	}
	fmt.Println(nums)
}

func MoveZeros3(nums []int) {
	w := 0
	r := len(nums) - 1
	for i := 0; i <= r; i++ {
		if nums[i] != 0 {
			nums[w] = nums[i]
			if i != w {
				nums[i] = 0
			}
			w++
		}
	}
	fmt.Println(nums)
}
