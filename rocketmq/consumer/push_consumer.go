package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"os"
)

func main() {
	sig := make(chan os.Signal)
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName("DEFAULT_CONSUMER_GROUP"),
	)

	err = c.Subscribe("TEST_TOPIC", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		rlog.Info("Subscribe Callback", map[string]interface{}{
			"msgs": msgs,
		})
		return consumer.ConsumeSuccess, nil
	})

	err = c.Start()
	<-sig

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

}
