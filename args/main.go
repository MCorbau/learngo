package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Count the Arguments
//
//  Print the count of the command-line arguments
//
// INPUT
//  bilbo ba

func main() {

	//Exercise 1
	// UNCOMMENT & FIX THIS CODE
	arg := os.Args
	count := len(arg) - 1

	// UNCOMMENT IT & THEN DO NOT TOUCH THIS CODE
	fmt.Printf("There are %d names.\n", count)

	//Exercise 2
	fmt.Printf(arg[0])

	//Exercise 3
	fmt.Println("Hi ", arg[1])
	fmt.Println("How are you?")

	//Exercise 4
	fmt.Println("Hello great ", arg[1])
	fmt.Println("Hello great ", arg[2])
	fmt.Println("Hello great ", arg[3])
	fmt.Println("Nice to meet you all")

	//Exercise 4
	fmt.Println("Hello great ", arg[1])
	fmt.Println("Hello great ", arg[2])
	fmt.Println("Hello great ", arg[3])
	fmt.Println("Hello great ", arg[4])
	fmt.Println("Hello great ", arg[5])
	fmt.Println("Nice to meet you all")
}
