package main

import "fmt"

func main() {
	var smallnum int
	var largenum int
	fmt.Print("Please enter a small number: ")
	fmt.Scan(&smallnum)
	fmt.Print("Please enter a large number: ")
	fmt.Scan(&largenum)
	var remainder = largenum % smallnum
	fmt.Println("The remainder is:", remainder)
}
