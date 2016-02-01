package main

import "fmt"

func largest(nums ...int) int {
	var num int
	for _, v := range nums {
		if v > num {
			num = v
		}
	}
	return num
}
func main() {
	value := largest(2, 8, 1, 3, 0, 8004, 41, 20)
	fmt.Println(value)
}
