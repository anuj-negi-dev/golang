package main

import "fmt"

func main() {
	// ---------- Integer Data Type --------

	// ---- signed integer: can store both negative and postive value ------
	/*
		 		formul 2^n-1 to 2^n-1 -1
			    int8	1 byte (8 bits)	-128 to 127
			 	int16	2 bytes (16 bits)	-32,768 to 32,767
				int32	4 bytes (32 bits)	-2,147,483,648 to 2,147,483,647
				int64	8 bytes (64 bits)	-9.22 × 10¹⁸ to 9.22 × 10¹⁸
				int	Platform-dependent	32-bit on 32-bit systems, 64-bit on 64-bit systems
	*/

	// ---- unsigned integer: only store postive value from 0 to 2^n (n is bits)------

	/*
		 *
			uint8	1 byte (8 bits)		0 to 255
			uint16	2 bytes (16 bits)	0 to 65,535
			uint32	4 bytes (32 bits)	0 to 4,294,967,295
			uint64	8 bytes (64 bits)	0 to 2^64 - 1
			uint	Platform-dependent	32-bit on 32-bit systems, 64-bit on 64-bit systems
	*/

	var age int = 12

	// ------- float is used to store decimal value ---------

	/*
		 *
			* The precision of float32 starts breaking down after about 7 significant digits (beyond this, it starts rounding off or producing inaccuracies).

			The precision of float64 goes up to about 15 to 17 digits, which makes it much more accurate.

			By default is take flafloat64
	*/

	pi := 3.14

	// ------- boolean -----------

	isAdult := true

	// ------- string ----------

	role := "admin"

	// ----- byte ----------

	var ch byte = 'A'

	fmt.Println(age, pi, isAdult, role, ch)

	x := int(ch)

	fmt.Print(x)
}
