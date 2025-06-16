package main
import "fmt"

func sumNumbers(numbers ...int) int {
	sum := 0
	for  _, value := range numbers {
		sum += value
	}
	return sum
}

func printDetails(student string, subjects ...string) {
	fmt.Println("hey ", student, ", here are your subjects - ")
	for _, sub := range subjects {
		fmt.Printf("%s, ", sub)
	}
}

func main() {

	fmt.Println(sumNumbers())
	fmt.Println(sumNumbers(10))
	fmt.Println(sumNumbers(10, 20))
	fmt.Println(sumNumbers(	110,20,89,1,233))
	fmt.Println(sumNumbers(1, -1))
	printDetails("joe", "physics", "biology")
}