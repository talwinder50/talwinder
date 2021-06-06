/*
Copyright SecureKey Technologies Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package session

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAuthenticate exhibits the user session authenticate
func TestAuthenticate(t *testing.T) {
	s := &UserSession{}
	auth := s.Authenticate("username", "password")
	assert.Equal(t, true, auth)
}
