package main

import (
	"fmt"
	"strconv"
)

func main() {
	age := 18
	// if else
	if age >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You are a minor")
	}

	score := 90
	// else if
	if score >= 60 {
		fmt.Println("A")
	} else if score >= 50 {
		fmt.Println("B")
	} else {
		fmt.Println("F")
	}

	// nested if

	isLoggedIn, isAdmin := true, false

	if isLoggedIn {
		if isAdmin {
			fmt.Println("Welcome, Admin!")
		} else {
			fmt.Println("Welcome, User!")
		}
	} else {
		fmt.Println("Please log in first.")
	}

	// if with Initialization Statement
	input := "42"
	if num, err := strconv.Atoi(input); err != nil {
		fmt.Printf("Failed to parse: %v\n", err)
	} else {
		fmt.Printf("Successfully parsed number: %d\n", num)
	}
}
