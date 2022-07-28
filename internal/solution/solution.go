package solution

import (
	"sort"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

/*
This is a demo task.

Write a function:func Solution(A []int) int

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.

Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].

*/

func isNegative(num int) bool {
	return num < 0
}

func Solution(A []int) int {
	// write your code in Go 1.4
	// fmt.Printf("input: %v\n", A)

	sort.Ints(A)

	// fmt.Printf("sorted: %v\n", A)

	size := len(A)

	result := 1
	for i := 0; i < size; i++ {
		if isNegative(A[i]) {
			return result
		} else if A[i] == result {
			result += 1
		} else if A[i] > result {
			return result
		}
	}

	return result
}
