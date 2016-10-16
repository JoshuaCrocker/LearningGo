package main

import "fmt"

const meaningOfLife = 42 // Doesn't need type

const ( // Multiple initialisation
	Pi       = 3.14
	Language = "Go"
)

const ( // They just increment
	A = iota // 0
	B = iota // 1
	C = iota // 2
)

const ( // Can also be written like this
	D = iota // 0
	E        // 1
	F        // 2
)

const ( // Throw away the 0
	_ = iota      // 0
	G = iota * 10 // 1 * 10 = 10
	H = iota * 10 // 2 * 10 = 20
)

const ( // Throw away the 0
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
)

func main() {

	fmt.Println(Pi)
	fmt.Println(Language)
	fmt.Println(meaningOfLife)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)

	fmt.Println(G)
	fmt.Println(H)

	fmt.Println()

	fmt.Println("Binary\t\t\t\tDecimal")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
}
