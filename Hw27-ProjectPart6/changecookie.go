package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/nu7hatch/gouuid"
)

type User struct {
	Age  string
	Name string
}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func loginHandler(res http.ResponseWriter, req *http.Request) {
	temp, err := template.ParseFiles("template.htmltemp")
	if err != nil {
		log.Fatalln(err)
	}

	if req.Method == "POST" {
		createCookie(res, req)
	}

	temp.Execute(res, nil)
}

func createCookie(res http.ResponseWriter, req *http.Request) {
	newUser := User{
		Age:  req.FormValue("name"),
		Name: req.FormValue("age"),
	}

	userSlice, err1 := json.Marshal(newUser)
	if err1 != nil {
		log.Fatalln(err1)
	}
	encodedData := base64.StdEncoding.EncodeToString(userSlice)
	code := getCode(encodedData)

	cookie, err2 := req.Cookie("session-fino")
	if err2 == http.ErrNoCookie {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-fino",
			Value: id.String() + "|" + encodedData + "|" + code,
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}

	xs := strings.Split(cookie.Value, "|")
	data := xs[1] + "invalidate me"
	usrCode := xs[2]

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	if usrCode == getCode(data) {
		io.WriteString(res, "Code valid")
	} else {
		io.WriteString(res, "Code Invalid")
	}

}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", loginHandler)
	http.ListenAndServe(":8080", nil)
}