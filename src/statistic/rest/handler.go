// Copyright 2018 Kora ID. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rest

import (
	"github.com/labstack/echo"
	"github.com/naufaldinta13/cuxs/cuxs"
	"github.com/naufaldinta13/graphic-api/datastore/model"
	"github.com/naufaldinta13/graphic-api/src/auth"
)

type Handler struct{}

func (h *Handler) URLMapping(r *echo.Group) {
	r.GET("/log", h.getLog, auth.CheckPrivilege(""))
}

func (h *Handler) getLog(c echo.Context) (e error) {
	ctx := c.(*cuxs.Context)
	var mx []*model.StatisticLog
	params, _ := ctx.FormParams()

	if _, e = auth.UserSession(ctx); e == nil {
		if mx, e = getLog(params); e == nil {
			ctx.Data(mx)
		}
	}

	return ctx.Serve(e)
}
