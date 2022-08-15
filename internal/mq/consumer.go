package mq

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-cloud/internal/cache"
)

type StreamGroup struct {
	Topic     string //消息队列
	GroupName string //消费者所在组名
	Consumer  string //消费者名
	start     string //0从头开始消费 $从最后一条消息开始消费
}

//消息处理
type ResponseMsg struct {
	ID       string
	Topic    string
	Group    string
	consumer string
	Body     map[string]interface{}
}
type MsgHandler func(msg *ResponseMsg) error

func NewStreamGroup(topic, groupName, consumer, start string) *StreamGroup {
	return &StreamGroup{
		Topic:     topic,
		GroupName: groupName,
		Consumer:  consumer,
		start:     start,
	}
}

//Consume count:一次性读取消息的个数，handler:消息处理函数
func (g *StreamGroup) Consume(c context.Context, count int64, handler MsgHandler) error {
	//创建消费组
	err := cache.XGroupCreate(c, g.Topic, g.GroupName, g.start)
	if err != nil {
		return err
	}
	//读取消费组中的消息
	for {
		if err := g.ReadGroup(c, count, ">", handler); err != nil {
			return err
		}
		if err := g.ReadGroup(c, count, "0", handler); err != nil {
			return err
		}
	}

}

/**
READGROUP GROUP group consumer [COUNT count] [BLOCK milliseconds] [NOACK] STREAMS key [key ...] ID [ID ...]
group ：消费组名
consumer ：消费者名。
count ： 读取数量。
milliseconds ： 阻塞毫秒数。
key ： 队列名。
ID ： 消息 ID  >:表示只接收未投递给其他消费者的消息;
这里先使用>拉取一次最新消息，再使用0拉取已经投递却没有ACK的消息，保证消息都能够成功消费。
*/
func (g *StreamGroup) ReadGroup(c context.Context, count int64, ID string, h MsgHandler) error {
	//阻塞的获取消息
	result, err := cache.XReadGroup(c, &redis.XReadGroupArgs{
		Group:    g.GroupName,
		Consumer: g.Consumer,
		Streams:  []string{g.Topic, ID},
		Count:    count,
	})
	if err != nil {
		return err
	}
	//处理消息
	for _, msg := range result[0].Messages {
		err = h(&ResponseMsg{
			ID:       msg.ID,
			Topic:    g.Topic,
			Group:    g.GroupName,
			consumer: g.Consumer,
			Body:     msg.Values,
		})
		if err == nil {
			//进行消息确认
			err = cache.XAck(c, g.Topic, g.GroupName, msg.ID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
