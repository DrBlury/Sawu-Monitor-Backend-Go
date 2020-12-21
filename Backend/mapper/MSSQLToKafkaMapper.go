package mapper

import (
	"sawu-monitor/entities"
)

// MapMssqlToKafka converts the MSSQL model to the Kafka/Internal model
func MapMssqlToKafka(mssqlEvent entities.MSSQLNextStepEvent) entities.KafkaNextStepEvent {
	kafkaEvent := new(entities.KafkaNextStepEvent)
	kafkaEvent.ComingFromID = mssqlEvent.ComingFromID
	kafkaEvent.CorrelationID = mssqlEvent.CorrelationID
	kafkaEvent.CorrelationState = mssqlEvent.CorrelationState
	kafkaEvent.Data = mssqlEvent.Data
	kafkaEvent.ID = mssqlEvent.ID
	kafkaEvent.NextRetryAt = mssqlEvent.NextRetryAt
	kafkaEvent.RetryCount = mssqlEvent.RetryCount
	kafkaEvent.ProcessName = mssqlEvent.ProcessName
	kafkaEvent.ProcessStep = mssqlEvent.ProcessStep
	kafkaEvent.ProcessInstanceID = mssqlEvent.ProcessInstanceID
	kafkaEvent.WaitID = mssqlEvent.WaitID
	kafkaEvent.TimeStamp = mssqlEvent.TimeStamp

	return *kafkaEvent
}
