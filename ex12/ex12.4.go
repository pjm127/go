package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.6, 35.4, 45.5, 56.5}

	for i, v := range t {
		fmt.Println(i, v)
	}

}
