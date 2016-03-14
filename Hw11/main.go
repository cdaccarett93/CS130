package main

import (
	"fmt"
	"net/http"
)

func formvalue(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	fmt.Fprintf(w, "Hello %s", name)
}

func main() {
	http.HandleFunc("/", formvalue)
	http.ListenAndServe(":8080", nil)

}
