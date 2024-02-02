package main

import "fmt"

func main() {
	// Initializing a pointer to an int
	ptr := new(int)
	fmt.Println("Pointer:", ptr)           // Outputs the memory address
	fmt.Println("Value at pointer:", *ptr) // Outputs the value (0)

	// Assigning a value to the location pointed by the pointer
	*ptr = 100
	fmt.Println("New value at pointer:", *ptr)

	slice := make([]int, 3) // Length and capacity are 3
	slice[0] = 10
	slice[1] = 20
	slice[2] = 30
	fmt.Println(slice)

	employeeSalary := make(map[string]int)
	employeeSalary["John"] = 40000
	employeeSalary["Jane"] = 45000
	fmt.Println(employeeSalary)

	ch := make(chan int, 2) // Buffered channel of size 2
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
