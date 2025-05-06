package main

import "fmt"

type User struct {
	name  string
	email string
	age   int
}

func main() {
	user1 := User{name: "Saravanan", email: "saravanan@gmail.com", age: 21}
	user2 := User{name: "pranesh", email: "pranesh@gmail.com", age: 21}

	fmt.Println(user1.age)
	fmt.Println(user2.email)
}
