package main

import "fmt"

func main() {
	var age int

	fmt.Println("Initial value of age =", age)
	fmt.Println("What is your age?")
	fmt.Scan(&age)
	fmt.Println("Your age is", age)

	var name string

	fmt.Println("Initial value of name =", name)
	fmt.Println("What is your name?")
	fmt.Scan(&name)
	fmt.Println("Your name is", name)
}
