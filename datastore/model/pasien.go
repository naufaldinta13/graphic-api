// Copyright 2019 KORA ID. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package model

import (
	"github.com/naufaldinta13/cuxs/orm"
)

func init() {
	orm.RegisterModel(new(Pasien))
}

// Pasien model for pasien table.
type Pasien struct {
	ID   int64  `orm:"column(id);auto" json:"id"`
	Name string `orm:"column(name);size(45)" json:"name"`
}

// Save inserting or updating Pasien struct into pasien table.
// It will updating if this struct has valid Id
// if not, will inserting a new row to pasien.
// The field parameter is an field that will be saved, it is
// usefull for partial updating data.
func (m *Pasien) Save(fields ...string) (err error) {
	o := orm.NewOrm()
	if m.ID > 0 {
		_, err = o.Update(m, fields...)
	} else {
		m.ID, err = o.Insert(m)
	}
	return
}

// Delete permanently deleting pasien data
// this also will truncated all data from all table
// that have relation with this pasien.
func (m *Pasien) Delete() (err error) {
	_, err = orm.NewOrm().Delete(m)
	return
}

// Read execute select based on data struct that already assigned.
func (m *Pasien) Read(fields ...string) error {
	o := orm.NewOrm()
	return o.Read(m, fields...)
}
