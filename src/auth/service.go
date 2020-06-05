// Copyright 2016 PT. Qasico Teknologi Indonesia. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/naufaldinta13/cuxs/cuxs"
	"github.com/naufaldinta13/graphic-api/datastore/model"
)

// SessionData structur data current user logged in.
type SessionData struct {
	Token string      `json:"token"`
	User  *model.User `json:"user"`
}

// StartSession mendapatkan data user entity dengan token
// untuk menandakan session user yang sedang login.
func StartSession(user *model.User, token ...string) (sd *SessionData, e error) {
	sd = new(SessionData)
	// buat token baru atau menggunakan yang sebelumnya
	if len(token) == 0 {
		sd.Token = cuxs.JwtToken("id", user.ID)
	} else {
		sd.Token = token[0]
	}

	sd.User = user

	return
}

// UserSession mendapatkan session data dari user yang mengirimkan request.
func UserSession(ctx *cuxs.Context) (*SessionData, error) {
	if u := ctx.Get("user"); u != nil {
		c := u.(*jwt.Token).Claims.(jwt.MapClaims)
		var userID int64

		// id adalah user id
		if c["id"] != nil {
			userID = int64(c["id"].(float64))
		}

		// memakai token sebelumnya
		token := ctx.Get("user").(*jwt.Token).Raw

		user := &model.User{ID: userID}
		if e := user.Read("ID"); e == nil {
			fmt.Println("--------kesini ?----------", )

			return StartSession(user, token)
		}

		fmt.Println("--------kemana ?----------", )

	}

	return nil, errors.New("Invalid jwt token")
}
