package main
import "fmt"

func main() {

	codes := map[string]string{"en": "English", "fr": "French"}
	fmt.Println(codes)
	fmt.Println(len(codes))
	fmt.Println(codes["en"])
	fmt.Println(codes["fr"])

	//getting a key value in map

	value, found := codes["en"]
	fmt.Println(value, found)
	value, found = codes["hh"]
	fmt.Println(value, found)

	//Adding a new value to the map

	codes["it"] = "Italian"
	fmt.Println(codes)

	//delete a key-value pair in map

	delete(codes, "en")
	fmt.Println(codes)

}