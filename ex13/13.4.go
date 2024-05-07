package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32   //4byte
	Score float64 //8byte
}

func main() {
	user := User{23, 55.3}
	fmt.Println(unsafe.Sizeof(user)) //unsafe.Sizeof 사이즈 크기

}
