package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Println("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its the weekend!")
	default:
		fmt.Println("its the weekday.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool.")
		case int:
			fmt.Println("I'm a int.")
		default:
			fmt.Printf("Dont know type %T\n", t)
		}
	}

	z := "something"
	fmt.Println(z.(type))

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
