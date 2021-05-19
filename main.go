package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// declaring variable
	var random_number int
	var guess int
	min := 1
	max := 6
	random_number = rand.Intn(max-min) + min

	// prompts user input
	fmt.Println("Pick random number")
	fmt.Println("Guess(1-6):")
	fmt.Scanln(&guess)

	// check if number is correct
	if guess == random_number {
		fmt.Println("CORRECT!!! RANDOM NUMBER IS ", random_number)
	} else {
		fmt.Println("INCORRECT")
	}
}
