package main
import "fmt"

// func operation(a int, b int) (int, int) {

// 	sum := a+b
// 	diff := a-b
// 	return sum, diff
// }

func operation(a int, b int) (sum int, diff int) { //declaring sum and diff type in function signature
	sum = a+b
	diff = a-b
	return
}

func main() {

	sum, diff := operation(20,10)
	fmt.Println(sum, diff)

}