package main
import (

	"fmt"
	"time"
	"sync"

)

func calculateSquare(i int, wg *sync.WaitGroup) {

	fmt.Println(i*i)
	wg.Done()
}

func main() {

	var wg sync.WaitGroup  // sync.WaitGroup is of type struct 
	start_time := time.Now()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go calculateSquare(i, &wg)
	}
	elapsed_time := time.Since(start_time)
	wg.Wait()
	fmt.Println("Function took", elapsed_time)

}