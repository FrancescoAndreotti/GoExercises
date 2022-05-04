package main

import "fmt"

func sum(a, b int) int {
	return a + b
}
func main() {
	result := sum(100, 500)
	fmt.Println(result)
}
