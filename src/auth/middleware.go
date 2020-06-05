// Copyright 2016 PT. Qasico Teknologi Indonesia. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/labstack/echo"
	"github.com/naufaldinta13/cuxs/cuxs"
	"github.com/naufaldinta13/graphic-api/datastore/model"
)

// CheckPrivilege is middleware that will check if user has authorize in endpoint
// note: pada handler UrlMapping ditambahkan parameter auth.CheckPrivilege() dan application modulnya
// e.g:  r.Post("",h.auth, auth.CheckPrivilege("dashboard"))
func CheckPrivilege(module string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return cuxs.Authorized()(func(c echo.Context) error {
			ctx := cuxs.NewContext(c)
			if u := ctx.JwtUsers(&model.User{}); u != nil {
				return next(c)
			}
			return echo.ErrUnauthorized
		})
	}
}
