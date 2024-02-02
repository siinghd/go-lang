# Part 1 : Arrays

### Basic Declaration

- An array is declared with a fixed length and type.
- Syntax: `var arrayName [size]Type`

### Initializing Arrays

- Arrays can be initialized with specific values.
- Example: `arr := [3]int{1, 2, 3}`

### ... Array Literals

- Arrays can be initialized using array literals.
- Example: `arr := [...]int{1, 2, 3, 4}` (The compiler determines the length)

## Operations on Arrays

### Accessing Elements

- Access elements by their index: `value := arr[0]`

### Iterating over Arrays

- Use `for` or `for range` loops to iterate over arrays.

### Multidimensional Arrays

- Declare using multiple square brackets: `var matrix [3][3]int`

## Arrays and Functions

### Passing Arrays to Functions

- Arrays are passed by value, which means a copy is passed.
- To avoid copying, use pointers or slices.

### Returning Arrays from Functions

- Functions can return arrays, but this is often inefficient due to copying.

## Advanced Topics and Best Practices

### When to Use Arrays

- Use when the number of elements is known at compile time.
- Prefer slices for more flexible data structures.

### Arrays vs. Slices

- Slices are more common and versatile compared to arrays.
- Slices are reference types, while arrays are value types.

### Caveats and Pitfalls

- **Size as Part of the Type**: In Go, the size of the array is part of its type. This means `[3]int` and `[4]int` are different types.
- **Copy Behavior**: Assigning one array to another copies all the elements.
- **Large Arrays**: Large arrays can be costly to copy and can consume a lot of stack space.

## Examples

### Example 1: Basic Array Operations

```go
package main

import "fmt"

func main() {
    var arr [5]int
    arr[0] = 1
    arr[4] = 5
    fmt.Println("Array:", arr)
}
```

### Example 2: Iterating Over Arrays

```go
func printArray(arr [5]int) {
    for i, value := range arr {
        fmt.Println("Index:", i, "Value:", value)
    }
}

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    printArray(arr)
}
```

### Example 3: Multidimensional Array

```go
func main() {
    var matrix [2][2]int
    matrix[0][0] = 1
    matrix[0][1] = 2
    matrix[1][0] = 3
    matrix[1][1] = 4
    fmt.Println("Matrix:", matrix)
}
```

# Part 2: Slices 
 
## Introduction

This tutorial provides an in-depth look at slices in Go, a key data structure for flexible and efficient data management. We'll cover everything from basic usage to advanced techniques and common pitfalls.

## What Are Slices?

- Slices are a flexible, dynamic view into the elements of an array.
- They provide more functionality compared to arrays.

## Creating and Using Slices

### Declaring Slices

- Slices are declared without specifying a size.
- Example: `var slice []int`

### Initializing Slices

- Created with `make`: `slice := make([]int, length, capacity)`
- Slice literals: `slice := []int{1, 2, 3}`

### Slicing an Array

- Slices can be formed by "slicing" an array: `slice := array[start:end]`

## Operations on Slices

### Appending to Slices

- Use `append`: `slice = append(slice, element)`

### Copying Slices

- `copy` function: `copy(destSlice, srcSlice)`

### Reslicing

- Slices can be resliced: `newSlice := slice[start:end]`

## Slices in Functions

- Slices are reference types and can be modified within functions.

## Advanced Topics and Best Practices

### Capacity and Length

- Understanding the difference between length and capacity.
- Capacity grows automatically when appending beyond capacity.

### Memory Efficiency

- Overusing `append` can lead to performance issues due to frequent reallocations.
- Preallocate slices when the size is known in advance.

### Slices and Concurrency

- Slices are not thread-safe; use proper synchronization when using slices in concurrent environments.

### Caveats and Pitfalls

- **Modifying Underlying Array**: Modifying the elements of a slice modifies the underlying array, which can affect other slices.
- **Slice Zero Value**: The zero value of a slice is `nil`.
- **Slice Bounds**: Exceeding slice bounds can cause runtime panics.

## Examples

### Example 1: Basic Slice Operations

```go
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3}
    fmt.Println("Initial Slice:", slice)

    // Appending to slice
    slice = append(slice, 4)
    fmt.Println("Slice after append:", slice)
}
```

### Example 2: Modifying Slices in Functions

```go
func addToEachElement(slice []int, addValue int) {
    for i := range slice {
        slice[i] += addValue
    }
}

func main() {
    slice := []int{1, 2, 3}
    addToEachElement(slice, 1)
    fmt.Println("Modified Slice:", slice)
}
```

### Example 3: Slice Capacity and Reslicing

```go
func main() {
    slice := make([]int, 3, 5)
    fmt.Println("Length:", len(slice), "Capacity:", cap(slice))

    // Reslicing
    newSlice := slice[:cap(slice)]
    newSlice[3] = 4
    newSlice[4] = 5
    fmt.Println("Resliced Slice:", newSlice)
}
```
# Part 3: Structs

