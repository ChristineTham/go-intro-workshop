package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var age int

	fmt.Println("Initial value of age =", age)
	fmt.Println("What is your age?")
	fmt.Scan(&age)
	fmt.Println("Your age is", age)

	if age >= 18 {
		fmt.Println("You can drive in Australia")
	}

	if age >= 16 {
		fmt.Println("You can drive in the USA")
	}

	var name string

	fmt.Println("Initial value of name =", name)
	fmt.Println("What is your full name?")
	reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n')
	fmt.Println("Your name is", name)
}
