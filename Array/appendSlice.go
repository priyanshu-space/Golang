package main

import "fmt"

func main() {
        arr := [5]int{10, 20, 90, 70, 60}
        // slice := append(arr[:2], arr[3:]) : This line will throw error, as it can't be directlY appended.
		slice := append(arr[:2], arr[3:]...)
        fmt.Println(slice)
}