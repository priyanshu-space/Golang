package main
import "fmt"

type Student struct {

	//field names
	name string
	rollNo int
}

func main() {

	//1.

	// var st Student //This will only initialize struct with default or zero values.
    
	//2.
	// st := new(Student) // This will only initialize struct with default or zero values and returns pointer to the struct. Here st acts as a pointer to struct Student.

	/* 3.
	For Initializing struct with some values 

	 st := Student {
		name: "joe",
		rollNo: 12,
	} */

	//4.

	st := Student{"joe", 12}

	fmt.Printf("%+v \n", st)
	fmt.Println(st) // skip field and print only field values
}