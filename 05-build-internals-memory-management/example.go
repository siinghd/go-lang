package main

import (
	"fmt"
	"strconv"
	"time"
)

// processData processes data received on the input channel and sends results on the output channel
func processData(input <-chan string, output chan<- int, errors chan<- error) {
	for value := range input {
		number, err := strconv.Atoi(value)
		if err != nil {
			errors <- fmt.Errorf("invalid input: %s", value)
			continue
		}
		output <- number * number // Squaring the number
	}
}

func main() {
	inputData := []string{"4", "9", "a", "16"}
	inputChan := make(chan string)
	outputChan := make(chan int)
	errorChan := make(chan error)

	go processData(inputChan, outputChan, errorChan)

	// Sending data to be processed
	go func() {
		for _, val := range inputData {
			inputChan <- val
		}
		close(inputChan)
	}()

	// Receiving and handling output
	for i := 0; i < len(inputData); i++ {
		select {
		case result := <-outputChan:
			fmt.Println("Processed result:", result)
		case err := <-errorChan:
			fmt.Println("Error:", err)
		}
	}

	// Using time.Sleep to prevent the program from exiting immediately
	time.Sleep(1 * time.Second)
}
