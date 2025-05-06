// Find the Average of an Array
// Write a program that calculates the average of an array of numbers.

// Example:

// Input: [1, 2, 3, 4, 5]

// Output: 3

package main

import (
	"fmt"
)

func average(n []int) int {
	var sum int = 0
	for i := 0; i < len(n); i++ {
		sum += n[i]
	}
	return sum / len(n)
}

func main() {
	n := []int{1, 2, 3, 4, 5}
	fmt.Println(average(n))
	fmt.Println(n)
}
