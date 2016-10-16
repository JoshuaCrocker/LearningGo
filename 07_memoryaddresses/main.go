package main

import "fmt"

const cstMetersToYards float64 = 1.09361

func main() {
	a := 43

	fmt.Println("a: ", a)
	fmt.Println("a's address:", &a)
	fmt.Println()

	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * cstMetersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
}
