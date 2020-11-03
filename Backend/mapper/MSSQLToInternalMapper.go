package mapper

import (
	"sawu-monitor/entities"
)

func MapMssqlToInternal(mssqlEvent entities.MSSQLNextStepEvent) entities.NextStepEvent {
	nextStepEvent := new(entities.NextStepEvent)
	nextStepEvent.ComingFromID = mssqlEvent.ComingFromID
	nextStepEvent.CorrelationID = mssqlEvent.CorrelationID
	nextStepEvent.CorrelationState = mssqlEvent.CorrelationState
	nextStepEvent.Data = mssqlEvent.Data
	nextStepEvent.ID = mssqlEvent.ID
	nextStepEvent.NextRetryAt = mssqlEvent.NextRetryAt
	nextStepEvent.RetryCount = mssqlEvent.RetryCount
	nextStepEvent.ProcessName = mssqlEvent.ProcessName
	nextStepEvent.ProcessStep = mssqlEvent.ProcessStep
	nextStepEvent.ProcessInstanceID = mssqlEvent.ProcessInstanceID
	nextStepEvent.WaitID = mssqlEvent.WaitID
	nextStepEvent.TimeStamp = mssqlEvent.TimeStamp

	return *nextStepEvent
}
