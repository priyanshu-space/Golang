package main
import (
	"fmt"
	"strings"
)

func main() {

	str1 := "learning standard library in go"
	str2 := "library in go"

	result := strings.Count(str1,str2) //it matches the strings of str2 in str1
	fmt.Println(result)
}