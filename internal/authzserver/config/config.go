/*
 * Copyright 2022 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package config

import "cooool-blog-api/internal/authzserver/options"

type Config struct {
	*options.Options
}

func NewConfig(options *options.Options) (*Config, error) {
	return &Config{Options: options}, nil
}
