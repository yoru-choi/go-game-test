package core

import (
	"fmt"

	"github.com/IBM/sarama"
)

func getProducer(bootstrapServers []string) sarama.SyncProducer {
	var config *sarama.Config = sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err := sarama.NewSyncProducer(bootstrapServers, config)
	if err != nil {
		fmt.Println("failed to create Producer", err)
		return nil
	}
	return producer
}

func sendMessage(producer sarama.SyncProducer, topic string, value string) {
	kafkaMessage := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(value),
	}

	partition, offset, err := producer.SendMessage(kafkaMessage)
	if err != nil {
		fmt.Printf("%s error occured.", err.Error())
	} else {
		fmt.Printf("Message was saved to partion: %d.\nMessage offset is: %d.\n", partition, offset)
	}
}

func main() {
	var topic string = "test-topic"
	producer := getProducer([]string{"localhost:29091", "localhost:29092", "localhost:29093"})
	for index := 0; index < 1000; index++ {
		var value string = fmt.Sprintf("%s%d", "message", index)
		sendMessage(producer, topic, value)
	}
}
