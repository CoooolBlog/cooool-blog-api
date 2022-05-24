/*
 * Copyright 2022 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kristianhuang/go-cmp/errors"
	"golang.org/x/time/rate"
)

// ErrLimitExceeded defines Limit exceeded error.
var ErrLimitExceeded = errors.New("Limit exceeded")

// Limit drops (HTTP status 429) the request if the limit is reached.
func Limit(maxEventsPerSec float64, maxBurstSize int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(maxEventsPerSec), maxBurstSize)

	return func(c *gin.Context) {
		if limiter.Allow() {
			c.Next()

			return
		}

		// Limit reached
		_ = c.Error(ErrLimitExceeded)
		c.AbortWithStatus(429)
	}
}
