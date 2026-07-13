package main

import "fmt"

func main() {
	// Closures
	nextEvenNumber := evenNumbers()
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
}

func evenNumbers() func() int {
	i := 0
	return func() int {
		i += 2
		return i
	}
}

