/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package route

import (
	"cooool-blog-api/internal/apiserver/controller/v1/adminuser"

	"github.com/gin-gonic/gin"
)

func AdminUser(e *gin.Engine) {
	v1 := e.Group(V1)
	{
		adminUserController := adminuser.NewController()
		adminUserV1 := v1.Group("/admin_user")
		{
			adminUserV1.POST("", adminUserController.Create)
			adminUserV1.DELETE("", adminUserController.Delete)
		}
	}
}
