package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {

	http.HandleFunc("/", surfPage)
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))
	http.ListenAndServe(":8080", nil)
}

func surfPage(res http.ResponseWriter, req *http.Request) {

	var err error

	surferPage := template.New("surfPage.html")
	surferPage, err = surferPage.ParseFiles("surfPage.html")

	if err != nil {
		log.Fatalln(err)
	}

	err = surferPage.Execute(res, nil)

	if err != nil {
		log.Fatalln(err)
	}
}
