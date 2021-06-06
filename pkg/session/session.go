/*
Copyright SecureKey Technologies Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package session

import "github.com/trisolaria/talwinder/pkg/crypt"

//UserSession exposes the ability to authenticate a user's password against the underlying crypt Authenticator

type UserSession struct {
	auth *crypt.IndeterminantAuthenticator
}

func (s *UserSession) Authenticate(username, password string) bool {

	if !s.auth.Authenticate(username, password) {
		return false
	}

	return true
}
