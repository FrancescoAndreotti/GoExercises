package main

import "fmt"

func main() {
	var a = []int64{1, 2, 3, 4}

	for index, element := range a {
		fmt.Print(index, ") ", element, "\n")
	}
}
