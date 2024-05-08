package main

import "fmt"

func main() {

	str := "hello 월드"
	//arr := []rune(str)
	//[]int 동적배열
	//rune -> int32의 별칭타입

	//for i := 0; i < len(str); i++ {
	//	fmt.Printf("타입:%T 값:%d 문자값:%c\n", str[i], str[i], str[i])
	//}
	/*for i := 0; i < len(arr); i++ {
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", arr[i], arr[i], arr[i])
	}*/

	for _, v := range str {
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", v, v, v)
	}

}
