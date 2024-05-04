package main

import "fmt"

func main() {
	a := 4

	switch a {
	case 1, 4:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}

}
