package main

import (
	"fmt"

	"github.com/ibiscum/Functional-Programming-in-Go/Chapter7/Performance/pkg"
)

func main() {
	fmt.Println(pkg.RecursiveFact(10))
	//debug.SetMaxStack(100)
	//stackOverflowExample(0)
}
