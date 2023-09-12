package captcha

import (
	"context"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/mojocn/base64Captcha"
	"hotgo/internal/library/cache"
	"time"
)

type gStore struct {
	ctx        context.Context
	cache      *gcache.Cache
	expiration time.Duration
}

func NewGStore(ctx context.Context) base64Captcha.Store {

	return &gStore{ctx: ctx, expiration: base64Captcha.Expiration, cache: cache.Instance()}
}

// Set sets the digits for the captcha id.
func (s *gStore) Set(id string, value string) (err error) {
	err = s.cache.Set(s.ctx, id, value, s.expiration)
	return
}

// Get returns stored digits for the captcha id. Clear indicates
// whether the captcha must be deleted from the store.
func (s *gStore) Get(id string, clear bool) (val string) {
	res, err := s.cache.Get(s.ctx, id)
	if err != nil {
		return ""
	}
	if clear {
		_, _ = s.cache.Remove(s.ctx, id)
	}
	val = res.String()
	return
}

// Verify captcha answer directly
func (s *gStore) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v == answer
}
