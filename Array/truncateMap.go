package main
import "fmt"

func main() {

	codes := map[string]string{"en": "English", "fr": "French"}
    fmt.Println(codes)

	//truncating a map

	// for key, value := range codes : throwing error --> declared and not used: value
	// used "_" for unused values
	
	for key, _ := range codes {
		delete(codes, key)
	}
	fmt.Println(codes)
}