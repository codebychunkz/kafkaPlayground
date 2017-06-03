package main

import (
	"github.com/Shopify/sarama"
	"log"
	"time"
)

type Topic string

const (
	TEST_TOPIC Topic = "TestTopic"
)

type ProducerContext struct {
	producer  sarama.SyncProducer
	topic     Topic
	partition int32
}

type ConsumerContext struct {
	consumer  sarama.Consumer
	topic     Topic
	partition int32
	offset    int64
}

func consumerTask(ctx ConsumerContext) {

	consPartition, err := ctx.consumer.ConsumePartition(string(ctx.topic), ctx.partition, ctx.offset)
	if err != nil {
		log.Printf("Unable to consume partition %d", ctx.partition)
		log.Println(err)
		return
	}
	defer consPartition.Close()

	for {
		select {
		case message := <-consPartition.Messages():
			value := string(message.Value)
			log.Printf("Consumed value \"%s\" for partition %d and offset %d", value, message.Partition, message.Offset)
		}
	}

}

func producerTask(ctx ProducerContext) {

	for {
		time.Sleep(time.Millisecond * 250)
		msg := sarama.ProducerMessage{
			Topic:     string(ctx.topic),
			Partition: ctx.partition,
			Value:     sarama.StringEncoder("Hello, World!"),
		}

		part, offset, err := ctx.producer.SendMessage(&msg)
		if err != nil {
			log.Println("Unable to send message")
			log.Println(err)
		} else {
			log.Printf("Sent message with partition %d and offset %d", part, offset)
		}

	}

}
