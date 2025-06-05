package main
import "fmt"

func main() {

	var grade [5] int = [5]int{80,90,87,100,99}

	for index, element := range grade {
		fmt.Println(index, "=>", element)
	}
}