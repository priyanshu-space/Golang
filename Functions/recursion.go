package main
import "fmt"

func factorial(n int) int {
	if n==0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {

	n := 5
	result := factorial(n)
	fmt.Println("factorial of", n, "is:",result)
}

/*
Execution :

return 5 * factorial(4) = 120
	return 4 * factorial(3) = 24
			return 3 * factorial(2) = 6
				return 2 * factorial(1) = 2
					return 1 * factorial(0) = 1					

*/