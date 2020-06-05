// Copyright 2019 KORA ID. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package model

import (
	"github.com/naufaldinta13/cuxs/common"
	"github.com/naufaldinta13/cuxs/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(User))
}

// User model for user table.
type User struct {
	ID             int64     `orm:"column(id);auto" json:"id"`
	Username       string    `orm:"column(username);size(45)" json:"username"`
	Password       string    `orm:"column(password);size(45)" json:"password"`
	Name           string    `orm:"column(name);size(45)" json:"name"`
	LastActivityAt time.Time `orm:"column(last_activity_at);type(timestamp);null" json:"last_activity_at"`

	Pasien *Pasien `orm:"-" json:"pasien"`
}

// Save inserting or updating User struct into user table.
// It will updating if this struct has valid Id
// if not, will inserting a new row to user.
// The field parameter is an field that will be saved, it is
// usefull for partial updating data.
func (m *User) Save(fields ...string) (err error) {
	o := orm.NewOrm()
	if m.ID > 0 {
		_, err = o.Update(m, fields...)
	} else {
		m.ID, err = o.Insert(m)
	}
	return
}

// Delete permanently deleting user data
// this also will truncated all data from all table
// that have relation with this user.
func (m *User) Delete() (err error) {
	_, err = orm.NewOrm().Delete(m)
	return
}

// Read execute select based on data struct that already assigned.
func (m *User) Read(fields ...string) error {
	o := orm.NewOrm()
	return o.Read(m, fields...)
}

func (m *User) PasswordMatched(pwd string) bool {
	return common.PasswordHash(m.Password, pwd) == nil
}

//GetUser is used by jwtUser to get model user
func (m *User) GetUser(id int64) (interface{}, error) {
	m.ID = id
	e := m.Read()
	return m, e
}
