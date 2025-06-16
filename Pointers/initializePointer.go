package main
import "fmt"

func main() {

	s := "hello"

	var a *string = &s
	fmt.Println(a)
	var b = &s
	fmt.Println(b)
	c := &s
	fmt.Println(c)
}