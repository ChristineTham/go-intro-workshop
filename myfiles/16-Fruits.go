package main

import (
	"fmt"
	"strings"
)

func main() {
	fruits := []string{"apple", "Banana", "cherry", "orange", "watermelon", "Mango", "Papaya", "Blueberry", "plum", "Peach", "pear", "Pineapple", "Raspberry", "Strawberry", "lemon", "lime"}
	fruits = append(fruits, "dragonfruit", "jackfruit", "durian")

	fmt.Println(fruits)
	fmt.Println("Number of elements in the fruits slice", len(fruits))

	fmt.Println("Index , fruit , Title , ToUpper , ToLower")
	for index, fruit := range fruits {
		fmt.Println(index, ",", fruit, ",", strings.Title(fruit), ",", strings.ToUpper(fruit), ",", strings.ToLower(fruit))
	}
}
