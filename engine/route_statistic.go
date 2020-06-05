// Copyright 2018 Kora ID. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package engine

import "github.com/naufaldinta13/graphic-api/src/statistic/rest"

func init() {
	handlers["statistic"] = &rest.Handler{}
}
