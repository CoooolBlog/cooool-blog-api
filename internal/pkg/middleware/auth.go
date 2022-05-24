/*
 * Copyright 2022 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package middleware

import "github.com/gin-gonic/gin"

// AuthStrategy defines the set of methods used to do resource authentication.
type AuthStrategy interface {
	AuthFunc() gin.HandlerFunc
}

// AuthOperator used to switch between different authentication strategy.
type AuthOperator struct {
	strategy AuthStrategy
}

// SetStrategy used to set to another authentication strategy.
func (o *AuthOperator) SetStrategy(strategy AuthStrategy) {
	o.strategy = strategy
}

// AuthFunc execute resource authentication.
func (o *AuthOperator) AuthFunc() gin.HandlerFunc {
	return o.strategy.AuthFunc()
}
