package main

import (
	"context"
	"fmt"
	"pkg_v1/dtviper"
	"pkg_v1/kafka"
)

func main() {
	cfg := dtviper.ConfigInit("ciallo~", "message")

	cGroup := kafka.InitConsumerGroup(cfg.Viper.GetStringSlice("Kafka.Brokers"), cfg.Viper.GetStringSlice("Kafka.GroupIds")[0])
	for {
		err := cGroup.Consume(context.Background(), cfg.Viper.GetStringSlice("ciallo~"), kafka.ConsumerGroup)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
	_ = cGroup.Close()
}
