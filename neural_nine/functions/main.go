package function

import "fmt"

func main() {
	// Functions
	hello()

	fmt.Println(add(5, 4))
	fmt.Println(sub(5, 4))
	fmt.Println(sum(10, 20, 30, 40, 50))

	arr := []int{20, 40, 60, 80, 100}
	fmt.Println(sum(arr...))
}

func hello() {
	fmt.Println("Hello!!")
}

func add(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber
}

func sub(firstNumber, secondNumber int) int {
	return firstNumber - secondNumber
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
