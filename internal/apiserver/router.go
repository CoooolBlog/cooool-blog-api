/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package apiserver

import (
	"cooool-blog-api/internal/apiserver/route"

	"github.com/gin-gonic/gin"
)

type Route func(e *gin.Engine)

var (
	routes = []Route{
		route.Index,
		route.AdminUser,
	}
)

func initRouter(e *gin.Engine) *gin.Engine {
	for _, r := range routes {
		r(e)
	}

	return e
}
