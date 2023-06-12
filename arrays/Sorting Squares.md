# Sorting Function
### Language used: Golang

```
/*
	Sorting function with time complexity of O(n log n) and space complexity of O(n)
*/
func SortedSquares(nums []int) []int {
	fmt.Println("Original array:", nums)
	squaredNums := make([]int, 0, len(nums)-1)

	for _, num := range nums {
		squaredNums = append(squaredNums, num*num)

		if len(nums) == 1 {
			return squaredNums
		}
	}

	var i = 0
	sortedArray := make([]int, len(nums))

	for i <= len(squaredNums)-1 {
		var currNumberIndex = 0
		var currNumber = squaredNums[i]

		for _, num := range squaredNums {
			if currNumber > num {
				currNumberIndex++
			}
		}

		for currNumber != 0 && currNumber == sortedArray[currNumberIndex] {
			currNumberIndex++
		}

		sortedArray[currNumberIndex] = currNumber

		i++
	}

	fmt.Println("Sorted array:", sortedArray)
	return sortedArray
}

```
### Function Analysis

```
  squaredNums := make([]int, 0, len(nums)-1)

	for _, num := range nums {
		squaredNums = append(squaredNums, num*num)

		if len(nums) == 1 {
			return squaredNums
		}
	}
```
In this part of code we create a new array for storing the squared values.\
In case we had as input an array with only one member, we directly return it as result.

```
    var i = 0
    sortedArray := make([]int, len(nums))

	for i <= len(squaredNums)-1 {
		var currNumberIndex = 0
		var currNumber = squaredNums[i]

```

Then we declared a variable *i*
that we will use to iterate on each square in our array, and we create another array *sortedArray* that has the same size of the array we had as input.

We start the iteration assigning each time the value of *squaredNums[i]* to a variable *currNumber*.
We also declared a variable *currNumberIndex* that will be used to store the final index that the value will have.

 ```
        for _, num := range squaredNums {
			if currNumber > num {
				currNumberIndex++
			}
		}

```

In this nested loop, we iterate on the *squaredNums* arrays starting from the index 0 and comparing the *currNumber* value with the value at each specific index.
If the current value will be greater that the number on wich we are iterating, then the index will increase by 1.
```
  for currNumber != 0 && currNumber == sortedArray[currNumberIndex] {
			currNumberIndex++
  }

```
Before to store our value at the index found, we check for duplicate at the current index.
If a duplicate is present *currNumber == sortedArray[currNumberIndex]*, then the index in wich the current value will be stored will be increased accordingly.

```
    sortedArray[currNumberIndex] = currNumber

    i++

```

Finally we store the value at the index calculated and keep iterating toward the next element.

* * *

This approach sort an array calculating the index in which each value should be stored.\
It comes up from the assumption that: *if in array of 5 elements, a value has 4 number that are smaller, it should be stored as the last element.*\
Therefore, we can calculate the index in which an element should be stored by knowing how many number are smaller than a certain value.