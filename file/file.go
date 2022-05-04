package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("/Users/francescoandreotti/progetti/marchiedisegni-ipmonitor/docker-compose.yml"); err == nil {
		fmt.Printf("File exists\n")
	} else {
		fmt.Printf("File does not exist\n")
	}
}
