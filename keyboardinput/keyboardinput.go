package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scanf("%d", &i)

	if err != nil {
		fmt.Println(err)
	}

	switch {
	case i < 1:
		fmt.Println("Il numero inserito è minore di 1")
	case i < 10:
		fmt.Println("Il numero inserito è compreso tra 1 e 10")
	case i > 10:
		fmt.Println("Il numero inserito è maggiore di 10")
	}

}
