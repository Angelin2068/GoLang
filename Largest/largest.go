package main

import "fmt"

func main() {
	a, b, c := 10, 25, 15
	largest := a

	if b > largest {
		largest = b
	}
	if c > largest {
		largest = c
	}

	fmt.Println("Largest =", largest)
}
