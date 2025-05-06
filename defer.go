package main

import "fmt"

func main() {
	var a int = 10
	defer fmt.Println(a)
	a = 20
	fmt.Println(a)
}
