package main

import "net/http"

func serveIt(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain")
	res.Write([]byte("Test.\n"))
}

func main() {
	http.HandleFunc("/", serveIt)
	http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
}