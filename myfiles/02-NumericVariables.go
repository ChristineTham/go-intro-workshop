package main

import "fmt"

func main() {
	x := 70.0
	y := 30.0

	fmt.Println("x =", x, "y =", y)
	fmt.Println("x + y =", x+y)
	fmt.Println("x * y =", x*y)
	fmt.Println("x / y =", x/y)
	// fmt.Println("x % y =", x % y)
	fmt.Println()

	x = 60
	y = 30.1

	fmt.Println("x =", x, "y =", y)
	fmt.Println("x + y =", x+y)
	fmt.Println("x * y =", x*y)
	fmt.Println("x / y =", x/y)
	// fmt.Println("x % y =", x % y)
}
