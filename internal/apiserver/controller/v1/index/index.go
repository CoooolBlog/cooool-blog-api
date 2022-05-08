/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kristianhuang/go-cmp/rollinglog"
	"github.com/spf13/viper"
)

type Controller struct {
}

func (ic *Controller) Index(c *gin.Context) {
	list := viper.AllSettings()
	rollinglog.Info("hello")
	c.JSON(http.StatusOK, gin.H{"list": list})
}

func NewIndexController() *Controller {
	return &Controller{}
}
