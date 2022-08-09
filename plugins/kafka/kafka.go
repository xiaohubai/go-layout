package kafka

import (
	"time"

	"github.com/Shopify/sarama"
	"github.com/xiaohubai/go-layout/configs/global"
)

func Producer() sarama.SyncProducer {
	cfg := sarama.NewConfig()
	cfg.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	cfg.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	cfg.Producer.Retry.Max = 3                             // 重试三次
	cfg.Producer.Retry.Backoff = 100 * time.Millisecond
	p, err := sarama.NewSyncProducer(global.Cfg.Kafka.Address, cfg)
	if err != nil {
		panic(err)
	}
	return p
}

func Consumer() sarama.Consumer {
	cfg := sarama.NewConfig()
	c, err := sarama.NewConsumer(global.Cfg.Kafka.Address, cfg)
	if err != nil {
		panic(err)
	}
	return c
}

func WriteToKafka(topic, data string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	_, _, _ = global.KafkaProducer.SendMessage(msg)
}

/*
func ReadFromKafka(topics string) {
	topic := strings.Split(topics, ",")
	for _, t := range topic {
		switch t {
		case consts.TopicOfWarn:
			go v1.Warn(t)
		default:
			go v1.Warn(t)
		}
	}
}
*/
