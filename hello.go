package main

import "fmt"
import "learning/arrays"

func main() {
	fmt.Println("Hello, World!")
	letMeKnowIfPalindrome("Hello")
	letMeKnowIfPalindrome("racecar")
	letMeKnowIfPalindrome("abcdedcba")
}

func letMeKnowIfPalindrome(word string) {
	if arrays.TwoPointers(word) {
		fmt.Println("The word", word, "is a palindrome")
	} else {
		fmt.Println("The word", word, "is not a palindrome")
	}
}
