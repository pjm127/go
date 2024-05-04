package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin) // 표준입력
	for {
		fmt.Println("숫자를 입력하세요")
		var number int
		_, err := fmt.Scanln(&number) // 앞에가 입력개수인데 활용안하면 _ 처리
		if err != nil {
			fmt.Println("숫자로 입력하세요")
			//키보드 버퍼 삭제
			stdin.ReadString('\n')
			continue

		}
		fmt.Printf("입력하신 숫자는 %d입니다. \n", number)
		if number%2 == 0 {
			break
		}

	}
	fmt.Println("for 종료")
}
