/* func ReadFile(name string) ([]byte, error)
    ReadFile reads the named file and returns the contents. A successful call
    returns err == nil, not err == EOF. Because ReadFile reads the whole file,
    it does not treat an EOF from Read as an error to be reported. */


package main
import (
	"fmt"
	"os"
)

func main() {

	path := "/workspaces/Golang/Strings_IO_and_ErrorHandling/FileHandling/temp.txt"

	data, err := os.ReadFile(path)

	if err!=nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))


}