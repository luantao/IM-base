package rocket

import (
	"context"
	"os"
	"strings"

	rocketmq "github.com/luantao/IM-base/pkg/rocketmq-client-go"
	"github.com/luantao/IM-base/pkg/rocketmq-client-go/primitive"
	rocket_producer "github.com/luantao/IM-base/pkg/rocketmq-client-go/producer"

	"github.com/spf13/cast"
)

type Producer struct {
	rocketmq.Producer
}

func NewProducer(opts ...rocket_producer.Option) (rocketmq.Producer, error) {
	var err error
	producer := new(Producer)
	producer.Producer, err = rocketmq.NewProducer(opts...)
	return producer, err
}

func (producer *Producer) Start() error {
	return producer.Producer.Start()
}

func (producer *Producer) Shutdown() error {
	return producer.Producer.Shutdown()
}

func (producer *Producer) SendSync(ctx context.Context, mq ...*primitive.Message) (*primitive.SendResult, error) {
	mq = convertMessageFromContext(ctx, mq...)
	return producer.Producer.SendSync(ctx, mq...)
}
func (producer *Producer) SendAsync(ctx context.Context, mq func(ctx context.Context, result *primitive.SendResult, err error),
	msg ...*primitive.Message) error {
	msg = convertMessageFromContext(ctx, msg...)
	return producer.Producer.SendAsync(ctx, mq, msg...)
}

func (producer *Producer) SendOneWay(ctx context.Context, mq ...*primitive.Message) error {
	mq = convertMessageFromContext(ctx, mq...)
	return producer.Producer.SendOneWay(ctx, mq...)
}

func convertMessageFromContext(ctx context.Context, msg ...*primitive.Message) []*primitive.Message {
	env := cast.ToString(ctx.Value("env"))
	ver := cast.ToString(ctx.Value("ver"))

	if strings.Trim(env, " ") == "" {
		if strings.Trim(os.Getenv("ENV"), " ") == "" {
			return msg
		}
		env = os.Getenv("ENV")
		ver = os.Getenv("VER")
	}

	for i, message := range msg {
		msg[i] = buildMessage(env, ver, message)
	}
	return msg
}

func buildMessage(env, ver string, msg *primitive.Message) *primitive.Message {
	/* 使用message property传递环境版本标识，避免修改消息体
	m := Message{}
	m.Header = map[string]string{ENVIRONMENT: env, VERSION: ver}
	m.Body = string(msg.Body)
	body, err := json.Marshal(m)
	if err != nil {
		fmt.Println(fmt.Sprintf("marshal message err:%v", err))
		return msg
	}
	msg.Body = body
	*/
	msg.WithProperty(ENVIRONMENT, env)
	msg.WithProperty(VERSION, ver)
	return msg
}
