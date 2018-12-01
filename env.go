package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	os.Setenv("FOO", "1")
	os.Setenv("HELLSBELLS", "HAVING ME RINGING")
	fmt.Println("FOO: ", os.Getenv("FOO"))
	fmt.Println("BAR: ", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}