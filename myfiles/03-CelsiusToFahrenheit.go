package main

import "fmt"

func main() {
	celsius := 32.0
	fahrenheit := celsius*9/5 + 32
	fmt.Printf("%.2f celsius to fahrenheit %.2f\n", celsius, fahrenheit)

	celsiusArray := [...]float64{88, 32, 0, 38}
	for _, v := range celsiusArray {
		var celsiusItem = v
		fmt.Printf("%.2f celsius to fahrenheit %.2f\n", celsiusItem, celsiusItem*9/5+32)
	}
}
