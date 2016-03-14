package main

import (
	"github.com/nu7hatch/gouuid"
	"html/template"
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"encoding/base64"
)

type User struct {
	Name string
	Age  string
}

func main() {

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		pName := req.FormValue("name")
		pAge := req.FormValue("age")

		currentUser := User{
			Name: pName,
			Age: pAge,
		}

		bs, err := json.Marshal(currentUser)
		if err != nil{
			fmt.Println(err)
		}

		jsonB64 := base64.StdEncoding.EncodeToString(bs)

		cookie, err := req.Cookie("session-fino")
		if err != nil {

			id, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name:  "session-fino",
				Value: id.String() + " " + pName + " " + pAge + " " + jsonB64,
				//Secure: true
				HttpOnly: true,
			}

			http.SetCookie(res, cookie)
		}

		err = tpl.Execute(res, nil)
		if err != nil {
			http.Error(res, err.Error(), 500)
			log.Fatalln(err)
		}

	})

	http.ListenAndServe(":8080", nil)
}