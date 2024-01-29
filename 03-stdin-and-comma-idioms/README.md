# Handling of Standard Input and Comma Idioms in Go

## Introduction

This tutorial is about standard input (`stdin`) in Go and provides production-ready examples of the "comma ok" and "comma error" idioms. We'll also review a "guess the number" game to apply these concepts.

## Reading Standard Input in Go

Reading input from the user in Go can be accomplished using functions like `fmt.Scanln()` and `bufio.Scanner`.

## Comma Ok Idiom - Production Example

The "comma ok" idiom is essential for safely accessing values in maps and asserting interface types. Here's a production-ready example:

### Example: Accessing a Map Value

```go
func getUserRole(roles map[string]string, username string) string {
    if role, ok := roles[username]; ok {
        return role // Username exists
    }
    return "guest" // Default role
}
```

In this example, `roles` is a map with usernames as keys. The "comma ok" idiom checks if `username` is present in the map. It's a safe way to access map values without causing a panic.

## Comma Error Idiom - Production Example

The "comma error" idiom is widely used for error handling. Here's how it might appear in a production setting:

### Example: Reading a File

```go
func readFileContent(path string) (string, error) {
    content, err := ioutil.ReadFile(path)
    if err != nil {
        return "", err // Handle the error
    }
    return string(content), nil
}
```

This function reads a file and returns its content. The "comma error" idiom is used to check if the `ReadFile` operation was successful.

## Custom Comma Ok Idiom

The "comma ok" idiom can be extended to custom types, especially when dealing with type assertions and checking the presence of elements in complex data structures.

### Example: Custom Type Assertion

```go
type Shape interface {
    Draw() string
}

type Circle struct {}

func (c Circle) Draw() string {
    return "Drawing a circle"
}

// CustomCommaOk checks if the Shape is a Circle and returns a boolean accordingly
func CustomCommaOk(shape Shape) (Circle, bool) {
    circle, ok := shape.(Circle)
    return circle, ok
}
```

In this custom example, we're asserting whether a `Shape` interface is of type `Circle`. The function `CustomCommaOk` returns both the asserted type and a boolean indicating success.

## Custom Comma Error Idiom

Similarly, the "comma error" idiom can be customized for functions that need to return a value and a custom error.

### Example: Custom Error Handling

```go
type Result struct {
    Value string
}

// CustomCommaError simulates a function that returns a result and a custom error
func CustomCommaError(condition bool) (Result, error) {
    if condition {
        return Result{Value: "Success"}, nil
    } else {
        return Result{}, fmt.Errorf("custom error occurred")
    }
}
```

This function returns a `Result` and a custom error based on a condition. It's a pattern commonly used in Go to simultaneously handle return values and errors.


