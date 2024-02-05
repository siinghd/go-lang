package main

import (
	"fmt"
	"math"
)

// Part 1: Functions

// Basic function that adds two integers
func add(a, b int) int {
	return a + b
}

// Variadic function that sums an arbitrary number of integers
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Anonymous function assigned to a variable
var double = func(n int) int {
	return n * 2
}

// Part 2: Methods

// Struct for demonstrating methods
type Circle struct {
	Radius float64
}

// Method with a value receiver
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Method with a pointer receiver
func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

func main() {
	// Demonstrating functions
	fmt.Println("Add 3 and 5:", add(3, 5))
	fmt.Println("Sum of 2, 4, and 6:", sum(2, 4, 6))
	fmt.Println("Double 4:", double(4))

	// Demonstrating methods
	c := Circle{Radius: 5}
	fmt.Println("Circle area:", c.Area())
	c.Scale(2)
	fmt.Println("Scaled circle radius:", c.Radius)
}
