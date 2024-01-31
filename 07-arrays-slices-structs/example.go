package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Part 1: Arrays
type FixedSizeData [5]int // An array type to store fixed size data

// Part 2: Slices
type DynamicData []int // A slice type to store dynamic data

// Part 3: Structs
type DataRecord struct {
	ID      int
	Fixed   FixedSizeData
	Dynamic DynamicData
}

// Generates random fixed size data
func generateFixedSizeData() FixedSizeData {
	var data FixedSizeData
	for i := range data {
		data[i] = rand.Intn(100) // Random number between 0 and 99
	}
	return data
}

// Generates random dynamic data
func generateDynamicData(size int) DynamicData {
	data := make(DynamicData, size)
	for i := range data {
		data[i] = rand.Intn(100) // Random number between 0 and 99
	}
	return data
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	// Creating an array of DataRecord structs
	records := make([]DataRecord, 3)

	for i := range records {
		records[i] = DataRecord{
			ID:      i + 1,
			Fixed:   generateFixedSizeData(),
			Dynamic: generateDynamicData(rand.Intn(5) + 1), // Dynamic size between 1 and 5
		}
	}

	// Displaying the data records
	for _, record := range records {
		fmt.Printf("Record %d:\nFixed Data: %v\nDynamic Data: %v\n\n", record.ID, record.Fixed, record.Dynamic)
	}
}
