package main

import "fmt"

func main() {
	// Logical operators

	trueValue1 := true
	falseValue1 := false

	trueValue2 := true
	falseValue2 := false

	fmt.Println(trueValue1 && trueValue2)
	fmt.Println(trueValue1 && falseValue2)
	fmt.Println(trueValue1 || trueValue2)
	fmt.Println(trueValue1 || falseValue2)
	fmt.Println(!falseValue1)

	// Arithmetic operators

	fmt.Println(10 + 2)
	fmt.Println(10 - 2)
	fmt.Println(10 * 2)
	fmt.Println(10 / 2)
	fmt.Println(10 % 3)

	var a int = 10
	a += 5
	fmt.Println(a)

	var b int = 20
	b -= 5
	fmt.Println(b)

	var c int = 5
	c *= 2
	fmt.Println(c)

	var d int = 20
	d /= 4
	fmt.Println(d)

	var e int = 10
	e %= 3
	fmt.Println(e)

	var myInput int = 456

	// 	fmt.Println("Enter a number: ")
	// 	fmt.Scan(&myInput)
	fmt.Printf("Your number was %d\n", myInput)

	// Conditionals

	if myInput < 100 {
		fmt.Println("Less than a hundred")
	} else if myInput < 200 {
		fmt.Println("Less than a two hundred")
	} else {
		fmt.Println("Too large")
	}

	// Switch statements

	switch myInput {
	case 10:
		fmt.Println("10 was your input")
	case 20:
		fmt.Println("20 was your input")
	default:
		fmt.Println("Something else")
	}

	// Loops - for loops

	fmt.Println("While like for loop")
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	fmt.Println("Classic for loop")
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}

	fmt.Println("Range based for loop")
	for x := range 10 {
		fmt.Println(x)
	}
}
