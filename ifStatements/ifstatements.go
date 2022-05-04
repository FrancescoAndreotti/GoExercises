package main

import "fmt"

func main() {
	var x float32
	fmt.Println("Inserisci un numero: ")
	_, err := fmt.Scanf("%f", &x)
	if err != nil {
		fmt.Println(err)
	}
	if x > 0 {
		x = x / 2
	}
	fmt.Println(x)
}
