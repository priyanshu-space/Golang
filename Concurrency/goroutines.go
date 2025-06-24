package main

import (
	"fmt"
	"time"
)

func calculateSquare(i int) {

	fmt.Println(i*i)

}

func main() {

	start := time.Now()
	
	for i := 0; i<10; i++ {

		go calculateSquare(i) // go routines are underministic
	}

	elpased := time.Since(start)
	fmt.Println("function took ", elpased)

}