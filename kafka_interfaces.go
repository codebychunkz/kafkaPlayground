package main

import "github.com/Shopify/sarama"

func newClient(urls []string) (sarama.Client, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	return sarama.NewClient(urls, config)
}

func newProducer(client sarama.Client) (sarama.SyncProducer, error) {
	return sarama.NewSyncProducerFromClient(client)
}

func newConsumer(client sarama.Client) (sarama.Consumer, error) {
	return sarama.NewConsumerFromClient(client)
}
