package main
import "fmt"

func f() (int, int) {
	return 42, 56
}

func main() {

	a,b := f()
	// v := f() --> will throw compile time error because it is returning two values
	// to hide the value you don't want, use '_' (blank identifier)

	_,v := f()

	fmt.Println(a,b)
	fmt.Println(v)
}