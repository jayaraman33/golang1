package main

import "fmt"

func main() {

	fmt.Printf("Enter the string")
	var myString string
	fmt.Scanln(&str)
	count := 0
	for _, ch := range str {
		switch ch {
		case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
			count++
		}
	}
	fmt.Println(" %s string contains vowels count: %d\n", str, count)

}
