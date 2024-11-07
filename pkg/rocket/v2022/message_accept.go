package rocket

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/luantao/IM-base/pkg/http"
	gohttp "net/http"
	"os"
	"strings"
	"sync"

	acm "github.com/luantao/IM-base/pkg/config"
	log "github.com/luantao/IM-base/pkg/mlog"
	"github.com/luantao/IM-base/pkg/rocketmq-client-go/primitive"

	"go.uber.org/zap"
)

var (
	once sync.Once
	cmap *callMap
)

type callMap struct {
	sync.RWMutex
	data map[string]PushConsumerFn
}

func (m *callMap) set(key string, value PushConsumerFn) {
	m.Lock()
	defer m.Unlock()
	m.data[key] = value
}

func (m *callMap) get(key string) PushConsumerFn {
	m.RLock()
	defer m.RUnlock()
	return m.data[key]
}

func MessageAccept(ctx context.Context, configNodeName string, logger *log.Mlog, fn PushConsumerFn) error {
	accept := &Accept{ctx, logger, configNodeName, ""}

	once.Do(func() {
		cmap = &callMap{data: make(map[string]PushConsumerFn, 1)}
		go func() {
			err := accept.serve()
			if err != nil {
				panic(err)
			}
		}()
	})

	if err := register(ctx, configNodeName, logger); err == nil {
		cmap.set(acm.GetString(configNodeName+".instance_name"), fn)
	} else {
		return err
	}
	return nil
}

func isTest() bool {
	if os.Getenv("ENV") != "" {
		return true
	}
	return false
}

func register(ctx context.Context, configNodeName string, logger *log.Mlog) error {
	env := os.Getenv("ENV")
	if env == "" {
		panic("environment variable env was not set")
	}

	req := http.Client(ctx).Request()
	req.FormData.Set("register_environment", env)
	req.FormData.Set("register_version", os.Getenv("VER"))
	req.FormData.Set("register_config_node_name", configNodeName)
	resp, err := req.Post(acm.GetString("register.register_url"))
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return errors.New(resp.String())
	}

	return nil
}

type Accept struct {
	context        context.Context
	logger         *log.Mlog
	configNodeName string
	message        string
}

func (accept *Accept) ServeHTTP(resp gohttp.ResponseWriter, req *gohttp.Request) {
	message := req.FormValue("msg")
	properties := req.FormValue("properties")
	//consumer 传递环境标识
	accept.context = context.WithValue(accept.context, "env", req.Header.Get("env"))
	accept.context = context.WithValue(accept.context, "ver", req.Header.Get("ver"))
	if strings.Trim(message, " ") != "" {
		err := accept.callback(message, properties)
		if err != nil {
			err = accept.callback(message, properties)
			if err != nil {
				accept.logger.Named("MessageAccept").Error("send http callback error",
					zap.Error(err),
					zap.String("message", message))
			}
		}
	}
	resp.WriteHeader(200)
	resp.Write(bytes.NewBufferString("serve http successfully").Bytes())
	return
}

func (accept *Accept) callback(msg, properties string) error {
	message := &primitive.MessageExt{}
	err := json.Unmarshal([]byte(msg), message)
	if err != nil {
		return err
	}
	message.UnmarshalProperties([]byte(properties))
	instance := message.GetProperty("instance")
	fn := cmap.get(instance)
	_, err = fn(accept.context, message)
	return err
}

func (accept *Accept) serve() error {
	return gohttp.ListenAndServe(":"+acm.GetString("register.accept_http_port"), accept)
}
