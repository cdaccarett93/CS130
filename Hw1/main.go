package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
	Beer bool
}

func main() {
	p1 := person{
		Name: "Carlos",
		Age:  22,
	}
	if p1.Age > 21 {
		p1.Beer = true
	}
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
