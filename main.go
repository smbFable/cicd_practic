package main

import "fmt"

func Add(a, b int) int {
	return 2 * a + b
}

func main() {
	fmt.Println("2 + 2 =", Add(2, 2))
}
