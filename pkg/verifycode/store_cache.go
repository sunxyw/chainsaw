package verifycode

import (
	"gohub/pkg/app"
	"gohub/pkg/cache"
	"gohub/pkg/config"
	"time"

	"github.com/spf13/cast"
)

type CacheStore struct {
	KeyPrefix string
}

func (s *CacheStore) Set(id string, value string) {
	expireTime := time.Minute * time.Duration(config.Get[int64]("verifycode.expire_time"))

	if app.IsLocal() {
		expireTime = time.Minute * time.Duration(config.Get[int64]("verifycode.debug_expire_time"))
	}

	cache.Set(s.KeyPrefix+id, value, expireTime)
}

func (s *CacheStore) Get(id string, clear bool) string {
	key := s.KeyPrefix + id
	val := cache.Get(key)
	if clear {
		cache.Forget(key)
	}
	return cast.ToString(val)
}

func (s *CacheStore) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v == answer
}
