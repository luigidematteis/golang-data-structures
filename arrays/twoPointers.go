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
