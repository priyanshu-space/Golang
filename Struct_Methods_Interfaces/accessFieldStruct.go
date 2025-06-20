package main
import "fmt"

type Circle struct {

	radius int
	x int
	y int

}
func main() {

	var c Circle
	c.x=5
	c.y=6
	c.radius=5

	fmt.Printf("%+v \n", c)


}