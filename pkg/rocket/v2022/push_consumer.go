package rocket

import (
	acm "MyIM/pkg/config"
	log "MyIM/pkg/mlog"
	"MyIM/pkg/rocketmq-client-go/consumer"
	"MyIM/pkg/rocketmq-client-go/primitive"
	"context"
	"github.com/spf13/cast"

	"go.uber.org/zap"
)

// PushConsumerFn 消费者逻辑
type PushConsumerFn func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error)

// NewPushConsumer 创建 PushConsumer
func NewPushConsumer(ctx context.Context, configNodeName string, logger *log.Mlog, fn PushConsumerFn) error {
	InitLogger(logger)

	if isTest() {
		logger.Named("NewPushConsumer").Info("new push consumer run with test mode")
		return MessageAccept(ctx, configNodeName, logger, fn)
	}

	c, err := consumer.NewPushConsumer(
		ConsumerOptionsWithACM(configNodeName)...,
	)
	if err != nil {
		logger.Named("NewPushConsumer").Error("new push consumer error",
			zap.Error(err),
			zap.String("key", configNodeName),
		)
		return err
	}

	var topicName = acm.GetString(configNodeName + ".topic_name")
	newTopic := cast.ToString(ctx.Value("topic"))
	if newTopic != "" {
		topicName = newTopic
	}
	var tags = acm.GetString(configNodeName + ".tags")
	newTags := cast.ToString(ctx.Value("tags"))
	if newTags != "" {
		tags = newTags
	}
	var selector consumer.MessageSelector
	if tags != "" {
		selector = consumer.MessageSelector{
			Type:       consumer.TAG,
			Expression: tags,
		}
	}

	err = c.Subscribe(topicName, selector, fn)
	if err != nil {
		logger.Named("NewPushConsumer").Error("new push consumer subscribe error",
			zap.Error(err),
			zap.String("key", configNodeName),
		)
		return err
	}

	if err := c.Start(); err != nil {
		logger.Named("NewPushConsumer").Error("new push consumer start error",
			zap.Error(err),
			zap.String("key", configNodeName),
		)
		return err
	}

	go func() {
		<-ctx.Done()
		err = c.Shutdown()
		if err != nil {
			logger.Named("NewPushConsumer").Error("push consumer shutdown error",
				zap.Error(err),
				zap.String("key", configNodeName),
			)
			return
		}
		logger.Named("NewPushConsumer").Info("push consumer shutdown",
			zap.String("key", configNodeName),
		)
	}()
	return nil
}
