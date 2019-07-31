package main

import (
	"fmt"

	"github.com/BillKoder/training/HelloWorld/pkg/mathematics"
)

func main() {
	fmt.Println("Hello World")

	fmt.Println((mathematics.Sum([]int{1, 2, 3, 4})))
	fmt.Println(mathematics.Multiply([]int{2, 2}))
	fmt.Println(mathematics.Difference(5, 3))
}
