// Copyright 2018 Kora ID. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rest

import (
	"github.com/labstack/echo"
	"github.com/naufaldinta13/cuxs/cuxs"
	"github.com/naufaldinta13/graphic-api/src/auth"
)

type Handler struct{}

func (h *Handler) URLMapping(r *echo.Group) {
	r.POST("/signin", h.signin)
	r.POST("/pasien", h.pasien, auth.CheckPrivilege(""))
	r.GET("/me", h.me, auth.CheckPrivilege(""))
}

func (h *Handler) signin(c echo.Context) (e error) {
	ctx := c.(*cuxs.Context)
	var sd *auth.SessionData
	var req signinRequest

	if e = c.Bind(&req); e == nil {
		if sd, e = auth.StartSession(req.UserData); e == nil {
			ctx.Data(sd)
		}
	}

	return ctx.Serve(e)
}

func (h *Handler) pasien(c echo.Context) (e error) {
	ctx := c.(*cuxs.Context)
	var req pasienRequest

	if req.Session, e = auth.UserSession(ctx); e == nil {
		if e = ctx.Bind(&req); e == nil {
			if req.Session, e = auth.StartSession(req.Session.User, req.Session.Token); e == nil {
				ctx.Data(req.Session)
			}
		}
	}

	return ctx.Serve(e)
}

func (h *Handler) me(c echo.Context) (e error) {
	ctx := c.(*cuxs.Context)
	var sd *auth.SessionData

	if sd, e = auth.UserSession(ctx); e == nil {
		ctx.Data(sd)
	}

	return ctx.Serve(e)
}
