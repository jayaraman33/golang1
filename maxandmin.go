package main

import "fmt"

func main() {
	fmt.Println("Enter size of array")
	var n int
	fmt.Scanln(&n)
	var a [10]int
	fmt.Println("Enter the numbers")
	for i := 0; i < n; i++ {
		fmt.Scanln(&a[i])
	}
	var min = a[0]
	var max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}
