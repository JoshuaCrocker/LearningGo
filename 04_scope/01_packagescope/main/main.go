package main

import (
	"fmt"

	"github.com/JoshuaCrocker/LearningGo/04_scope/01_packagescope/vis"
)

var x int = 42

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)

	// Can't access vis.forename

	vis.PrintName()
}
