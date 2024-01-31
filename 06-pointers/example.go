package main

import "fmt"

// Person struct for the struct example
type Person struct {
	Name string
	Age  int
}

// incrementAge increments the age of a Person
func incrementAge(p *Person) {
	p.Age++
}

// modifyArray modifies the specified element of an array
func modifyArray(arr *[3]int) {
	(*arr)[1] = 10 // Modify the second element of the array
}

// modifySlice modifies the first element of a slice
func modifySlice(s []int) {
	s[0] = 99
}

func main() {
	// Example: Basic Pointer Operations
	var a int = 20
	var p *int = &a
	fmt.Println("Value of a:", a)          // Value of a
	fmt.Println("Address of a:", p)        // Address of a
	fmt.Println("Value at address p:", *p) // Value at address p

	// Example: Modifying Array Elements
	arr := [3]int{1, 2, 3}
	modifyArray(&arr)
	fmt.Println("Modified array:", arr)

	// Example: Working with Structs
	person := Person{Name: "Alice", Age: 30}
	incrementAge(&person)
	fmt.Println("Modified struct:", person)

	// Example: Edge Case with Slices
	mySlice := []int{1, 2, 3}
	modifySlice(mySlice)
	fmt.Println("Modified slice:", mySlice)
}
