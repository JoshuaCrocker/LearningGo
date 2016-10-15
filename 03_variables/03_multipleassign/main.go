package main

import "fmt"

func main() {
	var a, b int                       // Multiple declaration
	var c, d string = "Hello", "World" // Multiple Initialisation

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n\n", d)

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
}
