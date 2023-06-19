package arrays

import (
	"fmt"
	"strings"
)

/*
	Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.
	Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.
	Example 3:
	Input: s = "leetcode", k = 3
	Output: 2
	Explanation: "lee", "eet" and "ode" contain 2 vowels.
*/
func MaxVowelsInSubstring(s string, k int) int {
	left := 0
	ans := 0
	vowCount := 0
	const vowels = "aeiou"

	if k == 1 && strings.ContainsAny(s, vowels) {
		return 1
	}

	for i := 0; i < len(s); i++ {
		if strings.Contains(vowels, string(s[i])) {
			vowCount += 1
		}
		for i-left+1 > k {
			if strings.Contains(vowels, string(s[left])) {
				vowCount -= 1
			}
			left++
		}
		ans = max(ans, vowCount)
	}
	fmt.Println(ans)
	return ans
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
