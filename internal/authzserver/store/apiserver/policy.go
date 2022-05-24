/*
 * Copyright 2022 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package apiserver

import (
	"context"
	"encoding/json"

	pb "cooool-blog-api/internal/pkg/proto/apiserver/v1"

	"github.com/AlekSi/pointer"
	"github.com/avast/retry-go"
	"github.com/kristianhuang/go-cmp/errors"
	log "github.com/kristianhuang/go-cmp/rollinglog"
	"github.com/ory/ladon"
)

type policies struct {
	cli pb.CacheClient
}

func newPolicies(ds *datastore) *policies {
	return &policies{ds.cli}
}

// List returns all the authorization policies.
func (p *policies) List() (map[string][]*ladon.DefaultPolicy, error) {
	pols := make(map[string][]*ladon.DefaultPolicy)

	log.Info("Loading policies")

	req := &pb.ListPoliciesRequest{
		Offset: pointer.ToInt64(0),
		Limit:  pointer.ToInt64(-1),
	}

	var resp *pb.ListPoliciesResponse
	err := retry.Do(
		func() error {
			var listErr error
			resp, listErr = p.cli.ListPolicies(context.Background(), req)
			if listErr != nil {
				return listErr
			}

			return nil
		}, retry.Attempts(3),
	)
	if err != nil {
		return nil, errors.Wrap(err, "list policies failed")
	}

	log.Infof("Policies found (%d total)[username:name]:", len(resp.Items))

	for _, v := range resp.Items {
		log.Infof(" - %s:%s", v.Username, v.Name)

		var policy ladon.DefaultPolicy

		if err := json.Unmarshal([]byte(v.PolicyShadow), &policy); err != nil {
			log.Warnf("failed to load policy for %s, error: %s", v.Name, err.Error())

			continue
		}

		pols[v.Username] = append(pols[v.Username], &policy)
	}

	return pols, nil
}
