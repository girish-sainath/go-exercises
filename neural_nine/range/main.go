package main

import "fmt"

func main() {
	// Range

	numbers := []int{12, 34, 45, 67, 77}
	for idx, number := range numbers {
		fmt.Printf("Value at index %d is %d\n", idx, number)
	}

	numbers = []int{22, 44, 55, 77, 87}
	for _, number := range numbers {
		fmt.Printf("Value is %d\n", number)
	}

	myMap := map[string]int{"A": 1, "B": 2, "C": 3}
	for key, value := range myMap {
		fmt.Printf("Value for key %s is %d\n", key, value)
	}
	for key := range myMap {
		fmt.Printf("Key %s\n", key)
	}

	for _, value := range myMap {
		fmt.Printf("Value %d\n", value)
	}
}