package captcha

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

func (s *CacheStore) Set(key string, value string) error {
	expireTime := time.Minute * time.Duration(config.Get[int]("captcha.expire_time"))

	if app.IsLocal() {
		expireTime = time.Minute * time.Duration(config.Get[int]("captcha.debug_expire_time"))
	}

	cache.Set(s.KeyPrefix+key, value, expireTime)

	return nil
}

func (s *CacheStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	val := cache.Get(key)
	if clear {
		cache.Forget(key)
	}
	return cast.ToString(val)
}

func (s *CacheStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
