package main

import "fmt"

func main() {

	var fruits [5] string = [5] string{"apple", "mango","orange","grapes","papaya"}
	fmt.Println(fruits[4])
	fruits[1]="banana"
	fmt.Println(fruits)
}