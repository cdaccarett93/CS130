package main

import "fmt"

func main() {
	fmt.Print("What is your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello,", name)
}
