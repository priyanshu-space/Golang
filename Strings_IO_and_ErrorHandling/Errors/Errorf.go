/*
package fmt // import "fmt"

func Errorf(format string, a ...any) error
    Errorf formats according to a format specifier and returns the string as a
    value that satisfies error.

    If the format specifier includes a %w verb with an error operand,
    the returned error will implement an Unwrap method returning the operand.
    If there is more than one %w verb, the returned error will implement an
    Unwrap method returning a []error containing all the %w operands in the
    order they appear in the arguments. It is invalid to supply the %w verb
    with an operand that does not implement the error interface. The %w verb is
    otherwise a synonym for %v.
*/


package main

import (
	"fmt"
)

func process(i int) error {
	if i%2 == 0 {
		return fmt.Errorf("Only odd numbers are allowed, got %d", i)
	}
	return nil
}

func checkError(e error) {
	if e!=nil {
		fmt.Println(e)
	} else {
		fmt.Println("Operation Sucessful")
	}
}

func main() {

	err := process(3)
	checkError(err)

	err = process(2)
	checkError(err)

}