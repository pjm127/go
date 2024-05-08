package main

import "fmt"

type User struct {
	name string
	age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u
}

func main() {
	userPointer := NewUser("AAA", 20)

	fmt.Println(userPointer)

}
