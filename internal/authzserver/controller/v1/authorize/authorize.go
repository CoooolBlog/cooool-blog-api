/*
 * Copyright 2022 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package authorize

import (
	"github.com/CoooolBlog/cooool-blog-api/internal/pkg/bind"
	"github.com/CoooolBlog/cooool-blog-api/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/ory/ladon"
)

type AuthzController struct {
}

func NewAuthzController() *AuthzController {
	return &AuthzController{}
}

func (a *AuthzController) Authorize(c *gin.Context) {
	var r ladon.Request

	if err := bind.Bind(c, &r); err != nil {
		response.Write(c, err, nil)
		return
	}

	if r.Context == nil {
		r.Context = ladon.Context{}
	}
	r.Context["username"] = c.GetString("username")

}
