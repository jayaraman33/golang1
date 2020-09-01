package main

import "fmt"

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func main() {
	fmt.Print("Enter String: ")
	var first string
	fmt.Scanln(&first)
	fmt.Printf("%v\n", reverse(first))
}
