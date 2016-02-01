package main

import "fmt"

func main() {
	half := func(i int) (int, bool) {
		return i / 2, i%2 == 0
	}
	fmt.Println(half(10))
}
