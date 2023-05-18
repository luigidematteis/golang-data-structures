package main

import "fmt"
import "learning/arrays"

func main() {
	fmt.Println("Hello, World!")

	// check if the words are palindromes
	//letMeKnowIfPalindrome("Hello")
	//letMeKnowIfPalindrome("racecar")
	//letMeKnowIfPalindrome("abcdedcba")

	// check if in the given arr there is a pair of value that is the sum of the target
	//nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//arrays.TwoSum(nums, 112)

	arrays.FindSumOfTwoNumbersEqualToTargetGivenAnUnsortedArray([]int{3, 2, 4}, 6)
}

func letMeKnowIfPalindrome(word string) {
	if arrays.TwoPointers(word) {
		fmt.Println("The word", word, "is a palindrome")
	} else {
		fmt.Println("The word", word, "is not a palindrome")
	}
}
