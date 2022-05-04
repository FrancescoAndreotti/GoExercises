package main

import "fmt"

func students(student ...string) {
	fmt.Println(student)
}

func main() {
	students("Carlo", "Pietro", "Mirco")
}
