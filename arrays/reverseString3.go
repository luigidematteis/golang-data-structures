package arrays

import (
	"fmt"
	"strings"
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
