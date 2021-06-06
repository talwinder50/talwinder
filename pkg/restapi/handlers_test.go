/*
Copyright SecureKey Technologies Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package restapi

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/securecookie"
	"github.com/trisolaria/talwinder/pkg/conn"
	"github.com/trisolaria/talwinder/pkg/cookie"
)

const testCredentialRequestBody = `{
"username":"test",
"password":"waheguru"
}`

func TestSignInHandler(t *testing.T) {
	setUpDB()

	cfg := Config{
		Db: conn.Connector,
		Cookie: &cookie.Config{
			AuthKey: securecookie.GenerateRandomKey(64),
			EncKey:  securecookie.GenerateRandomKey(32),
			MaxAge:  900,
		},
	}

	t.Run("Successfully Login", func(t *testing.T) {
		req, err := http.NewRequest("Post", "/signIn", bytes.NewBuffer([]byte(testCredentialRequestBody)))
		if err != nil {
			t.Fatal(err)
		}
		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(cfg.SignIn)

		handler.ServeHTTP(rr, req)
		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check the response body is what we expect.
		expected := `Successfully Logged In`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}

// Not a good practice
func setUpDB() {
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
}
