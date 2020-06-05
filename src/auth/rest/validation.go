// Copyright 2018 Kora ID. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rest

import (
	"github.com/naufaldinta13/cuxs/orm"
	"github.com/naufaldinta13/graphic-api/datastore/model"
)

const (
	errRequiredPasien   = "pasien harus diisi."
	errRequiredUsername = "username harus diisi."
	errRequiredPassword = "password harus diisi."
	errInvalidPasien    = "Pasien tidak valid."
	errInvalidAuth      = "periksa kembali username dan password Anda."
)

func validPasien(id int64) (m *model.Pasien, e error) {
	e = orm.NewOrm().Raw("SELECT * FROM pasien WHERE id = ?", id).QueryRow(&m)

	return
}

func validUsername(username string) (m *model.User, e error) {
	e = orm.NewOrm().Raw("SELECT * FROM user WHERE username = ?", username).QueryRow(&m)

	return
}
