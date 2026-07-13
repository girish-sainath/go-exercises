package main

import "fmt"

func main() {

	// Arrays

	var arr1 [5]int

	arr1[1] = 20
	arr1[4] = 10

	fmt.Println(arr1)
	fmt.Println(arr1[1])

	arr2 := [5]int{10, 20, 30, 40, 50}

	fmt.Println(arr2)

	fmt.Println(len(arr2))

	var arr3 [5]string

	fmt.Println(arr3)

	var arr4 [5]bool

	fmt.Println(arr4)

	var arr2D [4][5]int

	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			arr2D[i][j] = i * j
		}
	}

	fmt.Println(arr2D)

	// Slices -> Lists

	s1 := make([]string, 5)
	s2 := make([]int, 5)

	s1[1] = "hola"
	s2[2] = 5

	s2 = append(s2, 56, 67, 45)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(len(s2))

	cp1 := make([]int, len(s2))
	copy(cp1, s2)

	cp1[1] = 30

	fmt.Println(cp1)
	fmt.Println(s2)

	fmt.Println(cp1[:4])

	// Maps -> Dictionaries

	m1 := make(map[string]int)

	m1["zero"] = 0
	m1["one"] = 1
	m1["two"] = 2
	m1["three"] = 3

	fmt.Println(m1)
	fmt.Println(m1["one"])
	fmt.Println(m1["four"])

	value, present := m1["four"]
	fmt.Println(value)
	fmt.Println(present)

	value, present = m1["zero"]
	fmt.Println(value)
	fmt.Println(present)

	value, present = m1["three"]
	fmt.Println(value)
	fmt.Println(present)

	delete(m1, "three")
	delete(m1, "four")

	fmt.Println(m1)
}
