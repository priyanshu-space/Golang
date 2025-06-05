package main
import "fmt"

func main() {

	arr := [...]int{10,20,30,40,50,60}
	
	slice_1 := arr[1:3]
	fmt.Println(slice_1)
}