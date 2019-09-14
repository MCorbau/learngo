package main

import "fmt"

func main() {
	var test bool
	test = true
	prenom, nom := "Martin", "Corbau"
	fmt.Println(prenom, nom, test)

	//Exercise 1`
	i, f, s, b := 314, 3.14, "Hello", true
	fmt.Println(
		"i:", i,
		"f:", f,
		"s:", s,
		"b:", b,
	)

	//Exercise 2

	a, b := 14, true
	fmt.Println(a, b)

	//Exercise 3
	a, c := 42, "Good"
	fmt.Println(a, c)

	//Exercise 4
	sum := 27 + 3.5
	fmt.Println(sum)

	//Exercise 5
	on, off := true, false

	_ = off
	fmt.Println(on)

	//Exercise 6
	age, yourAge := 23, 23
	ratio, age := 3.14, 42
	fmt.Println(age, yourAge, ratio)
}
