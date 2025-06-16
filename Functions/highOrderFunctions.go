package main
import "fmt"

// find area, perimeter and diameter according to the query

func printResults() {

}

func getFunction(query int) func(r float64) float64 {

	query_to_func := map[int]func(r float64) float64{
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter
	}

	 return query_to_func[query]
}