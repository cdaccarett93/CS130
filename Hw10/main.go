package main

import (
	"fmt"
	"net/http"
)

func reqUrl(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "localhost:8080/name%v", req.URL.Path)
}

func main() {
	http.HandleFunc("/", reqUrl)
	http.ListenAndServe(":8080", nil)
}
