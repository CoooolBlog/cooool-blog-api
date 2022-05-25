/*
 * Copyright 2022 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/CoooolBlog/cooool-blog-api/internal/authzserver"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	authzserver.NewApp("cooool-blog-authz-server").Run()
}
