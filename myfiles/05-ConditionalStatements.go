package main

import "fmt"

func main() {
	justRight := 88
	var temperature int
	fmt.Println("What temperature is the soup?")
	fmt.Scan(&temperature)

	if temperature == justRight {
		fmt.Println("The soup is just right")
	} else if temperature > justRight {
		fmt.Println("The soup TOO hot")
	} else {
		fmt.Println("The soup is not right")
	}
}
