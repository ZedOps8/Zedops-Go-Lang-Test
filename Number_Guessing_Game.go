This code creates a simple number guessing game where the player tries to guess a randomly generated number.


package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(100)
	var guess int
	attempts := 0

	fmt.Println("Guess the number between 0 and 99")

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)
		attempts++

		if guess < secretNumber {
			fmt.Println("Try a higher number.")
		} else if guess > secretNumber {
			fmt.Println("Try a lower number.")
		} else {
			fmt.Printf("Congratulations! You guessed it in %d attempts.\n", attempts)
			break
		}
	}
}
