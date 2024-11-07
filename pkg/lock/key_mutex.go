package lock

import (
	"MyIM/pkg/cache"
	"MyIM/pkg/config"
	"MyIM/pkg/mlog"
	gocache "github.com/patrickmn/go-cache"
	"go.uber.org/zap"
	"sync"
	"time"
)

var once sync.Once
var lockCache *gocache.Cache

// Init 初始化
func Init() {
	once.Do(func() {
		c, err := cache.New(config.GetDuration("cache.lock_expiration")*time.Millisecond, config.GetDuration("cache.lock_cleanup_interval")*time.Millisecond)
		if err != nil {
			mlog.Logger().Error("初始化本地锁缓存失败", zap.Error(err))
			panic(err)
		}
		lockCache = c.Cache
	})
}

// KeyMutexClient 获取锁
var KeyMutexClient = &keyMutex{}

// 缓存锁结构体
type keyMutex struct {
	mu sync.Mutex
}

// Locker 通过key获取锁
func (s *keyMutex) Locker(key string, timeout time.Duration) (mtx sync.Locker) {
	s.mu.Lock()
	defer s.mu.Unlock()
	mtxP, ok := lockCache.Get(key)
	if !ok {
		mtx = &sync.Mutex{}
		if timeout == 0 {
			timeout = config.GetDuration("cache.lock_expiration") * time.Millisecond
		}
		lockCache.Set(key, mtx, timeout)
		return
	}
	mtx = mtxP.(*sync.Mutex)
	return
}
