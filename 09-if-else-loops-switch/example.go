package main

import "fmt"

func main() {
	// Part 1: Conditional Statements - `if`, `else if`, `else`
	age := 25
	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 && age < 60 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior.")
	}

	// Part 2: Loops - `for` and `for range`
	// Basic for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Loop iteration:", i)
	}

	// for range loop
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
	}

	// Part 3: The `switch` statement
	// Basic switch
	dayOfWeek := 3
	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Other day")
	}

	// Switch with a short statement
	switch today := dayOfWeek; today {
	case 1, 2, 3, 4, 5:
		fmt.Println("It's a weekday")
	case 6, 7:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("Invalid day")
	}

	// Type switch
	var i interface{} = 10.0
	switch v := i.(type) {
	case int:
		fmt.Println("i is an int:", v)
	case float64:
		fmt.Println("i is a float64:", v)
	case string:
		fmt.Println("i is a string:", v)
	default:
		fmt.Println("i is another type")
	}
}
