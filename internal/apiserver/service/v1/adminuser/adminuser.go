/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package adminuser

import (
	"context"

	"cooool-blog-api/internal/apiserver/store"
	"cooool-blog-api/internal/pkg/model"

	metav1 "github.com/kristianhuang/go-cmp/meta/v1"
)

type AdminUserService interface {
	Create(ctx context.Context, au *model.AdminUser, options metav1.CreateOptions) error

	Delete(ctx context.Context, account string, opts metav1.DeleteOptions) error
}

type adminUserService struct {
	store store.Factory
}

func NewAdminUserService(s store.Factory) *adminUserService {
	return &adminUserService{store: s}
}
