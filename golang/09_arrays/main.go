package main

import "fmt"

func main() {
	/*
		Fixed Size: Once an array is defined, its size cannot be changed.

		Type Safety: The length of the array is part of its type. For example, [3]int and [5]int are treated as entirely different types in Go.

		Value Semantics: Arrays in Go are values, meaning assigning an array to a new variable or passing it to a function copies the entire underlying data.
	*/

	// Declares an array of 5 integers (all initialized to 0)
	var arr [5]int
	fmt.Println(arr) // [0, 0, 0, 0, 0]

	var fruits = [3]string{"Apple", "Banana", "Mango"} // Standard Literal
	nums := [2]int{2, 4}                               // shorthand
	score := [...]int{3, 4, 5, 0, 1} // Size is automatically determined as 5

	statusCodes := [3]int{0: 200} // index-Based Initialization

	colors := [3]string{"Red", "Green", "Blue"}

	fmt.Println(colors[0]) // accessing array elements

	colors[1] = "violet" // update array element

	fmt.Println("length of the arr is", len(colors)) // get the length of arr

	// Classic for loop

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	// for ... range loop
	for index, color := range colors {
		fmt.Printf("color @ %v is %v\n", index, color)
	}

	// multidimensional arrays

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	// access [row][col]
	fmt.Println(matrix[1][0])

	fmt.Println(colors, fruits, nums, score, statusCodes)
}
