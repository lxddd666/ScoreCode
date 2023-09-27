// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package etcd implements service Registry and Discovery using etcd.
package etcd

import (
	"time"

	etcd3 "go.etcd.io/etcd/client/v3"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gsvc"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
)

var (
	_ gsvc.Registry = &Registry{}
)

// Registry implements gsvc.Registry interface.
type Registry struct {
	client       *etcd3.Client
	kv           etcd3.KV
	lease        etcd3.Lease
	keepaliveTTL time.Duration
	logger       glog.ILogger
}

// Option is the option for the etcd registry.
type Option struct {
	Logger       glog.ILogger
	KeepaliveTTL time.Duration
}

const (
	// DefaultKeepAliveTTL is the default keepalive TTL.
	DefaultKeepAliveTTL = 10 * time.Second
)

// New creates and returns a new etcd registry.
func New(address string, option ...Option) gsvc.Registry {
	endpoints := gstr.SplitAndTrim(address, ",")
	if len(endpoints) == 0 {
		panic(gerror.NewCodef(gcode.CodeInvalidParameter, `invalid etcd address "%s"`, address))
	}
	client, err := etcd3.New(etcd3.Config{
		Endpoints: endpoints,
	})
	if err != nil {
		panic(gerror.Wrap(err, `create etcd client failed`))
	}
	return NewWithClient(client, option...)
}

// NewWithClient creates and returns a new etcd registry with the given client.
func NewWithClient(client *etcd3.Client, option ...Option) *Registry {
	r := &Registry{
		client: client,
		kv:     etcd3.NewKV(client),
	}
	if len(option) > 0 {
		r.logger = option[0].Logger
		r.keepaliveTTL = option[0].KeepaliveTTL
	}
	if r.logger == nil {
		r.logger = g.Log()
	}
	if r.keepaliveTTL == 0 {
		r.keepaliveTTL = DefaultKeepAliveTTL
	}
	return r
}

// extractResponseToServices extracts etcd watch response context to service list.
func extractResponseToServices(res *etcd3.GetResponse) ([]gsvc.Service, error) {
	if res == nil || res.Kvs == nil {
		return nil, nil
	}
	var (
		services         []gsvc.Service
		servicePrefixMap = make(map[string]*Service)
	)
	for _, kv := range res.Kvs {
		service, err := gsvc.NewServiceWithKV(
			string(kv.Key), "",
		)
		if err != nil {
			return services, err
		}
		s := NewService(service)
		if v, ok := servicePrefixMap[service.GetPrefix()]; ok {
			v.Endpoints = append(v.Endpoints, service.GetEndpoints()...)
		} else {
			servicePrefixMap[s.GetPrefix()] = s
			services = append(services, s)
		}
	}
	return services, nil
}
