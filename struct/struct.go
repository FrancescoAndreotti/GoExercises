package main

import "fmt"

type House struct {
	numRooms int
	price    float32
	city     string
}

func main() {
	var firtsHouse House

	firtsHouse.numRooms = 3
	firtsHouse.price = 112
	firtsHouse.city = "Napoli"

	fmt.Println(firtsHouse)
}
