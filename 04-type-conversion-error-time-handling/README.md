# Type Conversion, Error Handling, Time Management, and Panic/Recover

## Introduction

This comprehensive tutorial covers type conversion, error handling, time management, and the use of panic/recover in Go, with practical examples for each.

## Type Conversion in Go

### Example: Integer to Float Conversion

```go
var myInt int = 10
var myFloat float64 = float64(myInt)
fmt.Println(myFloat) // Output: 10.0
```

- **Caveat**: When converting from a larger to a smaller type, or from floating-point to integer, there's a risk of overflow or data loss.

## Comprehensive Error Handling in Go

### Example: Error Checking

```go
file, err := os.Open("nonexistent.txt")
if err != nil {
    log.Fatalf("Error opening file: %v", err)
}
defer file.Close()
```

- **Custom Error**: Creating custom errors using `fmt.Errorf` or by implementing the `Error()` method.

## Managing Time in Go

### Example: Formatting Time

```go
currentTime := time.Now()
formattedTime := currentTime.Format("2006-01-02 15:04:05")
fmt.Println("Formatted Time:", formattedTime)
```

- **Time Arithmetic**: Adding and subtracting durations from `time.Time` objects.

## Panic and Recover in Go

### Example: Using Panic and Recover

```go
func riskyFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()
    panic("something bad happened")
}

riskyFunction()
```

- **Best Practice**: Use `panic`/`recover` carefulluy, as a last resort, not as a substitute for regular error handling.


