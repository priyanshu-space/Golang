package main
import (
	"fmt"
	"path/filepath"
)

func main() {

	path := filepath.Join("dir1","dir2","file.txt") // prints the filepath 
	fmt.Println(path)
	fmt.Println(filepath.Dir(path)) // returns the dir path of the file(last argument)
	fmt.Println(filepath.Base(path)) //returns the last element of the path
	fmt.Println(filepath.Ext(path)) // returns the extension of the file
	
}