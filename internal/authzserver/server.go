/*
 * Copyright 2022 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package authzserver

import (
	"context"

	"cooool-blog-api/internal/authzserver/analytics"
	"cooool-blog-api/internal/authzserver/config"
	"cooool-blog-api/internal/authzserver/load"
	"cooool-blog-api/internal/authzserver/load/cache"
	"cooool-blog-api/internal/authzserver/store/apiserver"
	genericoptions "cooool-blog-api/internal/pkg/options"
	genericapiserver "cooool-blog-api/internal/pkg/server"

	"github.com/kristianhuang/go-cmp/errors"
	log "github.com/kristianhuang/go-cmp/rollinglog"
	"github.com/kristianhuang/go-cmp/shutdown"
	"github.com/kristianhuang/go-cmp/shutdown/shutdownmanagers/posixsignal"
	"github.com/kristianhuang/go-cmp/storage"
)

// RedisKeyPrefix defines the prefix key in redis for analytics data.
const RedisKeyPrefix = "analytics-"

type authzServer struct {
	gs               *shutdown.GracefulShutdown
	clientCA         string
	rpcServer        string
	redisOptions     *genericoptions.RedisOptions
	genericAPIServer *genericapiserver.GenericAPIServer
	analyticsOptions *analytics.AnalyticsOptions
	redisCancelFunc  context.CancelFunc
}

type prepareAuthzServer struct {
	*authzServer
}

func createAuthzServer(conf *config.Config) (*authzServer, error) {
	gs := shutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	genericConfig, err := buildGenericConfig(conf)
	if err != nil {
		return nil, err
	}

	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}

	server := &authzServer{
		gs:               gs,
		redisOptions:     conf.RedisOptions,
		genericAPIServer: genericServer,
		analyticsOptions: conf.AnalyticsOptions,
	}

	return server, nil
}

func (a *authzServer) BeforeRun() prepareAuthzServer {
	_ = a.initialize()
	initRouter(a.genericAPIServer.Engine)

	return prepareAuthzServer{a}
}

func (s prepareAuthzServer) Run() error {
	stopCh := make(chan struct{})

	// start shutdown managers
	if err := s.gs.Start(); err != nil {
		log.Fatalf("start shutdown manager failed: %s", err.Error())
	}

	// in order to ensure that the reported data is not lost,
	// please ensure the following graceful shutdown sequence.
	s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		s.genericAPIServer.Close()
		if s.analyticsOptions.Enable {
			analytics.GetAnalytics().Stop()
		}
		s.redisCancelFunc()

		return nil
	}))

	<-stopCh

	return nil
}

func buildGenericConfig(cfg *config.Config) (genericConfig *genericapiserver.Config, lastErr error) {
	genericConfig = genericapiserver.NewConfig()
	if lastErr = cfg.GenericServerOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	if lastErr = cfg.FeatureOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	// TODO https

	if lastErr = cfg.InsecureServingOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	return
}

func (a *authzServer) buildStorageConfig() *storage.Config {
	return &storage.Config{
		Host:                  a.redisOptions.Host,
		Port:                  a.redisOptions.Port,
		Addrs:                 a.redisOptions.Addrs,
		MasterName:            a.redisOptions.MasterName,
		Username:              a.redisOptions.Username,
		Password:              a.redisOptions.Password,
		Database:              a.redisOptions.Database,
		MaxIdle:               a.redisOptions.MaxIdle,
		MaxActive:             a.redisOptions.MaxActive,
		Timeout:               a.redisOptions.Timeout,
		EnableCluster:         a.redisOptions.EnableCluster,
		UseSSL:                a.redisOptions.UseSSL,
		SSLInsecureSkipVerify: a.redisOptions.SSLInsecureSkipVerify,
	}
}

func (a *authzServer) initialize() error {
	ctx, cancel := context.WithCancel(context.Background())
	a.redisCancelFunc = cancel

	// keep redis connected
	go storage.ConnectToRedis(ctx, a.buildStorageConfig())

	// cron to reload all secrets and policies from iam-apiserver
	cacheIns, err := cache.GetCacheInsOr(apiserver.GetAPIServerFactoryOrDie(a.rpcServer, a.clientCA))
	if err != nil {
		return errors.Wrap(err, "get cache instance failed")
	}

	load.NewLoader(ctx, cacheIns).Start()

	// start analytics service
	if a.analyticsOptions.Enable {
		analyticsStore := storage.RedisCluster{KeyPrefix: RedisKeyPrefix}
		analyticsIns := analytics.NewAnalytics(a.analyticsOptions, &analyticsStore)
		analyticsIns.Start()
	}

	return nil
}
