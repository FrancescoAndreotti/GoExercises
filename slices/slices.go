package main

import "fmt"

func main() {

	mystring := "hello world"

	s1 := mystring[0 : len(mystring)/2]
	s2 := mystring[len(mystring)/2+1 : len(mystring)]
	fmt.Println(s1)
	fmt.Println(s2)
}
