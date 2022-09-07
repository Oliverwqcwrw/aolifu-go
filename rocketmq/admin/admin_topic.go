package main

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
)

func main() {
	testAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})))
	err = testAdmin.CreateTopic(
		context.Background(),
		admin.WithTopicCreate("TEST_NEW_TOPIC2"),
		admin.WithBrokerAddrCreate("127.0.0.1:10911"),
	)

	err = testAdmin.DeleteTopic(
		context.Background(),
		admin.WithTopicDelete("TEST_NEW_TOPIC2"),
		//admin.WithBrokerAddrDelete("127.0.0.1:10911"),	//optional
		//admin.WithNameSrvAddr(nameSrvAddr),				//optional
	)

	if err != nil {
		rlog.Info(err.Error(), nil)
	}

}
