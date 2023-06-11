package main

import "fmt"
import "learning/arrays"

func main() {
	fmt.Println("Hello, World!")

	// check if the words are palindromes
	letMeKnowIfPalindrome("Hello")
	letMeKnowIfPalindrome("racecar")
	letMeKnowIfPalindrome("abcdedcba")

	// check if in the given arr there is a pair of value that is the sum of the target
	arrays.TwoSum([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 112)

	arrays.FindSumOfTwoNumbersEqualToTargetGivenAnUnsortedArray([]int{3, 2, 4}, 6)

	// given two sorted integer arrays, return an array that combines both of them and is also sorted
	arrays.CombineArrays([]int{1, 4, 7, 20, 32}, []int{3, 5, 6})

	// given two strings s and t, return true if s is a subsequence of t, or false otherwise
	var isSubsequence = arrays.IsSubsequence("ace", "abcde")
	fmt.Println(isSubsequence)

	// reverse strings
	arrays.ReverseString([]byte("hello"))

	arrays.ReverseString2([]byte("Hannah"))

	// square of a sorted array
	arrays.SquareOfSortedArray([]int{-4, -1, 0, 3, 10})
}

func letMeKnowIfPalindrome(word string) {
	if arrays.TwoPointers(word) {
		fmt.Println("The word", word, "is a palindrome")
	} else {
		fmt.Println("The word", word, "is not a palindrome")
	}
}
