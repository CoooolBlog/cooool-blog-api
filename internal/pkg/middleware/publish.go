/*
 * Copyright 2022 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package middleware

import (
	"context"
	"net/http"
	"strings"

	"cooool-blog-api/internal/authzserver/load"

	"github.com/gin-gonic/gin"
	"github.com/kristianhuang/go-cmp/json"
	log "github.com/kristianhuang/go-cmp/rollinglog"
	"github.com/kristianhuang/go-cmp/storage"
)

func Publish() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if c.Writer.Status() != http.StatusOK {
			log.L(c).Debugf("request failed with http status code `%d`, ignore publish message", c.Writer.Status())

			return
		}

		var resource string

		pathSplit := strings.Split(c.Request.URL.Path, "/")
		if len(pathSplit) > 2 {
			resource = pathSplit[2]
		}

		method := c.Request.Method

		switch resource {
		case "policies":
			notify(c, method, load.NoticePolicyChanged)
		case "secrets":
			notify(c, method, load.NoticeSecretChanged)
		default:
		}
	}
}

func notify(ctx context.Context, method string, command load.NotificationCommand) {
	switch method {
	case "POST", "PUT", "DELETE", "PATH":
		redisStore := &storage.RedisCluster{}
		message, _ := json.Marshal(load.Notification{Command: command})

		if err := redisStore.Publish(load.RedisPubSubChannel, string(message)); err != nil {
			log.L(ctx).Errorw("publish redis message failed", "error", err.Error())
		}
		log.L(ctx).Debugw("publish redis message", "method", method, "command", command)
	default:
	}
}
