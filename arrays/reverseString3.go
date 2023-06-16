package arrays

import (
	"fmt"
	"strings"
	"unicode"
)

/*
	Given a string s, reverse the order of characters in each word within a sentence while still preserving
	whitespace and initial word order.
	Example 1:
	Input: s = "Let's take LeetCode contest"
	Output: "s'teL ekat edoCteeL tsetnoc"
*/
func ReverseWords(s string) string {
	left := 0
	st := strings.Split(s, "")
	rs := make([]string, 0, len(s))
	for i, c := range st {
		var temp = i
		fmt.Println("We are on the letter:", c)
		if c == " " || i == len(s)-1 {
			if i == len(s)-1 {
				temp++
			}
			for left < temp {
				rs = append(rs, st[temp-1])
				temp--
			}
			if i != len(s)-1 {
				rs = append(rs, " ")
			}
			left = i
			left++
		}
	}
	fmt.Println(strings.Join(rs, ""))
	return strings.Join(rs, "")
}

// ReverseWords2
// Input: s = "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"
func ReverseWords2(s string) string {
	n := len(s)
	var ans []byte
	checkpoint := -1
	for i := 0; i < n; i++ {
		if s[i] == ' ' || i == n-1 {
			isLastIndex := false
			if i == n-1 {
				i++
				isLastIndex = true
			}
			for j := i - 1; j > checkpoint; j-- {
				ans = append(ans, s[j])
			}
			if !isLastIndex {
				ans = append(ans, byte(' '))
			}
			checkpoint = i
		}
	}
	fmt.Println(string(ans))
	return string(ans)
}

/*
	Given a string s, reverse the string according to the following rules:
	All the characters that are not English letters remain in the same position.
	All the English letters (lowercase or uppercase) should be reversed.
	Return s after reversing it.
*/
func ReverseOnlyLetters(s string) string {
	ans := make([]byte, len(s))
	k := 0
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Println()
		fmt.Println("Letter at index:", i, "is:", string(s[i]))
		if !unicode.IsLetter(rune(s[i])) {
			ans[i] = s[i]
			fmt.Println("Array at this step:", string(ans))
		} else {
			fmt.Println("If s[k] will be a non letter we increment k:", string(s[k]))
			fmt.Println(!unicode.IsLetter(rune(s[k])))
			for !unicode.IsLetter(rune(s[k])) {
				ans[k] = s[k]
				k++
				fmt.Println("Array at this step:", string(ans))
			}
			ans[k] = s[i]
			k++
			fmt.Println("Array at this step:", string(ans))
		}
	}
	fmt.Println(string(ans))
	return string(ans)
}

/*
	Given a 0-indexed string word and a character ch, reverse the segment of word that starts at index 0 and ends at
	the index of the first occurrence of ch (inclusive). If the character ch does not exist in word, do nothing.
	For example, if word = "abcdefd" and ch = "d", then you should reverse the segment that starts at 0 and ends
	at 3 (inclusive). The resulting string will be "dcbaefd".
	Return the resulting string.
	Example:
	Input: word = "abcdefd", ch = "d"
	Output: "dcbaefd"
*/
func reversePrefix(word string, ch byte) string {
	w := -1
	n := len(word)
	rs := make([]byte, n)
	for i := 0; i < n; i++ {
		if word[i] == ch && i != 0 && w == -1 {
			r := i
			for j := 0; j <= r; j++ {
				fmt.Println("j ->", j)
				fmt.Println("r ->", r)
				fmt.Println(string(rs[j]), string(word[r]))
				fmt.Println(string(rs[r]), string(word[j]))
				rs[j] = word[r]
				rs[r] = word[j]
				r--
				fmt.Println("j ->", j)
				fmt.Println("r ->", r)
				fmt.Println("Array at this step:", string(rs))
			}
			w = i
		}
		if w != -1 && w != i {
			rs[i] = word[i]
		}
	}
	if rs[0] != ch || word[0] == ch || word[n-1] == ch && w == -1 {
		fmt.Println("return word")
		return word
	}
	fmt.Println(string(rs))
	fmt.Println("return rs")
	return string(rs)
}
