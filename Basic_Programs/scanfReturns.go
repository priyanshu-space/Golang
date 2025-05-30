package main
import "fmt"

func main() {

	var a string
	var b int
	fmt.Print("Enter a string and a number: ")

	//var count int
	//var err error

    count, err := fmt.Scanf("%s %d", &a, &b)
	fmt.Println("count:",count)
	fmt.Println("error: ",err)
	fmt.Println("String is",a)
	fmt.Println("number is",b)

}