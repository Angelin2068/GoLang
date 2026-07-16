package main

import "fmt"

func main() {
	str := "Hello Go"
	count := 0

	for _, ch := range str {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}

	fmt.Println("Vowel Count =", count)
}
