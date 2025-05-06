//multiple contition in switch case

package main

import "fmt"

func main() {

	day := 2
	switch day {
	case 1, 3, 5:
		fmt.Println("Odd week days")
	case 2, 4:
		fmt.Println("Even week days")
	case 6, 7:
		fmt.Println("Weekend days")
	}

}
