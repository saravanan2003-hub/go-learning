package main

import "fmt"

func main() {
	day := 7

	switch day {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("A week have only seven daysx")
	}

}
