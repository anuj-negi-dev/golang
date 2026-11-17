package main

import "fmt"

func main() {
	var name string = "golang" // Explicit type
	var age = 12               // Type infernce
	isAdult := true            // shorthand to decalare variables
	fmt.Println(name, age, isAdult)

	// ----  Multiple variable declaration ----

	var x, y, z int = 10, 20, 30 // using explicit type

	// Grouped var declaration
	var (
		host   = "localhost"
		port   = 300
		active = true
	)

	a, b, name := 10, 20, "admin" // using shorthand

	fmt.Println(x, y, z, host, port, active, a, b, name)
}
