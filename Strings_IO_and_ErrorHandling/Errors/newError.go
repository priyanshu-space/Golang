/*
go doc errors new
package errors // import "errors"

func New(text string) error
    New returns an error that formats as the given text. Each call to New
    returns a distinct error value even if the text is identical.
*/
package main
import (
	"fmt"
	"errors"
)
func process(i int) error {
	if i%2==0 {
		return errors.New("Only odd numbers allowed")
	}
	
	return nil

}

func checkError(e error) {
	if e==nil {
		fmt.Println(e)
	}
	fmt.Println("Operation Sucessful")

}


func main() {
	err := process(3)
	checkError(err)
	err = process(2)
	checkError(err)
}
