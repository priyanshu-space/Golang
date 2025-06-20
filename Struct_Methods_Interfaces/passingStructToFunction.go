package main
import "fmt"

type Circle struct {
	x int
	y int
	radius float64
	area float64
}


// Passing by value

/*
func calcArea(c Circle) {

	const PI float64 = 3.14
	var area float64
	area = (PI*c.radius*c.radius)
	c.area = area

}
*/

// Passing by reference

func calcArea(c *Circle) {

	const PI float64 = 3.14
	var area float64
	area = (PI*c.radius*c.radius)
	(*c).area = area

}

func main() {

	c := Circle{x:5, y:5, radius:7,area:0}
	fmt.Printf("%+v \n", c)

	//passing by value
	// calcArea(c)
	// fmt.Printf("%+v \n", c)

	//passing by reference

	calcArea(&c)
	fmt.Printf("%+v \n", c)
}