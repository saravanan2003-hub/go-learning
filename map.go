package main

import "fmt"

func main() {
	ages := make(map[string]int)

	ages["saravanan"] = 21
	ages["pranesh"] = 21

	fmt.Println(ages["saravanan"])
}
