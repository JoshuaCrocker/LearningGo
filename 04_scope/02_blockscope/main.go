package main

import "fmt"

func main() {
	// 'a' doesn't exist here

	a := "Hello, world" // Scope 'a' begins here
	x := 42
	fmt.Println(x)
	foo()
	fmt.Println(x)

	fmt.Println(a)
	fmt.Println(name)
} // 'a' ends here

func foo() {
	x := 21 // Different instance of 'x' to in the main() block.
	fmt.Println(x)
}

var name string = "Joshua Crocker" // Will be available anywhere within this scope - order doesn't matter here
