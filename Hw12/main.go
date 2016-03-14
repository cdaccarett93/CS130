package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func EnterName(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("nameTemplate.html")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.Execute(res, nil)
	fmt.Fprintf(res, "%v", req.FormValue("name"))
}

func main() {
	http.HandleFunc("/", EnterName)
	http.ListenAndServe(":8080", nil)
}
