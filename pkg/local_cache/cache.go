package local_cache

import (
	"errors"
	"github.com/patrickmn/go-cache"
	"time"
)

type Cache struct {
	*cache.Cache
}

// New 代理go cache New，增加默认过期时间和过期key清除间隔时间判断，如果设置非法则返回错误
func New(defaultExpiration, cleanupInterval time.Duration) (*Cache, error) {
	if defaultExpiration <= 0 || cleanupInterval <= 0 {
		return nil, errors.New("defaultExpiration or cleanupInterval not set")
	}
	return &Cache{Cache: cache.New(defaultExpiration, cleanupInterval)}, nil
}

// Set 代理go cache Set，增加过期时间校验，过期时间为0返回错误
func (cache *Cache) Set(k string, x interface{}, d time.Duration) error {
	if d <= 0 {
		return errors.New("expire duration not set")
	}
	cache.Cache.Set(k, x, d)
	return nil
}

// Add 代理go cache Add方法，增加过期时间校验，过期时间为0返回错误
func (cache *Cache) Add(k string, x interface{}, d time.Duration) error {
	if d <= 0 {
		return errors.New("expire duration not set")
	}
	return cache.Cache.Add(k, x, d)
}

// SetDefault 屏蔽SetDefault方法，禁止缓存无过期时间
func (cache *Cache) SetDefault(k string, x interface{}) {
	cache.Cache.SetDefault(k, x)
}
