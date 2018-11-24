package main

import "os"

func main() {

	panic("a issue!")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
