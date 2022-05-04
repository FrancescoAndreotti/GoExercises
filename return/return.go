package main

import "fmt"

func values() (string, string) {
	return "Hello", "World!"
}

func main() {
	x, y := values()
	fmt.Println(x + " " + y)
}
