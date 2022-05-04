package main

import "fmt"

var (
	Marco  float32 = 110
	Luca   float32 = 68
	Franco float32 = 80
	Mario  float32 = 92
	Andrea float32 = 72.4
)

func main() {
	media := (Marco + Luca + Franco + Mario + Andrea) / 5
	fmt.Printf("Il peso medio è %.2f\n", media)
}
