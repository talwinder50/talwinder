package main

import (
	"encoding/gob"
	"github.com/gorilla/securecookie"
	"github.com/trisolaria/talwinder/pkg/cookie"
	"github.com/trisolaria/talwinder/pkg/models"
	"github.com/trisolaria/talwinder/pkg/restapi"
	"net/http"

	"github.com/trisolaria/talwinder/pkg/conn"
)

const defaultMaxAge = 900
// For testing only
func main() {
	config := conn.SophonicConnection{
		ServerName: "localhost:3306",
		User:       "root",
		Password:   "interview",
		DB:         "user",
	}

	connectionString := conn.GetConnectionString(config)
	err := conn.ConnectSophon(connectionString)
	if err != nil {
		panic(err.Error())
	}

	cfg := restapi.Config{
		Db: conn.Connector,
		Cookie: &cookie.Config{
			AuthKey: securecookie.GenerateRandomKey(64),
			EncKey: securecookie.GenerateRandomKey(32),
			MaxAge: defaultMaxAge,
		},
	}

	http.HandleFunc("/signIn", cfg.SignIn)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	gob.Register(models.User{})
}


