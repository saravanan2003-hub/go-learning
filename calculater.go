package main

import (
	"errors"
	"fmt"
	"math"
)

func add(a, b float64, operator string) float64 {
	return a + b
}

func sub(a, b float64, operator string) float64 {
	return a - b
}

func mul(a, b float64, opeaator string) float64 {
	return a * b
}

func div(a, b float64, operator string) (interface{}, error) {
	if b == 0 {
		return 0, errors.New("Can't divide by 0")
	}
	return a / b, nil
}

func modules(a, b int, operator string) (interface{}, error) {
	if b == 0 {
		return 0, errors.New("Can't modules by 0")
	}

	return a % b, nil
}

func power(a, b float64) float64 {
	return math.Pow(a, b)
}

func factorial(a int) (interface{}, error) {

	if a < 0 {
		return 0, errors.New("Negative numbers not allowed")
	} else if a == 0 {
		return 1, nil
	} else {
		mul := 1
		for i := 1; i <= a; i++ { // use range
			mul *= i
		}
		return mul, nil
	}
}

func main() {
	var a, b float64
	var operator string

	fmt.Print("Enter your operation symbol : ")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Print("Enter two numbers : ")
		fmt.Scan(&a, &b)

		fmt.Println(fmt.Sprintf("Your answer is  %v", add(a, b, operator)))

	case "-":
		fmt.Print("Enter two numbers : ")
		fmt.Scan(&a, &b)

		fmt.Println(fmt.Sprintf("Your answer is %v", sub(a, b, operator)))

	case "*":
		fmt.Print("Enter two numbers : ")
		fmt.Scan(&a, &b)

		fmt.Println(fmt.Sprintf("Your answer is %v", mul(a, b, operator)))

	case "/":
		fmt.Print("Enter two numbers : ")
		fmt.Scan(&a, &b)

		result, err := div(a, b, operator)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println(fmt.Sprintf("Your answer is %v", result))
		}

	case "%":
		fmt.Print("Enter two numbers : ")
		fmt.Scan(&a, &b)

		result, err := modules(int(a), int(b), operator)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println(fmt.Sprintf("Your answer is %v", result))
		}

	case "^":
		fmt.Print("Enter two numbers : ")
		fmt.Scan(&a, &b)

		fmt.Println(power(a, b))

	case "!":
		fmt.Print("Enter a number : ")
		fmt.Scan(&a)

		result, err := factorial(int(a))
		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println(fmt.Sprintf("Your answer is %v", result))
		}

	default:
		fmt.Println("Enter valid sign or Enter valid operator")
	}

}
