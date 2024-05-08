package main

import (
	"fmt"
	"strings"
)

func ToUpper1(str string) string {
	var rst string
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			//소문자면
			rst += string('A' + (c - 'a')) // 대문자로
		} else {
			rst += string(c)
		}
	}
	return rst
}

func ToUpper2(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			//소문자면
			builder.WriteRune('A' + (c - 'a')) //한글자씩 합쳐서
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main() {
	var str string = "Hello World"
	fmt.Println(ToUpper1(str))
	fmt.Println(ToUpper2(str))
	fmt.Println(strings.ToUpper(str)) //사실 strings 패키지에 있다

}
