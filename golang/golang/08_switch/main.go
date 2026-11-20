package main

import (
	"fmt"
)

func main() {
	// switch statement
	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("First day of week")
	case "Tuesday":
		fmt.Println("Second day of week")
	default:
		fmt.Println("Midweek or weekend")
	}

	// Multiple Cases in switch
	day = "Saturday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	// switch statement without expression
	score := 85

	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}

	// fallthrough

	num := 1

	switch num {
	case 1:
		fmt.Println("One")
		fallthrough // Forces execution into case 2 regardless of its value
	case 2:
		fmt.Println("Two (executed via fallthrough)")
	case 3:
		fmt.Println("Three")
	}

}
