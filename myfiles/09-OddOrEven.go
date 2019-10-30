package main

import "fmt"

func main() {
	var number int
	fmt.Println("What is the number?")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}
}
