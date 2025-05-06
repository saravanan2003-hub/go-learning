package main

import "fmt"

func main() {
	// a := 10
	// b := 20
	// c := 30
	// d := 40
	// e := 50

	// defer fmt.Println(a)
	// defer fmt.Println(b)
	// defer fmt.Println(c)
	// defer fmt.Println(d)
	// defer fmt.Println(e)

	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}

}
