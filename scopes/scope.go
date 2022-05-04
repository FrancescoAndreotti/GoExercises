package main

import "fmt"

//Global Variable
var x int = 10

func main() {
	//Local Variable
	var l int = 7

	fmt.Println("This is a local variable: ", l)
	fmt.Println("This is a global variable: ", x)
}
