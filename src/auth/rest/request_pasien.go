// Copyright 2018 Kora ID. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rest

import (
	"fmt"
	"github.com/naufaldinta13/cuxs/validation"
	"github.com/naufaldinta13/graphic-api/datastore/model"
	"github.com/naufaldinta13/graphic-api/src/auth"
)

type pasienRequest struct {
	PasienID int64 `json:"pasien_id" valid:"required"`

	Session *auth.SessionData `json:"-"`
	Pasien  *model.Pasien     `json:"-"`
}

func (r *pasienRequest) Validate() *validation.Output {
	o := &validation.Output{Valid: true}

	var e error

	fmt.Println("-------kemari------")

	if r.PasienID != 0 {
		if r.Pasien, e = validPasien(r.PasienID); e != nil {
			o.Failure("pasien_id.invalid", errInvalidPasien)
		}
	}

	if r.Pasien != nil {
		r.Session.User.Pasien = r.Pasien
	}

	return o
}

func (r *pasienRequest) Messages() map[string]string {
	return map[string]string{
		"pasien_id.required": errRequiredPasien,
	}
}
