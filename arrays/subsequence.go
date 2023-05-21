package arrays

func IsSubsequence(string1 string, string2 string) bool {
	var i = 0
	var j = 0

	for i < len(string1) && j < len(string2) {
		if string1[i] == string2[j] {
			i++
		}
		j++
	}

	return i == len(string1)
}
