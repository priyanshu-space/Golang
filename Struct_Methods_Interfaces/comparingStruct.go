package main
import "fmt"

type s1 struct {
	x int
}

type s2 struct {
	x int
}

// you can't compare two values of different struct type s1 and s2, it will throw compile time error. Values can be only compared within the same type of the struct.

func main() {

	c := s1{x:5}
	c1 := s1{x:6}
	c2 := s1{x:5}

	if c!=c1 {
		fmt.Println("c and c1 have different values")
	}
	if c==c2 {
		fmt.Println("c and c2 have the same values")
	}
}