package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error:", r)
		}
	}()

	fmt.Print("Enter the number of hours until the meeting: ")
	var input string
	fmt.Scanln(&input)

	hours, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error: Invalid input. Please enter a number.")
		return
	}

	if hours < 0 {
		panic("Error: Number of hours cannot be negative.")
	}

	meetingTime := time.Now().Add(time.Duration(hours) * time.Hour)
	fmt.Println("The meeting is scheduled for:", meetingTime.Format("2006-01-02 15:04:05"))
}
