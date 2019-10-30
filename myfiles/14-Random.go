package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	min := 10
	max := 30
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(max-min) + min
	fmt.Println()

	guess := 0
	for guess != number {
		fmt.Printf("Guess the random number? (min = %d max = %d)\n", min, max)
		fmt.Scan(&guess)
	}

	fmt.Println("You guessed", guess, "the random number was", number)
}
