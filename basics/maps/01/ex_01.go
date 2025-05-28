package main

import (
	"fmt"
)



func main() {
	// var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
	// b := map[KeyType]ValueType{key1:value1, key2:value2,...}

	var m = map[int]string{123: "Chris", 456: "John", 789: "Gabi"}
	fmt.Println("map:", m)

	m[123] = "Chris Santos"
	fmt.Println("map after update:", m)

	// Adding a new key-value pair
	m[101] = "Alice"
	fmt.Println("map after adding new key-value pair:", m)

	m[456] = "Rafael"
	fmt.Println("map after updating:", m)

	// Deleting a key-value pair
	delete(m, 456)
	fmt.Println("map after deleting key 456:", m)

	// Checking if a key exists
	if value, exists := m[789]; exists {
		fmt.Println("Key 789 exists with value:", value)
	} else {
		fmt.Println("Key 789 does not exist")
	}

	// Iterating over a map
	for key, value := range m {
		fmt.Printf("Key: %d - Value: %s\n", key, value)
	}

	// Length of the map
	fmt.Println("Length of the map:", len(m))

	// Creating an empty map
	emptyMap := make(map[string]string)
	fmt.Println("Empty map:", emptyMap)

	// Adding elements to the empty map
	emptyMap["UUI12"] = "One"
	emptyMap["UUI13"] = "Two"
	fmt.Println("Empty map after adding elements:", emptyMap)

	// Example of a map with string keys and int values
	stringIntMap := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 8,}

	for key, value := range stringIntMap {
		if value < 5 {
			fmt.Println("Value less than 5:", key)
		}
		fmt.Println("Value:", key)
	}
}