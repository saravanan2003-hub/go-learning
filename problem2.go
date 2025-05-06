// Find the Largest of Three Numbers
// Write a program that takes three numbers as input and finds the largest of them.

// Example:

// Input: 5, 9, 2

// Output: 9

package main

import "fmt"

func find_max_number(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > c && b > c {
		return b
	} else {
		return c
	}
}

func main() {
	fmt.Println(find_max_number(100, 30, 70))
}
