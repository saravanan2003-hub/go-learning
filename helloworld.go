// 1. Sum of Natural Numbers
// Write a program that computes the sum of the first n natural numbers. The program should take n as input from the user.

// Example:

// Input: 5

// Output: 15 (1 + 2 + 3 + 4 + 5 = 15)

package main

import "fmt"

func natural_numbers(n int) int {
	var sum int = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(natural_numbers(5))
}
