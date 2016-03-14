package main

import (
	"github.com/nu7hatch/gouuid"
	"net/http"
	"fmt"
)
func main()  {
	http.HandleFunc("/", func(res http.ResponseWriter, req, *http.Request){
		cookie, err := req.Cookie("session-fino")

		if err != nil{
			id, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name: "session-fino",
				Value: id.String(),
				HttpOnly: true;
			}
			http.SetCookie(res, cookie)
		}
		fmt.Println(cookie)
	})
	http.ListenAndServe(":8080", nil)
}
