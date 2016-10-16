package main

import "fmt"

func max(x int) int {
	return 42 + x
}

func main() {
	max := max(7)
	fmt.Println(max) // max is now the result, not the function
	// cannot call function max() from here anymore...
}

// This is very bad coding practice; don't do this.
