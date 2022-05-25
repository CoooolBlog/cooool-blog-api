/*
 * Copyright 2022 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package bind

import (
	"github.com/CoooolBlog/cooool-blog-api/internal/pkg/code"
	"github.com/gin-gonic/gin"
	"github.com/kristianhuang/go-cmp/errors"
	"github.com/kristianhuang/go-cmp/validator"
)

// BindJson bind json data and validate.
func BindJson(c *gin.Context, data interface{}) error {
	if err := c.ShouldBindJSON(data); err != nil {
		return errors.WithCode(code.ErrBind, err.Error())
	}

	if err := validate(data); err != nil {
		return err
	}

	return nil
}

// BindQuery bind query data and validate.
func BindQuery(c *gin.Context, data interface{}) error {
	if err := c.ShouldBindQuery(data); err != nil {
		return errors.WithCode(code.ErrBind, err.Error())
	}

	if err := validate(data); err != nil {
		return err
	}

	return nil
}

// BindUri bind uri data and validate.
func BindUri(c *gin.Context, data interface{}) error {
	if err := c.ShouldBindUri(data); err != nil {
		return errors.WithCode(code.ErrBind, err.Error())
	}

	if err := validate(data); err != nil {
		return err
	}

	return nil
}

// Bind bind data and validate.
func Bind(c *gin.Context, data interface{}) error {
	if err := c.ShouldBind(data); err != nil {
		return errors.WithCode(code.ErrBind, err.Error())
	}

	if err := validate(data); err != nil {
		return err
	}

	return nil
}

func validate(data interface{}) error {
	if err := validator.Struct(data); err != nil {
		return errors.WithCode(code.ErrValidation, err.(*validator.ValidationErrors).TranslateErrs()[0].Error())
	}

	return nil
}
