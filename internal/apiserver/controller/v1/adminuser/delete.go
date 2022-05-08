/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package adminuser

import (
	"cooool-blog-api/internal/pkg/code"
	"cooool-blog-api/internal/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/kristianhuang/go-cmp/errors"
	"github.com/kristianhuang/go-cmp/validator"
)

type delForm struct {
	Account string `json:"account" validate:"required,gt=6" form:"account" uri:"account"`
}

func (a *AdminUserController) Delete(c *gin.Context) {
	delForm := &delForm{}
	if err := c.ShouldBind(delForm); err != nil {
		response.Write(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	if err := validator.Struct(delForm); err != nil {
		response.Write(c, errors.WithCode(code.ErrValidation, err.(*validator.ValidationErrors).TranslateErrs()[0].Error()), nil)

		return
	}

	response.Write(c, nil, nil)
}
