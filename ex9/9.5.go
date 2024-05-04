package main

import "fmt"

func HasRichFriend() bool {
	return true
}

func GetFriendsCount() int {
	return 3
}

func main() {
	price := 35000

	if price >= 50000 {
		if HasRichFriend() {
			fmt.Println("앗 신발끈이")
		} else {
			fmt.Println("나눠서..")
		}
	} else if price > 30000 {
		if GetFriendsCount() > 3 {
			fmt.Println("아이쿠 신발끈이")
		} else {
			fmt.Println("사람이 없어서 나눠내자")
		}
	} else {
		fmt.Println("내가 쏜다")
	}

}
