package main

import "fmt"

func rec(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	rec(n - 1)
	fmt.Println("After", n)
}

func main() {
	rec(3)

}
