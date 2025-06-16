package main
import "fmt"

func modify(s *string){

	*s = "World"
}

func main() {

	a := "hello"
	fmt.Println(a)
	modify(&a)
	fmt.Println(a)
}