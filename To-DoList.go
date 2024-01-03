// This code allows the user to create a simple to-do 
//list by entering tasks. Typing "quit" exits the program, and the list of tasks is displayed

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	tasks := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a task (or 'quit' to exit): ")
		scanner.Scan()
		task := scanner.Text()

		if task == "quit" {
			break
		}

		tasks = append(tasks, task)
	}

	fmt.Println("Tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}
