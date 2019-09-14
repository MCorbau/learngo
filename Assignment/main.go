package main

import (
	"fmt"
	"path"
)

func main() {
	//Exercise 1
	color := "green"
	color = "blue"
	fmt.Println(color)

	//Exercise 2
	color = "green"
	color = "dark " + color
	fmt.Println(color)

	//Exercise 3
	n := 0.
	n = 3.14 * 2
	fmt.Println(n)

	//Exercise 4
	var (
		perimeter     int
		width, height = 5, 6
	)
	perimeter = 2 * (width + height)
	fmt.Println(perimeter)

	//Exercise 5
	var (
		lang    string
		version int
	)

	lang, version = "go", 2

	fmt.Println(lang, "version", version)

	//Exercise 6
	var (
		planet string
		isTrue bool
		temp   float64
	)

	planet, isTrue, temp = "Mars", true, 19.5

	fmt.Println("Air is good on", planet)
	fmt.Println("It's", isTrue)
	fmt.Println("It is", temp, "degrees")

	//Exercise 7
	_, b := multi()
	fmt.Println(b)

	//Exercise 8
	color, color2 := "red", "blue"
	color, color2 = "orange", "green"
	fmt.Println(color, color2)

	//Exercise 9

	red, blue := "red", "blue"
	red, blue = blue, red

	fmt.Println(red, blue)

	//Exercise 9

	p, _ := path.Split("secret/file.txt")
	fmt.Println(p)

}

func multi() (int, int) {
	return 5, 4
}
