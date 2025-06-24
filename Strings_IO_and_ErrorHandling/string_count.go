package main
import (
	"fmt"
	"strings"
)

func main() {

	str1 := "learning standard library in go"
	str2 := "library"

	result := strings.Count(str1, str2) // it counts the str2 in str1 and Count is case sensitive.
	fmt.Println(result)
}