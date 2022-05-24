/*
 * Copyright 2022 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package authzserver

import (
	"github.com/gin-gonic/gin"
)

type Route func(e *gin.Engine)

var (
	routes = []Route{}
)

func installMiddleware(e *gin.Engine) {

}

func initRouter(e *gin.Engine) *gin.Engine {
	installMiddleware(e)

	for _, r := range routes {
		r(e)
	}

	return e
}
