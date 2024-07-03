\\Q.concurrent prime number generator using goroutines and channels. The program should generate prime numbers up to a given limit and print them. Include comments to explain each part of the code.

package main

import (
	"fmt"
)

// generate sends the sequence 2, 3, 4, ... to channel 'ch'
func generate(ch chan<- int) {
	// Infinite loop to generate numbers starting from 2
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'
	}
}

// filter copies values from 'in' to 'out', removing those divisible by 'prime'
func filter(in <-chan int, out chan<- int, prime int) {
	// Infinite loop to receive values from 'in' channel
	for {
		i := <-in // Receive value from 'in'
		// If the value is not divisible by 'prime', send it to 'out' channel
		if i%prime != 0 {
			out <- i // Send 'i' to 'out' if not divisible by 'prime'
		}
	}
}

// sieve generates prime numbers up to a given limit using the Sieve of Eratosthenes
func sieve(limit int) {
	ch := make(chan int) // Create a new channel
	go generate(ch)      // Start generate() as a goroutine

	for {
		prime := <-ch // Receive the first value from 'ch', which is a prime number
		// If the prime number exceeds the limit, stop the function
		if prime > limit {
			return
		}
		fmt.Println(prime) // Print the prime number
		ch1 := make(chan int)  // Create a new channel for filtering
		go filter(ch, ch1, prime) // Start filter() as a goroutine to filter out multiples of 'prime'
		ch = ch1 // Update 'ch' to the new filtered channel
	}
}

func main() {
	var limit int
	// Prompt the user to enter the limit
	fmt.Print("Enter the limit: ")
	fmt.Scan(&limit) // Read the user input into 'limit'

	// Call the sieve function to generate and print primes up to the given limit
	sieve(limit)
}
