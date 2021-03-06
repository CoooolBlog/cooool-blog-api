/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package adminuser

import (
	"context"

	"cooool-blog-api/internal/pkg/model"

	metav1 "github.com/kristianhuang/go-cmp/meta/v1"
)

func (a *adminUserService) Create(ctx context.Context, au *model.AdminUser, options metav1.CreateOptions) error {
	if err := a.store.AdminUser().Create(ctx, au, options); err != nil {
		return err
	}

	return nil
}
