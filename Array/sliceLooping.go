package main
import "fmt"

func main() {

	arr := []int{10,20,30,40,50}

	for index,value := range arr {
		fmt.Println(index, "=>", value)
	}

	//incase you don't want index but want to return value

	for _,value := range arr {
		fmt.Println(value)
	}
}