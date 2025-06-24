/* 
package os // import "os"

func Open(name string) (*File, error)
    Open opens the named file for reading. If successful, methods on the
    returned file can be used for reading; the associated file descriptor has
    mode O_RDONLY. If there is an error, it will be of type *PathError.
*/

package main

import (
	"fmt"
	"os"
)

func main() {

	path := "/workspaces/Golang/Strings_IO_and_ErrorHandling/FileHandling/temp.txt"

	file, _ := os.Open(path)
	b := make([]byte, 4)

	for {

		n, err := file.Read(b)
		if err!=nil {
			fmt.Println("Error", err)
			break
		}
		fmt.Println(b[0:n])
		fmt.Println(string(b[0:n]))
	}
	
}

// more to understand


