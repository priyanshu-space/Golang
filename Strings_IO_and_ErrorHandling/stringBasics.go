package main
import (
	"fmt"
	//"strings"
)

func main() {

	name := "Priyanshu" // will store as slice of []byte
	str := "hello world"
	fmt.Println(name[0]) // 80 'byte of p'
	fmt.Printf("%c \n", name[0]) //p
	fmt.Printf("%c \n", name[1]) //r
	fmt.Printf("%c \n", name[2]) //i
	fmt.Printf("%c \n", name[3]) //y
	fmt.Printf("%c \n", str[0])
	fmt.Println(len(name))
	fmt.Println(len(str))


	// looping over a string by byte(uint8)

	// for i := 0; i<len(name); i++ {
	// 	fmt.Printf("%c \n", name[i])
	// }

	// looping over a string by rune(int32)
	// range return two outputs: 1. index, 2. Output

	// for i, ch := range str {
	// 	fmt.Printf("Index: %d, char: %c\n", i,ch)
	// }



}

