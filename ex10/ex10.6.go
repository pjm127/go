package main

import (
	"fmt"
)

func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		fallthrough // 이러면 break안하고 밑에도 실행
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("default")

	}

}
