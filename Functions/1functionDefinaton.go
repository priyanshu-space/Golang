package main
import "fmt"

func printGreeting(str string) {
	fmt.Print("hello, ", str )
}

func main() {
	printGreeting("Joe")
}