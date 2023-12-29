//This code generates and prints a random number between 0 and 99 using the math/rand package.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)
	fmt.Printf("Random number: %d\n", randomNumber)
}
