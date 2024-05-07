package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.3, 35.5, 21.1, 43.3}

	for i := 0; i < 5; i++ {
		fmt.Println(t[i])
	}
}
