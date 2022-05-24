/*
 * Copyright 2022 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package options

import (
	"cooool-blog-api/internal/authzserver/analytics"
	genericoptions "cooool-blog-api/internal/pkg/options"
	"cooool-blog-api/internal/pkg/server"

	"github.com/kristianhuang/go-cmp/flag"
	"github.com/kristianhuang/go-cmp/rollinglog"
)

type Options struct {
	RPCServer              string `json:"rpcserver" mapstructure:"rpcserver"`
	ClientCA               string `json:"client-ca-file" mapstructure:"client-ca-file"`
	GenericServerOptions   *genericoptions.GenericServerOptions
	InsecureServingOptions *genericoptions.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	FeatureOptions         *genericoptions.FeatureOptions         `json:"feature" mapstructure:"feature"`
	RedisOptions           *genericoptions.RedisOptions           `json:"redis" mapstructure:"redis"`
	Log                    *rollinglog.Options                    `json:"log" mapstructure:"log"`
	AnalyticsOptions       *analytics.AnalyticsOptions            `json:"analytics" mapstructure:"analytics"`
}

func NewOptions() *Options {
	return &Options{
		RPCServer:              "127.0.0.1:8081",
		ClientCA:               "",
		GenericServerOptions:   genericoptions.NewGenericServerOptions(),
		InsecureServingOptions: genericoptions.NewInsecureServingOptions(),
		FeatureOptions:         genericoptions.NewFeatureOptions(),
		RedisOptions:           genericoptions.NewRedisOptions(),
		Log:                    rollinglog.NewOptions(),
		AnalyticsOptions:       analytics.NewAnalyticsOptions(),
	}
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *server.Config) error {
	return nil
}

// Flags returns flags for a specific APIServer by section name.
func (o *Options) Flags() (fss flag.NamedFlagSets) {
	o.GenericServerOptions.AddFlags(fss.FlagSet("generic"))
	o.AnalyticsOptions.AddFlags(fss.FlagSet("analytics"))
	o.RedisOptions.AddFlags(fss.FlagSet("redis"))
	o.FeatureOptions.AddFlags(fss.FlagSet("features"))
	o.InsecureServingOptions.AddFlags(fss.FlagSet("insecure serving"))
	// TODO https
	o.Log.AddFlags(fss.FlagSet("logs"))

	// Note: the weird ""+ in below lines seems to be the only way to get gofmt to
	// arrange these text blocks sensibly. Grrr.
	fs := fss.FlagSet("misc")
	fs.StringVar(&o.RPCServer, "rpcserver", o.RPCServer, "The address of cooool-blog rpc server. "+
		"The rpc server can provide all the secrets and policies to use.")
	fs.StringVar(&o.ClientCA, "client-ca-file", o.ClientCA, ""+
		"If set, any request presenting a client certificate signed by one of "+
		"the authorities in the client-ca-file is authenticated with an identity "+
		"corresponding to the CommonName of the client certificate.")

	return fss
}
