package arrays

import "fmt"

func CombineArrays(array1 []int, array2 []int) {

	var combinedArrays []int

	i := 0
	j := 0

	for i < len(array1) && j < len(array2) {
		if array1[i] < array2[j] {
			combinedArrays = append(combinedArrays, array1[i])
			i++
		} else {
			combinedArrays = append(combinedArrays, array2[j])
			j++
		}

		fmt.Println(combinedArrays)
	}

	for i < len(array1) {
		combinedArrays = append(combinedArrays, array1[i])
		i++
		fmt.Println(combinedArrays, "Array1")
	}

	for j < len(array2) {
		combinedArrays = append(combinedArrays, array2[j])
		j++
		fmt.Println(combinedArrays, "Array2")
	}

	fmt.Println("Given arrays:", array1, array2)
	fmt.Println("Combined array:", combinedArrays)
}
