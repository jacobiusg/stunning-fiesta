package main

import (
	"fmt"

	"github.com/bentranter/chalk"
)

func main() {
	h := "Hello World!"
	i := 1
	for i <= 10 {
		fmt.Println(chalk.Green(h))
		i = i + 1
	}
}
