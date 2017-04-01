package main

import "fmt"

func main() {
	h := "Hello everyboday!!"
	i := 1
	for i <= 20 {
		fmt.Println(chalk.red(h))
		i = i + 1
	}
}
