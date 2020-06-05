// Copyright 2018 Kora ID. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rest

import (
	"github.com/naufaldinta13/cuxs/validation"
	"github.com/naufaldinta13/graphic-api/datastore/model"
)

type signinRequest struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`

	UserData *model.User `json:"-"`
}

func (r *signinRequest) Validate() *validation.Output {
	o := &validation.Output{Valid: true}

	var e error

	if r.Username != "" && r.Password != "" {
		// memastikan bahwa username yang di berikan valid
		if r.UserData, e = validUsername(r.Username); e != nil {
			o.Failure("username.invalid", errInvalidAuth)
		} else {
			if !r.UserData.PasswordMatched(r.Password) {
				o.Failure("username.invalid", errInvalidAuth)
			}
		}
	}

	return o
}

func (r *signinRequest) Messages() map[string]string {
	return map[string]string{
		"username.required": errRequiredUsername,
		"password.required": errRequiredPassword,
	}
}
