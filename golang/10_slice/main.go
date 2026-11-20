package main

import "fmt"

func main() {
	// declaration
	var nums []int

	fmt.Print(nums == nil) // true
	fmt.Println(len(nums)) // 0
	fmt.Println(cap(nums)) // 0

	// Initializing Slices

	fruits := []string{"Apple", "Banana", "Cherry"}

	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	var colors []string = make([]string, 3, 5)
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	colors = append(colors, "Yellow", "Purple", "Orange")
	fmt.Println(colors)
	fmt.Println(len(colors))
	fmt.Println(cap(colors))
}
