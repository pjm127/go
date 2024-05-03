package main

import "fmt"

func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0 // 반환값 이름 적으면 리턴값없이 리턴 가능
		success = false
		return
	}
	result = a / b
	success = true
	return
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	d, success := Divide(9, 0) //success 한개만 있으면 중복이라 안되는데 여러개라서 중복선언 가능
	fmt.Println(d, success)
}
