package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Printing and formatting

	fmt.Println("Hello World!")
	var name string = "World"
	name = "Girish"
	var age int8 = 10
	fmt.Println(name)
	fmt.Println(reflect.TypeOf(age))

	var gravitionalConstant float64 = 10
	gravitionalConstant = 9.8
	fmt.Println(gravitionalConstant)
	fmt.Println(reflect.TypeOf(gravitionalConstant))

	const PI = 3.141592653589793
	fmt.Println(PI)

	hasMoney := false
	fmt.Println(hasMoney)

	fmt.Println("Hello " + name + ", how are you doing? Are you " + fmt.Sprint(age) + " years old?")
	fmt.Printf("Hello %s, how are you doing? Are you %d years old?\n", name, age)
}
