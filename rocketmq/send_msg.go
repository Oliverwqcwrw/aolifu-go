package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main() {
	p, _ := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"localhost:9876"})),
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

	// do something with result

}
