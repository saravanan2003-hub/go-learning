//	func (receiver TypeName) MethodName(params) returnType {
//	    // code
//	} //// it is a basic syntax for method
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) sayhello() {
	fmt.Println("Welcome to Go launuage ", p.name, "your age is ", p.age)
	return
}

func (p *Person) update_name(new_name string) {
	p.name = new_name
}

func main() {
	p := Person{name: "Saravanan", age: 21}
	p.sayhello()

	p1 := Person{name: "Pranesh", age: 21}
	p1.sayhello()

	p2 := Person{name: "Bala", age: 19}
	p2.sayhello()

	p.update_name("kumar")
	p.sayhello()
}
