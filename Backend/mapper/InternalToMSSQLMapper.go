package mapper

import (
	"sawu-monitor/entities"
)

func MapInternalToMssql(internalEvent entities.NextStepEvent) entities.MSSQLNextStepEvent {
	mssqlNextStepEvent := new(entities.MSSQLNextStepEvent)
	mssqlNextStepEvent.ComingFromID = internalEvent.ComingFromID
	mssqlNextStepEvent.CorrelationID = internalEvent.CorrelationID
	mssqlNextStepEvent.CorrelationState = internalEvent.CorrelationState
	mssqlNextStepEvent.Data = internalEvent.Data
	mssqlNextStepEvent.ID = internalEvent.ID
	mssqlNextStepEvent.NextRetryAt = internalEvent.NextRetryAt
	mssqlNextStepEvent.RetryCount = internalEvent.RetryCount
	mssqlNextStepEvent.ProcessName = internalEvent.ProcessName
	mssqlNextStepEvent.ProcessStep = internalEvent.ProcessStep
	mssqlNextStepEvent.ProcessInstanceID = internalEvent.ProcessInstanceID
	mssqlNextStepEvent.WaitID = internalEvent.WaitID
	mssqlNextStepEvent.TimeStamp = internalEvent.TimeStamp

	return *mssqlNextStepEvent
}
