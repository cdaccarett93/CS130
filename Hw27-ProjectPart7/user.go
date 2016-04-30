package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/nu7hatch/gouuid"
)

type User struct {
	Age   string
	Name  string
	State bool
}

func decodeUser(c *http.Cookie) User {
	xs := strings.Split(c.Value, "|")
	usrData := xs[1]

	bs, err := base64.URLEncoding.DecodeString(usrData)
	if err != nil {
		log.Println("Error decoding base64", err)
	}

	var u User
	err = json.Unmarshal(bs, &u)
	if err != nil {
		fmt.Println("error unmarshalling: ", err)
	}
	return u
}

func newVisitor() *http.Cookie {
	mm := initialUser()
	id, _ := uuid.NewV4()
	return makeCookie(mm, id.String())
}

func currentVisitor(u User, id string) *http.Cookie {
	mm := marshalUser(u)
	return makeCookie(mm, id)
}

func makeCookie(mm []byte, id string) *http.Cookie {
	b64 := base64.URLEncoding.EncodeToString(mm)
	code := getCode(b64)
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: id + "|" + b64 + "|" + code,
		// Secure: true,
		HttpOnly: true,
	}
	return cookie
}

func marshalUser(u User) []byte {
	bs, err := json.Marshal(u)
	if err != nil {
		fmt.Println("error: ", err)
	}
	return bs
}

func initialUser() []byte {
	u := User{
		Age:   "",
		Name:  "",
		State: false,
	}
	return marshalUser(u)
}