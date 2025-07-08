package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile("demo-file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err!=nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//append a string to a file.

	if _, err := file.WriteString("Hope you had a good day"); err!=nil {
		fmt.Println(err)
	}
}