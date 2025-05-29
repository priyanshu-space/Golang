package main
import "fmt"

func main() {

	var name string = "Hello"
	var user string = "Priyanshu"

	fmt.Print(name, " ", user, "\n")

// printf - format specifier
// %v - formats the value in default format
// %d - decimal format specifier
// %T - data type of the value
// %q - quoted characters/string
// %s - plain string
// %f - floating numbers
// %.2f - floating numbers up to 2 decimal.

    fmt.Printf("%v Nice to see you", user)
}