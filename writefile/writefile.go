package main

import "os"

func main() {
	file, err := os.Create("writefile/file.txt")

	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString("Roma, Milano, Napoli, Lucca, Pisa")
}
