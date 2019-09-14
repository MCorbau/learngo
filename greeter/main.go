package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		name  = os.Args[1]
		name2 = os.Args[2]
		name3 = os.Args[3]
	)

	fmt.Println("Hello great", name, "!")
	fmt.Println("And hellooo to you magnificient", name2, "!")
	fmt.Println("Welcome", name3, "!")

	name, age := "bilbo baggins", 13

	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("And, I love adventures!")
}
