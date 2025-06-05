package main
import "fmt"

func main() {

	var arr [5] int = [5]int{10,20,30,40,50}
	slice := arr[1:3] // slice have its own indexing.
	fmt.Println(slice)
	fmt.Println(arr)

	slice[0] = 90 // It also affects the underlying array.
	fmt.Println("After modification")
	fmt.Println(slice)
	fmt.Println(arr)
}