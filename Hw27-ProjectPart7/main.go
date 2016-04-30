package main

import (
	"net/http"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl, _ = template.ParseGlob("templates/*.htmltemp")
}

func index(res http.ResponseWriter, req *http.Request) {

	cookie := genCookie(res, req)
	u := decodeUser(cookie)

	if req.Method == "POST" {
		u.State = true
		u.Name = req.FormValue("name")
		u.Age = req.FormValue("age")

		xs := strings.Split(cookie.Value, "|")
		id := xs[0]

		cookie = currentVisitor(u, id)
		http.SetCookie(res, cookie)

	}
	tpl.ExecuteTemplate(res, "index.htmltemp", u)
}

func login(res http.ResponseWriter, req *http.Request) {

	cookie := genCookie(res, req)

	if req.Method == "POST" && req.FormValue("password") == "secret" {
		u := decodeUser(cookie)
		u.State = true
		u.Name = req.FormValue("name")

		xs := strings.Split(cookie.Value, "|")
		id := xs[0]

		cookie := currentVisitor(u, id)
		http.SetCookie(res, cookie)

		http.Redirect(res, req, "/", 302)
		return
	}
	tpl.ExecuteTemplate(res, "login.htmltemp", nil)
}

func logout(res http.ResponseWriter, req *http.Request) {
	cookie := newVisitor()
	http.SetCookie(res, cookie)
	http.Redirect(res, req, "/", 302)
}

func genCookie(res http.ResponseWriter, req *http.Request) *http.Cookie {

	cookie, err := req.Cookie("session-id")
	if err != nil {
		cookie = newVisitor()
		http.SetCookie(res, cookie)
	}

	if strings.Count(cookie.Value, "|") != 2 {
		cookie = newVisitor()
		http.SetCookie(res, cookie)
	}

	if tampered(cookie.Value) {
		cookie = newVisitor()
		http.SetCookie(res, cookie)
	}

	return cookie
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)
}