package cache

import (
	"time"

	mcpkg "github.com/patrickmn/go-cache"
	"github.com/spf13/cast"
)

type MemoryStore struct {
	Client *mcpkg.Cache
}

func NewMemoryStore() *MemoryStore {
	return NewMemoryStoreWithExpiration(5*time.Minute, 10*time.Minute)
}

func NewMemoryStoreWithExpiration(defaultExpiration time.Duration, cleanupInterval time.Duration) *MemoryStore {
	ms := &MemoryStore{}
	ms.Client = mcpkg.New(defaultExpiration, cleanupInterval)
	return ms
}

func (ms *MemoryStore) Set(key string, value string, expireTime time.Duration) {
	ms.Client.Set(key, value, expireTime)
}

func (ms *MemoryStore) Get(key string) string {
	value, _ := ms.Client.Get(key)
	return cast.ToString(value)
}

func (ms *MemoryStore) Has(key string) bool {
	_, found := ms.Client.Get(key)
	return found
}

func (ms *MemoryStore) Forget(key string) {
	ms.Client.Delete(key)
}

func (ms *MemoryStore) Forever(key string, value string) {
	ms.Client.Set(key, value, mcpkg.NoExpiration)
}

func (ms *MemoryStore) Flush() {
	ms.Client.Flush()
}

func (ms *MemoryStore) IsAlive() error {
	return nil
}

func (ms *MemoryStore) Increment(parameters ...interface{}) {
	key := cast.ToString(parameters[0])
	if len(parameters) > 1 {
		ms.Client.Increment(key, cast.ToInt64(parameters[1]))
	} else {
		ms.Client.Increment(key, 1)
	}
}

func (ms *MemoryStore) Decrement(parameters ...interface{}) {
	key := cast.ToString(parameters[0])
	if len(parameters) > 1 {
		ms.Client.Decrement(key, cast.ToInt64(parameters[1]))
	} else {
		ms.Client.Decrement(key, 1)
	}
}
