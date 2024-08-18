package main

import (
	"fmt"

	"github.com/elliotchance/pie/pie"
	"github.com/ibiscum/Functional-Programming-in-Go/Chapter11/preGenerics/pkg"
)

func main() {
	out := pie.Ints{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}.
		Filter(func(i int) bool {
			return i%2 == 0
		}).
		Map(func(i int) int { return i * i })

	fmt.Printf("result: %v\n", out)

	MyDogs := []pkg.Dog{
		{
			Name: "Bucky",
			Age:  1,
		},
		{
			Name: "Keeno",
			Age:  15,
		},
		{
			Name: "Tala",
			Age:  16,
		},
		{
			Name: "Amigo",
			Age:  7,
		},
	}

	results := pkg.Dogs(MyDogs).
		Filter(func(d pkg.Dog) bool {
			return d.Age > 10
		}).SortUsing(func(a, b pkg.Dog) bool {
		return a.Age < b.Age
	})
	fmt.Printf("results: %v\n", results)
}
