package main

import "fmt"

func main() {
	val := 600851475143
	largest := 1

	if val%2 == 0 {
		val = val / 2
		largest = 2
		if val%2 == 0 {
			val = val / 2
		}
	}

	current := 3
	for val > 1 {
		if val%current == 0 {
			val = val / current
			largest = current
			if val%current == 0 {
				val = val / current
			}
		}
		current += 2
	}
	fmt.Println(largest)
}
