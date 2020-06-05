// Copyright 2018 Kora ID. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rest

import (
	"github.com/naufaldinta13/cuxs/orm"
	"github.com/naufaldinta13/graphic-api/datastore/model"
	"net/url"
)

type ResponseStatistic struct {
	Day  string  `json:"day"`
	Suhu float64 `json:"suhu"`
}

func getLog(params url.Values) (mx []*model.StatisticLog, e error) {
	//default tanggal satu aja hehe
	day := "01"

	if recoredAt := params.Get("recorded_at"); recoredAt != "" {
		day = recoredAt
	}

	if _, e = orm.NewOrm().Raw("SELECT * FROM statistic_log WHERE day(recorded_at) = ? ORDER BY id desc;", day).QueryRows(&mx); e == nil {
		return mx, nil
	}

	return nil, e
}

func getStatistic(params url.Values) (mx []*ResponseStatistic, e error) {
	//default tanggal satu aja hehe
	day := "01"

	if recoredAt := params.Get("recorded_at"); recoredAt != "" {
		day = recoredAt
	}

	if _, e = orm.NewOrm().Raw("SELECT * FROM statistic_log WHERE day(recorded_at) = ? ORDER BY id desc;", day).QueryRows(&mx); e == nil {
		return mx, nil
	}

	return nil, e
}
