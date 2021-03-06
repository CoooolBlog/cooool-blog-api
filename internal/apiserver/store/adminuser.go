/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package store

import (
	"context"

	v1 "cooool-blog-api/internal/pkg/model"

	metav1 "github.com/kristianhuang/go-cmp/meta/v1"
)

type AdminUserStore interface {
	Create(ctx context.Context, adminUserModel *v1.AdminUser, opts metav1.CreateOptions) error

	Delete(ctx context.Context, account string, opts metav1.DeleteOptions) error
}
