package arrays

func ReverseString(s []byte) {
	var i = len(s) - 1
	var j = 0
	var rs []byte
	for i >= 0 {
		rs = append(rs, s[i])
		i--
	}
	for j <= len(s)-1 {
		s[j] = rs[j]
		j++
	}
}
