package main
import "fmt"

func main() {

	codes := map[string]string{"en": "English", "fr": "French"}
	
	for key, value := range codes {
		fmt.Println(key, "=>", value)
	}

	//truncating a map

	for key, value := range codes {
		delete(codes, key)
	}
	fmt.Println(codes)
}