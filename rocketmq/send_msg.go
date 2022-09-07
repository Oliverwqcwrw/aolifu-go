package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
)

func main() {
	p, _ := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		producer.WithRetry(2),
	)
	err := p.Start()
	if err != nil {
		fmt.Printf("send message error: %s\n", err)
	}

	result, err := p.SendSync(context.Background(), &primitive.Message{
		Topic: "TEST_TOPIC",
		Body:  []byte("Hello RocketMQ Go Client!"),
	})
	fmt.Printf("send message success: result=%s\n", result.String())
	err = p.SendOneWay(context.Background(), &primitive.Message{
		Topic: "TEST_TOPIC",
		Body:  []byte("Hello RocketMQ Go Client!"),
	})
	if err != nil {
		rlog.Info("123", nil)
	}

}
