/*
 * Copyright 2022 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package apiserver

import (
	"cooool-blog-api/internal/apiserver/config"
	"cooool-blog-api/internal/apiserver/store"
	"cooool-blog-api/internal/apiserver/store/mysql"
	genericoptions "cooool-blog-api/internal/pkg/options"
	genericapiserver "cooool-blog-api/internal/pkg/server"

	log "github.com/kristianhuang/go-cmp/rollinglog"
	"github.com/kristianhuang/go-cmp/shutdown"
	"github.com/kristianhuang/go-cmp/shutdown/shutdownmanagers/posixsignal"
	"github.com/kristianhuang/go-cmp/validator"
)

type apiServer struct {
	gs            *shutdown.GracefulShutdown
	genericServer *genericapiserver.GenericAPIServer
	redisOptions  *genericoptions.RedisOptions
}

type preparedAPIServer struct {
	*apiServer
}

// 使用 apiserver 的配置项填充至 server 的配置项，用以满足启动 server 的必要条件。
func buildGenericConfig(conf *config.Config) (apiServerConfig *genericapiserver.Config, err error) {
	apiServerConfig = genericapiserver.NewConfig()
	if err = conf.GenericServerOptions.ApplyTo(apiServerConfig); err != nil {
		return
	}

	if err = conf.FeatureOptions.ApplyTo(apiServerConfig); err != nil {
		return
	}

	if err = conf.InsecureServingOptions.ApplyTo(apiServerConfig); err != nil {
		return
	}

	return
}

func createServer(config *config.Config) (*apiServer, error) {
	gs := shutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	genericServerConfig, err := buildGenericConfig(config)
	if err != nil {
		return nil, err
	}

	genericAPIServer, err := genericServerConfig.Complete().New()
	if err != nil {
		return nil, err
	}

	// init mysql store.
	storeIns, err := mysql.GetMysqlFactory(config.MySQLOptions)
	if err != nil {
		return nil, err
	}
	store.SetClient(storeIns)

	// init validator.
	validator.Init(config.Validator)

	server := &apiServer{
		gs:            gs,
		genericServer: genericAPIServer,
		redisOptions:  config.RedisOptions,
	}

	return server, nil
}

func (s *apiServer) BeforeRun() preparedAPIServer {
	// init redis
	s.initRedisStore()

	// init router
	initRouter(s.genericServer.Engine)

	s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		mysqlStore, _ := mysql.GetMysqlFactory(nil)
		if mysqlStore != nil {
			_ = mysqlStore.Close()
		}
		s.genericServer.Close()

		return nil
	}))

	return preparedAPIServer{s}
}

func (s preparedAPIServer) Run() error {

	// start shutdown managers.
	if err := s.gs.Start(); err != nil {
		log.Fatalf("start shutdown manager failed: %s", err.Error())
	}

	return s.genericServer.Run()
}
