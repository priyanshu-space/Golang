package main
import (
	"fmt"
	"strings"
)

func main() {

	str1 := "learning standard library in go"
	str2 := "library"

	result := strings.ReplaceAll(str1, str2, "library in python") // it replaces the str2 with the new and then replaces the occurance in str1 .
	fmt.Println(result)
	fmt.Println(str2) //no effect in original
}