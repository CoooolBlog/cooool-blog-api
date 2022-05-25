/*
 * Copyright 2022 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package apiserver

import (
	"github.com/CoooolBlog/cooool-blog-api/internal/apiserver/route"

	"github.com/gin-gonic/gin"
)

type Route func(e *gin.Engine)

var (
	routes = []Route{
		route.Index,
		route.AdminUser,
	}
)

func installMiddleware(e *gin.Engine) {

}

func installRoutes(e *gin.Engine) {
	for _, r := range routes {
		r(e)
	}
}

func initRouter(e *gin.Engine) *gin.Engine {
	installMiddleware(e)
	installRoutes(e)

	return e
}
