package main

import (
	"fmt"
	"mycalculator/mathutils"
)

func main() {

	sum := mathutils.Add(10,5)
	product := mathutils.Multiply(10, 5)
	fmt.Println(sum)
	fmt.Println(product)
}