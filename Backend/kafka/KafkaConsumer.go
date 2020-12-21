package kafka

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sawu-monitor/config"
	"sawu-monitor/connector"
	"sawu-monitor/entities"
	"strings"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var separator string

// DoKafkaConsumerStuff is the method to start the kafka consumer
func DoKafkaConsumerStuff() {
	// Load config.yml
	var defaults config.Conf
	defaults.GetDefaults()

	//Set default broker ip if not set
	broker, isPresent := os.LookupEnv("kafka_broker_ip")
	if isPresent == false {
		broker = defaults.Kafka.Broker.IPAddress
	}

	//Set default separator ip if not set
	separator, isPresent = os.LookupEnv("sawu_separator_string")
	if isPresent == false {
		separator = defaults.Sawu.Separator
	}

	//Set default consumergroup if not set
	group, isPresent := os.LookupEnv("kafka_consumer_group")
	if isPresent == false {
		group = defaults.Kafka.Consumer.ConsumerGroup
	}

	var consumerTopics []string

	//Set default broker ip if not set
	topics, isPresent := os.LookupEnv("kafka_consumer_topics")
	if isPresent == false {
		consumerTopics = defaults.Kafka.Consumer.Topics
	} else {
		consumerTopics = strings.Split(topics, ", ")
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  broker,
		"group.id":           group,
		"session.timeout.ms": 6000,
		"auto.offset.reset":  "earliest"})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %v\n", consumer)

	err = consumer.SubscribeTopics(consumerTopics, nil)

	run := true

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := consumer.Poll(100)
			if ev == nil {
				continue
			}
			switch event := ev.(type) {
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n",
					event.TopicPartition, string(event.Value))
				fmt.Println("==============================")

				processEvent, err := deSerialize(string(event.Value))
				if true == bool(err) {
					recover()
				}

				jsonString, _ := json.Marshal(processEvent)
				fmt.Println(string(jsonString))

				connector.CreateNewEvent(processEvent)
				if err {
					fmt.Println("I deserialized and retrieved an error. Aborting.")
					continue
				}

				// if event.Headers != nil {
				// 	fmt.Printf("%% Headers: %v\n", event.Headers)
				// }
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", event.Code(), event)
				if event.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Printf("Ignored %v\n", event)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	consumer.Close()
	os.Exit(0)
}

func deSerialize(kafkaRecord string) (entities.KafkaNextStepEvent, bool) {

	seperatorIndex := strings.Index(kafkaRecord, separator)
	runes := []rune(kafkaRecord)
	processEvent := new(entities.KafkaNextStepEvent)
	processEvent.Data = string(runes[seperatorIndex+len(separator) : len(runes)])

	caughtSeparator := false
	eventValue := string(kafkaRecord)
	eventData := strings.Split(eventValue, ",")
	for dataIterator := 0; dataIterator < len(eventData); dataIterator++ {
		VariableAndValue := strings.Split(eventData[dataIterator], "=")
		// Switch case to add the values to the event
		switch VariableAndValue[0] {
		case "id":
			processEvent.ID = VariableAndValue[1]
		case "timestamp":
			processEvent.TimeStamp = VariableAndValue[1]
		case "processName":
			processEvent.ProcessName = VariableAndValue[1]
		case "processInstanceID":
			processEvent.ProcessInstanceID = VariableAndValue[1]
		case "processStep":
			processEvent.ProcessStep = VariableAndValue[1]
		case "internal":
			processEvent.Internal = VariableAndValue[1]
		case "retryCount":
			processEvent.RetryCount = VariableAndValue[1]
		case "nextRetryAt":
			processEvent.NextRetryAt = VariableAndValue[1]
		case "waitID":
			processEvent.WaitID = VariableAndValue[1]
		case "error":
			processEvent.Error = VariableAndValue[1]
		case "correlationState":
			processEvent.CorrelationState = VariableAndValue[1]
		case "correlationID":
			processEvent.CorrelationID = VariableAndValue[1]
		case "comingFromID":
			processEvent.ComingFromID = VariableAndValue[1]
		case "$e%":
			caughtSeparator = true
			break
		default:
			if !caughtSeparator {
				fmt.Println(VariableAndValue)
				fmt.Println("Whatever that is... It's not a process Event.")
				return *processEvent, true
			}
		}
	}
	return *processEvent, false
}
