package integrations

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func StreamLogToKafka(Log string, kafkaConfigs []KafkaLogConfig) error {
	for _, config := range kafkaConfigs {
		log.Printf("[KAFKA] Dialing %v:%v Topic: %s", config.Broker, config.Port, config.Topic)
		producer, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": config.Broker + ":" + config.Port,
		})
		if err != nil {
			log.Printf("err: %s", err)
			return err
		}
		defer producer.Close()
		if producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &config.Topic, Partition: kafka.PartitionAny},
			Value:          []byte(Log),
		}, nil) != nil {
			log.Printf("err: %s", err)
			return err
		}
		producer.Flush(1000 * 2)
		log.Printf("[KAFKA] Log send to %v:%v", config.Broker, config.Port)
	}
	return nil
}
