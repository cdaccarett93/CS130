package main

import "fmt"

func main() {
	fmt.Println("Even numbers from 0 - 100 are: ")
	for even := 1; even <= 100; even++ {
		if even%2 == 0 {
			fmt.Println(even)
		}
	}
}
