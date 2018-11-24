package main

import "fmt"

type point struct {
	x, y int
}

var pf = fmt.Printf

func main() {

	p := point{1, 2}
	pf("%v\n", p)

	pf("%+v\n", p)

	pf("%#v\n", p)

	pf("%T\n", p)

	pf("%t\n", p)

	pf("%d\n", p)

	pf("%b\n", p)

	pf("%c\n", p)

	pf("%x\n", p)

	pf("%f\n", p)

}
