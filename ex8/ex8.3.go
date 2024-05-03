package main

import "fmt"

const (
	Pig     int = iota
	Cow     int = iota
	Chicken int = iota
)

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("Pig")
	} else if animal == Cow {
		fmt.Println("Cow")
	} else if animal == Chicken {
		fmt.Println("Chicken")
	} else {
		fmt.Println("Animal")
	}
}

func main() {
	PrintAnimal(Pig)
	PrintAnimal(Cow)
	PrintAnimal(6)
}
