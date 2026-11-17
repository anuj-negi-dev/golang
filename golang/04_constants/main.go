package main

import "fmt"

func main() {
	const name = "golang" // unsinged const

	const uname string = "javascript" // signed const

	// const grouping
	const (
		port = 5000
		host = "localhost"
	)

	const (
		a = 100 // A is 100
		b       // B automatically 100 ban jayega (repeats A's value)
		c = 200 // C is 200
		d       // D automatically 200 ban jayega (repeats C's value)
	)

	fmt.Println(a, b, c, d)

	const (
		A = iota // 0
		B
	)

	const (
		admin = iota
		manager
		user
	)
}
