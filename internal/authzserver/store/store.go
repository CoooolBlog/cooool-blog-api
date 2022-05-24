/*
 * Copyright 2022 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package store

//go:generate mockgen -self_package=cooool-blog/internal/authzserver/store -destination mock_store.go -package store cooool-blog/iam/internal/authzserver/store Factory,SecretStore,PolicyStore

var client Factory

// Factory defines the iam platform storage interface.
type Factory interface {
	Policies() PolicyStore
	Secrets() SecretStore
}

// Client return the store client instance.
func Client() Factory {
	return client
}

// SetClient set the iam store client.
func SetClient(factory Factory) {
	client = factory
}