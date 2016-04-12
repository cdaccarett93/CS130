package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

func getData(data string) string {
	h := hmac.New(sha256.New, []byte("cd#r=*resw6rEjez"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Webpage(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
		fmt.Println("Created cookie!")
	}

	if req.Method == "POST" {
		newValue := req.FormValue("cookieType")
		userCookie := getData(newValue)
		tempValue := strings.Split(cookie.Value, " | ")
		cookie.Value = tempValue[0] + " | " + userCookie
		tempValue = nil
		http.SetCookie(res, cookie)
		fmt.Println("Updated cookie!")
	}

	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.Execute(res, nil)
}

func main() {
	// set the path URL
	http.HandleFunc("/", Webpage)

	http.ListenAndServe(":8080", nil)
}
