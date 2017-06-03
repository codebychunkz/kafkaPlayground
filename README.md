# kafkaPlayground
Playing around with kafka, docker and Go

Used https://github.com/wurstmeister/kafka-docker docker image and tutorial for Kafka deployment
Used https://github.com/Shopify/sarama as the Kafka client for Go

## Prerequisites
* Docker
* Docker compose
* Go (tested with 1.8)

## How to
* Navigate to the docker folder (where the docker-compose.yml file is located)
* docker-compose pull   
* run docker-compose up -d
* run docker-compose scale kafka=2
* change the ip and port in the main.go file to the ip and port for kafka (run docker-compose ps to list the ports)
* Build and run
