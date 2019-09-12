package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println("test")
	fmt.Println(runtime.Version())
}
