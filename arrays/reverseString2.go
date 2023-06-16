package arrays

import "fmt"

func ReverseString2(s []byte) {
	fmt.Println("Hello from ReverseString2")

	var i = len(s) - 1
	var j = 0
	for i >= 0 && j <= len(s)-1 {
		if j >= i {
			break
		}
		var temp = s[j]
		s[j] = s[i]
		s[i] = temp
		i--
		j++
	}

	fmt.Println(string(s))
}
