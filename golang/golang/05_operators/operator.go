package main

import "fmt"

func main() {
	// --- Arithmetic operators ------
	const (
		a = 10
		b = 3
	)
	fmt.Println(a + b) // add
	fmt.Println(a - b) // sub
	fmt.Println(a * b) // multiplication
	fmt.Println(a / b) // divison
	fmt.Println(a % b) // divison

	i := 10
	i++ // increment increase the value by one
	i-- // decrement decrease the value by one

	//	Comparison operators

	fmt.Println(a == b) // check if two value are equal or not
	fmt.Println(a > b)  // check if a is greater than b
	fmt.Println(a >= b) // check if a is greater or eqal than b
	fmt.Println(a < b)  // check if a is smaller than b
	fmt.Println(a <= b) // check if a is smaller or equal than b
	fmt.Println(a != b) // check if a is not equal than b

	// Logical operators
	/*
		 *
			* && - returns true if both conditons are true
			* || - returns true if any conditons is true
			* ! - !false becomes true and !true becomes false
	*/
}
