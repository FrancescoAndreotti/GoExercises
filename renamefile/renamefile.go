package main

import "os"

func main() {

	src := "readfile/name.txt"
	dst := "ciao.txt"

	os.Rename(src, dst)
}
