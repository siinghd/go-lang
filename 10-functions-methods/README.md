# Part 1: Functions in Go

Functions are the building blocks of Go code. They allow you to encapsulate code logic and promote reusability. Understanding functions in Go is key to writing effective and organized code.

## Basic Function Declaration

### Syntax

- Functions are declared using the `func` keyword.
- Syntax: 
  ```go
  func functionName(parameterList) returnType {
      // function body
  }
  ```

### Parameters and Return Types

- Functions can have zero or more parameters.
- Functions can return one or more values.
- Named return values can be used for documentation purposes.

### Examples

```go
func greet(name string) string {
    return "Hello " + name
}

func add(a, b int) (result int) {
    result = a + b
    return // Named return
}
```

## Variadic Functions

- Variadic functions can accept a variable number of arguments.
- Defined by appending `...` before the type of the last parameter.

### Example

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```

## Anonymous Functions and Closures

### Anonymous Functions

- Functions without a name.
- Can be assigned to a variable or passed as an argument.

### Closures

- Anonymous functions that capture variables from the surrounding scope.
- Useful for creating function generators or maintaining state.

### Examples

```go
// Anonymous function
adder := func(a, b int) int {
    return a + b
}
fmt.Println(adder(5, 7))

// Closure
incrementor := func() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
inc := incrementor()
fmt.Println(inc()) // 1
fmt.Println(inc()) // 2
```

## Passing Functions as Arguments

- Functions can be passed as arguments to other functions.
- Enables higher-order functions.

### Example

```go
func applyOperation(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

result := applyOperation(3, 4, adder)
fmt.Println("Result:", result)
```

## Best Practices and Common Pitfalls

### Best Practices

- Keep functions focused on a single task.
- Use descriptive names for functions and parameters.
- Prefer shorter, readable functions over longer, complex ones.

### Pitfalls

- Avoid too many parameters in function signatures; consider using structs if the number of parameters is large.
- Beware of using named returns in long functions as they can reduce readability.


# Part 2: Methods in Go

Methods in Go attach a function to a specific type, allowing you to define behavior associated with that type. Understanding methods is crucial for structuring and organizing code in Go, especially when working with custom types and interfaces.

## Defining Methods

### Syntax

- Methods are functions with a special receiver argument.
- Syntax: 
  ```go
  func (receiver ReceiverType) MethodName(params) returnType {
      // method body
  }
  ```

### Value vs. Pointer Receivers

- A method can have either a value receiver or a pointer receiver.
- Value receivers operate on a copy of the original type.
- Pointer receivers can modify the original value.

### Example

```go
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
```

## Embedding and Composition

- Go supports composition over inheritance, achieved through embedding.
- Embedding a type allows access to its methods and fields.

### Example

```go
type Shape struct {
    X, Y float64
}

type Circle struct {
    Shape
    Radius float64
}

// Method on the embedded Shape type
func (s *Shape) Move(dx, dy float64) {
    s.X += dx
    s.Y += dy
}
```

## Best Practices and Common Pitfalls

### Best Practices

- Use pointer receivers to avoid copying on method calls or to allow methods to modify the receiving struct.
- Keep method receivers consistent for a given type (either all pointer receivers or all value receivers).

### Pitfalls

- Avoid embedding types with overlapping method sets, as it can lead to ambiguous calls.
- Be cautious of value receivers for large structs as they can be inefficient due to copying.

### Advanced Example: Composition and Method Overriding

```go
type Rectangle struct {
    Shape
    Width, Height float64
}

// Method overriding
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 4, Height: 3}
    rect.Move(1, 1)  // Accessing Move method from embedded Shape
    fmt.Println("Area of rectangle:", rect.Area())
}
```

In this example, `Rectangle` embeds `Shape` and overrides the `Area` method. The `Move` method from `Shape` is accessible from a `Rectangle` instance.

