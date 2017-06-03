package main

import (
	"github.com/Shopify/sarama"
	"log"
	"time"
)

//https://github.com/wurstmeister/kafka-docker
//https://medium.com/@Oskarr3/implementing-cqrs-using-kafka-and-sarama-library-in-golang-da7efa3b77fe
//https://gist.github.com/JnBrymn/6fc38872b4d312886908#file-sarama_example-go-L22
//https://github.com/Shopify/sarama/wiki
//https://github.com/Shopify/sarama/tree/master/examples

//docker-compose up -d
//docker-compose scale kafka=2

func main() {
	log.Println("START")

	client, err := newClient([]string{"192.168.99.100:32776", "192.168.99.100:32777"})

	if err != nil {
		log.Println("Unable to create Kafka client")
		log.Fatal(err)
	}
	defer client.Close()

	producer, err := newProducer(client)
	if err != nil {
		log.Println("Unable to create Kafka producer")
		log.Fatal(err)
	}
	defer producer.Close()

	consumer, err := newConsumer(client)
	if err != nil {
		log.Println("Unable to create Kafka consumer")
		log.Fatal(err)
	}
	defer consumer.Close()

	go producerTask(ProducerContext{topic: TEST_TOPIC,
		partition: 0,
		producer:  producer})

	go consumerTask(ConsumerContext{topic: TEST_TOPIC,
		partition: 0,
		offset:    sarama.OffsetOldest,
		consumer:  consumer})

	time.Sleep(time.Second * 1)

	log.Println("DONE")
}
