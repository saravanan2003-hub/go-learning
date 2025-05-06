// Prime Number Checker
// Write a program that checks if a given number is prime. A prime number is a number greater than 1 that has no divisors other than 1 and itself.

// Example:

// Input: 7

// Output: True (7 is prime)

// Input: 10

// Output: False (10 is not prime)

package main

import (
	"fmt"
	"math"
)

func prime(n int) bool {
	var isPrime = true
	if n < 1 {
		isPrime = false
	} else {
		for i := 2; i < int(math.Sqrt(float64(n))); i++ {
			if n%i == 0 {
				isPrime = false

			}
		}
	}

	return isPrime
}

func main() {
	fmt.Println(prime(7))
}
