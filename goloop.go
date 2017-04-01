package main

import (
	"fmt"
)

func main() {
	var h string = "Hello, to the world!"
	var i int = 1
	for i <= 10 {
		fmt.Println(h)
		i = i + 1
	}
}
