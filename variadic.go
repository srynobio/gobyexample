package main

import "fmt"

func sum(num ...int) {
	fmt.Print(num, "* ")
	total := 0
	for _, num := range num {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	sum(1, 2, 3, 1, 2, 3)

	//	nums := []int{1, 2, 3, 4}
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6}
	sum(nums...)
}
