package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr)

	slice := arr[1:5]
	fmt.Println(slice)

	fmt.Println(arr[:4])
	fmt.Println(arr[3:])

	fmt.Println(arr[:])

	s := make([]int, 6, 10)
	fmt.Println(s)      // [0 0 0]
	fmt.Println(len(s)) // 3
	fmt.Println(cap(s)) // 5
}
