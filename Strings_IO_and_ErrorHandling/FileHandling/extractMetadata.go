package main
import (
	"fmt"
	"os"
)

func main() {

	fileInfo, err := os.Stat("/workspaces/Golang/Strings_IO_and_ErrorHandling/FileHandling/temp.txt")

	if err!=nil {
		fmt.Println(err)
	}

	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size()) //returns 24 byte
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.IsDir())

}