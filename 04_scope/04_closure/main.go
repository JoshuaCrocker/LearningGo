package main

import "fmt"

// (1)
// var number int = 0

// func increment() int {
//      number++
//      return number
// }

// (2)
func wrapper() func() int {
	number := 0

	return func() int {
		number++
		return number
	}
}

func main() {
	x := 42

	fmt.Println(x)

	{ // This block is enclosed by the outer scope.
		fmt.Println(x) // 'x' is accessible

		y := "This is a string"

		fmt.Println(y)
	}

	// fmt.Println(y) // outside of scope of 'y'

	fmt.Println()
	fmt.Println()

	// Closure helps to limit the scope of variables used by multiple functions.

	// This is much better than declaring it at the package level (1), as long as our code doesn't get too complex

	number := 0

	increment := func() int { // anonymous function
		number++
		return number
	}

	// The wrapper method (2) is better still; it confines 'number' only to the scope that it is needed
	fmt.Println(increment()) // number = 1
	fmt.Println(increment()) // number = 2

	increment2 := wrapper()
	fmt.Println(increment2()) // 1
	fmt.Println(increment2()) // 2
}
