package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	p := fmt.Println

	p(rand.Intn(100), ",")
	p(rand.Intn(100))
	p()

	p(rand.Float64())

	p(rand.Float64())
	p((rand.Float64()*5)+5, ",")
	p((rand.Float64() * 5) + 5)
	p()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	p(r1)
}
