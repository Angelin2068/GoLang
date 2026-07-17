package main

import "fmt"

func main() {
	str := "Golang"

	for i := len(str) - 1; i >= 0; i-- {
		fmt.Print(string(str[i]))
	}
	fmt.Println()
}
