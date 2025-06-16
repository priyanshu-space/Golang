package main
import "fmt"

func main() {

	x := func(l int, b int) (int) {
		return l*b
	}

	// x := func(l int, b int) (int) {
	// 	return l*b
	// }(20,30)
	//fmt.Println(x)

	fmt.Println(x(20,30))
	fmt.Printf("%T \n", x)
}