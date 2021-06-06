/*
Copyright SecureKey Technologies Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name" gorm:"unique"`
	Password []byte `json:"-"`
}
