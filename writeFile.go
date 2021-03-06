package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat2", d1, 0644)
	check(err)

	f, err := os.Create("/tmp/dat3")
	check(err)
	defer f.Close()

	// f.write will take []byte and print as string.
	d2 := []byte{115, 111, 109, 101, 10, 99, 87}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("write\n")
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}
