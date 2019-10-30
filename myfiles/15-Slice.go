package main

import (
	"fmt"
	"strings"
)

func main() {
	letters := []string{"a", "b", "c", "d"}

	fmt.Println("First in the slice", letters[0])
	fmt.Println("Number of elements in the letters slice", len(letters))

	for index := 0; index < len(letters); index++ {
		letters[index] = strings.ToUpper(letters[index])
		fmt.Println(index, "in the slice", letters[index])
	}
}
