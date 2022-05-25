/*
 * Copyright 2022 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package authzserver

import (
	"github.com/CoooolBlog/cooool-blog-api/internal/authzserver/config"
	"github.com/CoooolBlog/cooool-blog-api/internal/authzserver/options"

	"github.com/kristianhuang/go-cmp/app"
	log "github.com/kristianhuang/go-cmp/rollinglog"
)

const commandDesc = `Authorization server to run ladon policies which can protecting your resources.
Find more ladon information at:
    https://github.com/ory/ladon`

func NewApp(use string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp(use, "Authorization Server",
		app.WithLong(commandDesc),
		app.WithOptions(opts),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(createRunFunc(opts)),
	)

	return application
}

func createRunFunc(opts *options.Options) app.RunFunc {
	return func(use string) error {
		log.Init(opts.Log)
		defer log.Flush()

		conf, err := config.NewConfig(opts)
		if err != nil {
			return err
		}

		return Run(conf)
	}
}

func Run(conf *config.Config) error {
	server, err := createAuthzServer(conf)
	if err != nil {
		return err
	}

	return server.BeforeRun().Run()
}
