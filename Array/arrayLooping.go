package main
import "fmt"

func main() {

	//var grades [5] int = [5]int{90,80,85,99,100}

	grades:=[...]int{90,80,85,99,100}

	for i:=0; i<len(grades); i++ {
		fmt.Println(grades[i])
	}
}