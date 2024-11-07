package rediss

import (
	"MyIM/pkg/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
)

var onceDefault sync.Once
var defaultClient *redis.ClusterClient

var onceUserList sync.Once
var userListClient *redis.ClusterClient

func Init() {
	remoteCache := config.GetString("remote_cache")
	if remoteCache == "redis" {
		//initRedis()
	}
	if remoteCache == "redis_cluster" {
		initRedisCluster()
	}
}

//func initRedis() {
//	onceDefault.Do(func() {
//		defaultClient = redis.NewClient(&redis.Options{
//			Addr:     config.GetString("redis.default.addr"),
//			Password: config.GetString("redis.default.password"), // no password set
//			DB:       config.GetInt("redis.default.db"),          // use default DB
//			PoolSize: config.GetInt("redis.default.pool"),
//		})
//
//		pong, err := defaultClient.Ping(context.Background()).Result()
//		if err == nil {
//			log.Printf("\033[1;30;42m[info]\033[0m default redis connect success %s\n", pong)
//		} else {
//			panic(fmt.Sprintf("\033[1;30;41m[error]\033[0m default redis connect error %s\n", err.Error()))
//		}
//	})
//	onceUserList.Do(func() {
//		userListClient = redis.NewClient(&redis.Options{
//			Addr:     config.GetString("redis.user_list.addr"),
//			Password: config.GetString("redis.user_list.password"), // no password set
//			DB:       config.GetInt("redis.user_list.db"),          // use default DB
//			PoolSize: config.GetInt("redis.user_list.pool"),
//		})
//		pong, err := userListClient.Ping(context.Background()).Result()
//		if err == nil {
//			log.Printf("\033[1;30;42m[info]\033[0m user list redis connect success %s\n", pong)
//		} else {
//			panic(fmt.Sprintf("\033[1;30;41m[error]\033[0m  user list  redis connect error %s\n", err.Error()))
//		}
//	})
//}

func initRedisCluster() {
	onceDefault.Do(func() {
		//addrs := config.GetStringSlice("redis_cluster.default.addrs")
		defaultClient = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    []string{"192.168.2.50:6379", "192.168.2.51:6379", "192.168.2.52:6379"},
			Password: config.GetString("redis_cluster.default.password"), // no password set
			PoolSize: config.GetInt("redis_cluster.default.pool"),
		})

		pong, err := defaultClient.Ping(context.Background()).Result()
		if err == nil {
			log.Printf("\033[1;30;42m[info]\033[0m default redis connect success %s\n", pong)
		} else {
			panic(fmt.Sprintf("\033[1;30;41m[error]\033[0m default redis connect error %s\n", err.Error()))
		}
	})
	onceUserList.Do(func() {
		userListClient = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    config.GetStringSlice("redis_cluster.user_list.addrs"),
			Password: config.GetString("redis_cluster.user_list.password"), // no password set
			PoolSize: config.GetInt("redis_cluster.user_list.pool"),
		})
		pong, err := userListClient.Ping(context.Background()).Result()
		if err == nil {
			log.Printf("\033[1;30;42m[info]\033[0m user list redis connect success %s\n", pong)
		} else {
			panic(fmt.Sprintf("\033[1;30;41m[error]\033[0m  user list  redis connect error %s\n", err.Error()))
		}
	})
}
func Default() *redis.ClusterClient {
	return defaultClient
}

func UserList() *redis.ClusterClient {
	return userListClient
}
