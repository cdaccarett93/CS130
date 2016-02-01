package main

import "fmt"

func half(i int) (int, bool) {
	return i /2, i%2 == 0
}

func main() {
	a, num := half(10)
	fmt.Println(a, num)
}
