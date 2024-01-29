package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(100) + 1

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter the number")

		scanner.Scan()

		input := scanner.Text()

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid input")
			continue
		}
		if guess < secretNumber {
			fmt.Println("Your number is lower!")
		} else if guess > secretNumber {
			fmt.Println("Your number higher!")
		} else {
			fmt.Println("Correct! :happy-face:")
			break
		}
	}
}
