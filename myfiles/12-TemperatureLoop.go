package main

import "fmt"

func main() {
	for true {
		var temperature float64
		fmt.Println("Enter temperature in fahrenheit (^C to quit)")
		fmt.Scan(&temperature)
		// temperature := 88.0

		fmt.Printf("%.2f fahrenheit to celsius %0.2f\n", temperature, (temperature*9/5)+32)
	}

}
