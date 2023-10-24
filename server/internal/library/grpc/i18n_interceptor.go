package grpc

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/library/contexts"
)

// customCredential 自定义认证
type customCredential struct {
	ctx context.Context
}

func NewCustomCredential(ctx context.Context) *customCredential {
	return &customCredential{ctx: ctx}
}

func (c *customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		consts.HttpLanguage: gconv.String(contexts.GetData(c.ctx)[consts.HttpLanguage]),
	}, nil
}

func (c *customCredential) RequireTransportSecurity() bool {
	return false
}
