/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package adminuser

import (
	"github.com/CoooolBlog/cooool-blog-api/internal/pkg/bind"
	"github.com/CoooolBlog/cooool-blog-api/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type delForm struct {
	Account string `json:"account" validate:"required,gt=6" form:"account" uri:"account"`
}

func (a *AdminUserController) Delete(c *gin.Context) {
	delForm := &delForm{}
	if err := bind.Bind(c, delForm); err != nil {
		response.Write(c, err, nil)

		return
	}

	response.Write(c, nil, nil)
}
