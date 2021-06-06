/*
Copyright SecureKey Technologies Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package restapi

import (
	"encoding/json"
	"fmt"
	"github.com/trisolaria/talwinder/pkg/cookie"
	"github.com/trisolaria/talwinder/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"log"

	"net/http"

	"github.com/jinzhu/gorm"
)

type Config struct {
	Db *gorm.DB
	store       *cookie.Store
	Cookie      *cookie.Config
}

const cookieName = "userSession"
// Create a struct that models the structure of a user, both in the request body, and in the DB

type Credentials struct {
	Password string
	Username string
}

func (cfg *Config) SignIn(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		WriteErrorResponsef(w,
			http.StatusBadRequest, "failed to decode credentials: %s", err.Error())
		return
	}

	var user models.User

	cfg.Db.Where("username = ?", creds.Username).First(&user)

	if user.Id == 0 {
		WriteErrorResponsef(w,
			http.StatusNotFound, "failed to find any record")
		return
	}

	// Salt and hash the password using the bcrypt algorithm
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(creds.Password))
	if  err != nil {
		WriteErrorResponsef(w,
			http.StatusInternalServerError, "Oops password not matched: %s", err.Error())
		return
	}

    store:= cookie.NewStore(cfg.Cookie)

	err = saveSession(store, r, w, user.Id)
	if err != nil {
		WriteErrorResponsef(w,
			http.StatusInternalServerError, "Failed to read from session: %s", err.Error())
		return
	}

	fmt.Println(w.Write([]byte("Successfully Logged In")))
}

func saveSession(cookieSession *cookie.Jars, r *http.Request, w http.ResponseWriter, user uint) error{
	 cookieJar, err := cookieSession.Open(r)
	if err != nil {
		return	err
	}

	cookieJar.Set(cookieName, user)

	err = cookieJar.Save(r, w)
	if err != nil {
	   return fmt.Errorf("failed to save user login cookie%s", err)
	}

    return nil
}


// ErrorResponse to send error message in the response.
type ErrorResponse struct {
	Message string `json:"errMessage,omitempty"`
}

// WriteErrorResponsef write error resp.

func WriteErrorResponsef(rw http.ResponseWriter, status int, msg string, args ...interface{}) {
	rw.WriteHeader(status)

	err := json.NewEncoder(rw).Encode(ErrorResponse{
		Message: fmt.Sprintf(msg, args...),
	})
	if err != nil {
		log.Println(err)
	}
}
