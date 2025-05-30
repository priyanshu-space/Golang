package main
import "fmt"

func main() {

	var name string
	var is_muggle bool

	fmt.Print("Enter you name and muggle type: ")
	fmt.Scanf("%s %t", &name, &is_muggle)
	fmt.Println("Hey your name is",name," of muggle type as",is_muggle)

}