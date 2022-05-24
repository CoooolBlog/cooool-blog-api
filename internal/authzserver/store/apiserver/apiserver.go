/*
 * Copyright 2022 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package apiserver

import (
	"sync"

	"cooool-blog-api/internal/authzserver/store"
	pb "cooool-blog-api/internal/pkg/proto/apiserver/v1"

	log "github.com/kristianhuang/go-cmp/rollinglog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type datastore struct {
	cli pb.CacheClient
}

func (ds *datastore) Secrets() store.SecretStore {
	return newSecrets(ds)
}

func (ds *datastore) Policies() store.PolicyStore {
	return newPolicies(ds)
}

var (
	apiServerFactory store.Factory
	once             sync.Once
)

// GetAPIServerFactoryOrDie return cache instance and panics on any error.
func GetAPIServerFactoryOrDie(address string, clientCA string) store.Factory {
	once.Do(func() {
		var (
			err   error
			conn  *grpc.ClientConn
			creds credentials.TransportCredentials
		)

		creds, err = credentials.NewClientTLSFromFile(clientCA, "")
		if err != nil {
			log.Panicf("credentials.NewClientTLSFromFile err: %v", err)
		}

		conn, err = grpc.Dial(address, grpc.WithBlock(), grpc.WithTransportCredentials(creds))
		if err != nil {
			log.Panicf("Connect to grpc server failed, error: %s", err.Error())
		}

		apiServerFactory = &datastore{pb.NewCacheClient(conn)}
		log.Infof("Connected to grpc server, address: %s", address)
	})

	if apiServerFactory == nil {
		log.Panicf("failed to get apiserver store fatory")
	}

	return apiServerFactory
}