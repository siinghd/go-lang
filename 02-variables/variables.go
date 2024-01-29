package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// Integer Types
	// 'int' size is dependent on the system architecture (32-bit or 64-bit)
	var age int = 30
	var year uint = 2024

	// Floating Point Types
	// 'float32' and 'float64' are for floating-point numbers with 32 and 64 bits of precision respectively
	var temperature float32 = 26.5
	var distance float64 = 384400.0 // Distance to the Moon in km

	// Complex Numbers
	// 'complex128' is used for complex numbers with double precision
	var complexNumber complex128 = cmplx.Sqrt(-5 + 12i)

	// Boolean Type
	// 'bool' represents a boolean and is either true or false
	var isSunny bool = true

	// String Type
	// 'string' is a sequence of characters with immutable nature
	var greeting string = "Hello, Go!"

	// Composite Types
	// Array
	// Arrays have a fixed size defined at compile time
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

	// Slice
	// Slices are similar to arrays but their size is dynamic, they are a reference type
	var colors []string = []string{"Red", "Green", "Blue"}

	// Struct
	// Structs are custom types that group together variables of different types
	type Location struct {
		Latitude, Longitude float64
	}
	var officeLocation Location = Location{Latitude: 27.2046, Longitude: 77.4977}

	// Map
	// Maps are key-value pairs, they can grow dynamically and are reference types
	var capitals map[string]string = map[string]string{
		"France": "Paris",
		"Japan":  "Tokyo",
		"India":  "New Delhi",
	}

	// Pointer
	// Pointers hold the memory address of a variable, allows for indirect manipulation of variable
	var pointerToAge *int = &age

	// Rune Type
	// 'rune' represents a Unicode character and is an alias for 'int32'
	var heart rune = '❤'

	// Print values of the variables
	fmt.Println("Integer:", age)
	fmt.Println("Unsigned Integer:", year)
	fmt.Println("Float32:", temperature)
	fmt.Println("Float64:", distance)
	fmt.Println("Complex Number:", complexNumber)
	fmt.Println("Boolean:", isSunny)
	fmt.Println("String:", greeting)
	fmt.Println("Array:", numbers, "\n   Last Value:", numbers[len(numbers)-1])
	fmt.Println("Slice:", colors)
	fmt.Println("Struct:", officeLocation, "\n   Single value:", officeLocation.Latitude)
	fmt.Println("Map:", capitals, "\n   Access the value:", capitals["France"])
	fmt.Println("Pointer:", pointerToAge, "Pointing to value:", *pointerToAge)
	fmt.Println("Rune (Unicode Character):", string(heart), "\n   Unicode value:", heart)
}
