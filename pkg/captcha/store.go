package captcha

import (
	"context"
	"errors"
	"fmt"
	"github.com/luantao/IM-base/pkg/rediss"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

const StorePrefix = "permission:auth:image_captcha"

// customizeRdsStore An object implementing Store interface
type Store struct {
	client *redis.ClusterClient
}

func NewStore() *Store {
	return &Store{client: rediss.Default()}
}

// customizeRdsStore implementing Set method of  Store interface
func (s *Store) Set(id string, value string) {
	err := s.client.Set(context.Background(), fmt.Sprintf("%s:%s", StorePrefix, id), value, time.Minute*10).Err()
	if err != nil {
		log.Println(err)
	}
}

// customizeRdsStore implementing Get method of  Store interface
func (s *Store) Get(id string, clear bool) (value string) {
	key := fmt.Sprintf("%s:%s", StorePrefix, id)
	val, err := s.client.Get(context.Background(), key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		log.Println(err)
		return ""
	}
	if clear {
		err := s.client.Del(context.Background(), key).Err()
		if err != nil {
			log.Println(err)
			return ""
		}
	}
	return val
}
