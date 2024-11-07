package cache

import (
	"errors"
	"time"

	cache "github.com/patrickmn/go-cache"
)

// Cache 缓存封装对象，对go cache的操作加强管控，杜绝无过期时间的k-v缓存写入
type goCache struct {
	*cache.Cache
}

// New 代理go cache New，增加默认过期时间和过期key清除间隔时间判断，如果设置非法则返回错误
func New(defaultExpiration, cleanupInterval time.Duration) (*goCache, error) {
	if defaultExpiration <= 0 || cleanupInterval <= 0 {
		return nil, errors.New("defaultExpiration or cleanupInterval not set")
	}
	return &goCache{cache.New(defaultExpiration, cleanupInterval)}, nil
}

// Set 代理go cache Set，增加过期时间校验，过期时间为0返回错误
func (cache *goCache) Set(k string, x interface{}, d time.Duration) error {
	if d <= 0 {
		return errors.New("expire duration not set")
	}
	cache.Cache.Set(k, x, d)
	return nil
}

// Add 代理go cache Add方法，增加过期时间校验，过期时间为0返回错误
func (cache *goCache) Add(k string, x interface{}, d time.Duration) error {
	if d <= 0 {
		return errors.New("expire duration not set")
	}
	return cache.Cache.Add(k, x, d)
}

// SetDefault 屏蔽SetDefault方法，禁止缓存无过期时间
func (cache *goCache) SetDefault(k string, x interface{}) {
	// forbidden
	return
}
