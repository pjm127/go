package main

import "fmt"

func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	d, success := Divide(9, 0) //success 한개만 있으면 중복이라 안되는데 여러개라서 중복선언 가능
	fmt.Println(d, success)
}
