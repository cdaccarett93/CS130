package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func Webpage(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	cookie, err := req.Cookie("cookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "cookie",
			Value: "0",
		}
	}

	count, _ := strconv.Atoi(cookie.Value)
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(res, cookie)

	io.WriteString(res, cookie.Value)
}

func main() {
	http.HandleFunc("/", Webpage)

	fmt.Println("Server Is Up!")
	http.ListenAndServe(":8080", nil)
}
