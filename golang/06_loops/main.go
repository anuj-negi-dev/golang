package main

import "fmt"

func main() {
	// classic for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// break and continue

	for i := 1; i <= 10; i++ {
		if i == 2 {
			continue
		}
		if i == 5 {
			break
		}
	}

	// infinite loop
	for {
		fmt.Println("Infinite loop")
	}
}