## What Are Structs?

- Structs are composite data types that group together variables (fields) under one name.
- They are used to create custom data types, making data handling more structured and clear.

## Defining and Instantiating Structs

### Basic Definition

- Syntax for defining a struct: `type StructName struct { Fields }`
- Example: `type Person struct { Name string; Age int }`

### Creating Instances

- Instantiating a struct: `person := Person{Name: "Alice", Age: 30}`

### Pointers to Structs

- Using pointers to work with structs efficiently and to modify their fields within functions.

## Struct Operations

### Accessing and Modifying Fields

- Accessing fields: `person.Name`
- Modifying fields: `person.Age = 31`

### Comparing Structs

- Structs are comparable if each field is comparable.
- Use `==` to compare struct instances.

### Nested Structs

- Structs can contain other structs, enabling complex data structures.

## Advanced Topics

### Anonymous (Embedded) Structs

- Embedding one struct within another without a field name.
- Promotes fields and methods to the embedding struct.

### Tags

- Struct field tags provide metadata for fields, commonly used with encoding/decoding libraries.

### Structs and Interfaces

- Implementing interfaces with structs.
- Using structs to provide concrete implementations of interface methods.

## Best Practices

- Use structs to model real-world data and entities.
- Keep struct design simple and cohesive.
- Use pointer receivers in methods to modify struct fields or for efficiency with large structs.

## Examples

### Example 1: Basic Struct Usage

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    person := Person{Name: "Alice", Age: 30}
    fmt.Println("Person:", person)
}
```

### Example 2: Nested Structs and Tags

```go
type Address struct {
    City, State string
}

type Person struct {
    Name    string
    Age     int
    Address Address
}

func main() {
    person := Person{
        Name: "Bob",
        Age:  35,
        Address: Address{
            City: "New York",
            State: "NY",
        },
    }
    fmt.Println("Person with Address:", person)
}
```

### Example 3: Structs with Methods

```go
type Rectangle struct {
    Width, Height float64
}

// Method to calculate area
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    fmt.Println("Area:", rect.Area())
}
```
# Part 4: Map

## What Are Maps?

- Maps are a collection of key-value pairs where each key is unique.
- They are similar to dictionaries or hash tables in other languages.

## Creating and Using Maps

### Declaring Maps

- Maps are declared using the `map` keyword.
- Syntax: `var mapName map[KeyType]ValueType`
- Example: `var employeeSalary map[string]int`

### Initializing Maps

- Maps can be initialized using the `make` function: `employeeSalary = make(map[string]int)`
- Map literals: `employeeSalary := map[string]string{"John": "Developer", "Emma": "Designer"}`

### Adding and Updating Elements

- Add or update elements using the key: `employeeSalary["Mike"] = 50000`

### Retrieving Elements

- Retrieve elements by their key: `salary := employeeSalary["Mike"]`

### Checking for Existence

- Use the "comma ok" idiom to check if a key exists: `salary, ok := employeeSalary["Mike"]`

## Advanced Map Operations

### Iterating over Maps

- Use a `for range` loop to iterate over maps.
- Iteration order over maps is not guaranteed.

### Deleting Elements

- Remove elements using the `delete` function: `delete(employeeSalary, "Mike")`

### Maps and Concurrency

- Maps are not safe for concurrent use without additional synchronization.
- Use sync.RWMutex for concurrent read/write access.

## Best Practices and Common Pitfalls

### Choosing Key Types

- Use comparable types for keys (int, string, etc.).
- Avoid using slices or other non-comparable types as keys.

### Nil Maps

- A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic.

### Memory Efficiency

- Preallocate maps if the size is approximately known.

## Examples

### Example 1: Basic Map Operations

```go
package main

import "fmt"

func main() {
    employeeSalary := make(map[string]int)
    employeeSalary["John"] = 30000
    employeeSalary["Emma"] = 45000

    fmt.Println("Employee Salaries:", employeeSalary)
}
```

### Example 2: Iterating and Deleting

```go
func main() {
    colors := map[string]string{"Red": "#FF0000", "Green": "#00FF00", "Blue": "#0000FF"}

    // Iterating over the map
    for color, hex := range colors {
        fmt.Println(color, ":", hex)
    }

    // Deleting an element
    delete(colors, "Green")

    fmt.Println("Updated Colors:", colors)
}
```

### Example 3: Maps with Struct Values

```go
type Employee struct {
    ID     int
    Salary int
}

func main() {
    employees := make(map[string]Employee)
    employees["John"] = Employee{ID: 1, Salary: 30000}
    employees["Emma"] = Employee{ID: 2, Salary: 45000}

    for name, info := range employees {
        fmt.Printf("%s: %+v\n", name, info)
    }
}
```


