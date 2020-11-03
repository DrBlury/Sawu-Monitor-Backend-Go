package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"sawu-monitor/entities"
)

func SendNextStepEvent(topic string, nextStepEvent entities.NextStepEvent) {

	//broker := os.Args[1]
	//topic := os.Args[2]

	broker := "192.168.1.9"

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Producer %v\n", p)

	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	deliveryChan := make(chan kafka.Event)

	value := serialize(nextStepEvent)
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value),
		Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
	}, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)
}

func serialize(nextStepEvent entities.NextStepEvent) string {
	event := fmt.Sprintf("id=%s,timestamp=%s,processName=%s,processInstanceID=%s,processStep=%s,internal=%s,retryCount=%s,$e%%,%s",
		nextStepEvent.ID, nextStepEvent.TimeStamp, nextStepEvent.ProcessName, nextStepEvent.ProcessInstanceID, nextStepEvent.ProcessStep, nextStepEvent.Internal, nextStepEvent.RetryCount, nextStepEvent.Data)

	return event
}
