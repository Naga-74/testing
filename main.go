package main

import "fmt"

func main() {
	result := add(3, 4)
	fmt.Println("Result:", result)
}

func add(a, b int) int {
	return a + b
}
