// Copyright 2019 KORA ID. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package model

import (
	"time"

	"github.com/naufaldinta13/cuxs/orm"
)

func init() {
	orm.RegisterModel(new(StatisticLog))
}

// StatisticLog model for statistic_log table.
type StatisticLog struct {
	ID         int64     `orm:"column(id);auto" json:"id"`
	Pasien     *Pasien   `orm:"column(pasien_id);rel(fk)" json:"pasien,omitempty"`
	Suhu       string    `orm:"column(suhu);size(45)" json:"suhu"`
	RecordedAt time.Time `orm:"column(recorded_at);type(timestamp)" json:"recorded_at"`
}

// Save inserting or updating StatisticLog struct into statistic_log table.
// It will updating if this struct has valid Id
// if not, will inserting a new row to statistic_log.
// The field parameter is an field that will be saved, it is
// usefull for partial updating data.
func (m *StatisticLog) Save(fields ...string) (err error) {
	o := orm.NewOrm()
	if m.ID > 0 {
		_, err = o.Update(m, fields...)
	} else {
		m.ID, err = o.Insert(m)
	}
	return
}

// Delete permanently deleting statistic_log data
// this also will truncated all data from all table
// that have relation with this statistic_log.
func (m *StatisticLog) Delete() (err error) {
	_, err = orm.NewOrm().Delete(m)
	return
}

// Read execute select based on data struct that already assigned.
func (m *StatisticLog) Read(fields ...string) error {
	o := orm.NewOrm()
	return o.Read(m, fields...)
}
