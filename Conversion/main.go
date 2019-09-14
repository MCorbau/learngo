package main

import "fmt"

func main() {
	//Exercise 1
	a, b := 10, 5.5
	fmt.Println(float64(a) + b)

	//Exercise 2
	a, b = 10, 5.5
	a = int(b)
	fmt.Println(float64(a) + b)

	//Exercise 3
	fmt.Println(5.5)

	//Exercise 4
	age := 2
	fmt.Println(7.5 + float64(age))

	//Exercise 5
	min := int8(127)
	max := int16(1000)

	// FIX THE CODE HERE
	fmt.Println(max + int16(min))
}
