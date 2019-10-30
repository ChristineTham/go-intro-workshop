package main

import "fmt"

func main() {
	var max int
	fmt.Println("Enter triangle size")
	fmt.Scan(&max)
	// max = 7

	for i := 0; i < max; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := max; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
