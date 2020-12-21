package mapper

import (
	"sawu-monitor/entities"
)

// MapKafkaToMssql converts the Kafka/Internal model to the MSSQL model
func MapKafkaToMssql(kafkaEvent entities.KafkaNextStepEvent) entities.MSSQLNextStepEvent {
	mssqlNextStepEvent := new(entities.MSSQLNextStepEvent)
	mssqlNextStepEvent.ComingFromID = kafkaEvent.ComingFromID
	mssqlNextStepEvent.CorrelationID = kafkaEvent.CorrelationID
	mssqlNextStepEvent.CorrelationState = kafkaEvent.CorrelationState
	mssqlNextStepEvent.Data = kafkaEvent.Data
	mssqlNextStepEvent.ID = kafkaEvent.ID
	mssqlNextStepEvent.NextRetryAt = kafkaEvent.NextRetryAt
	mssqlNextStepEvent.RetryCount = kafkaEvent.RetryCount
	mssqlNextStepEvent.ProcessName = kafkaEvent.ProcessName
	mssqlNextStepEvent.ProcessStep = kafkaEvent.ProcessStep
	mssqlNextStepEvent.ProcessInstanceID = kafkaEvent.ProcessInstanceID
	mssqlNextStepEvent.WaitID = kafkaEvent.WaitID
	mssqlNextStepEvent.TimeStamp = kafkaEvent.TimeStamp

	return *mssqlNextStepEvent
}
