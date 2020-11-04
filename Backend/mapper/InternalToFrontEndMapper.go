package mapper

import (
	"sawu-monitor/entities"
)

func MapInternalToFrontend(internalEvent entities.NextStepEvent) entities.FrontendNextStepEvent {
	frontendNextStepEvent := new(entities.FrontendNextStepEvent)
	frontendNextStepEvent.ComingFromID = internalEvent.ComingFromID
	frontendNextStepEvent.CorrelationID = internalEvent.CorrelationID
	frontendNextStepEvent.CorrelationState = internalEvent.CorrelationState
	frontendNextStepEvent.Data = internalEvent.Data
	frontendNextStepEvent.ID = internalEvent.ID
	frontendNextStepEvent.NextRetryAt = internalEvent.NextRetryAt
	frontendNextStepEvent.RetryCount = internalEvent.RetryCount
	frontendNextStepEvent.ProcessName = internalEvent.ProcessName
	frontendNextStepEvent.ProcessStep = internalEvent.ProcessStep
	frontendNextStepEvent.ProcessInstanceID = internalEvent.ProcessInstanceID
	frontendNextStepEvent.WaitID = internalEvent.WaitID
	frontendNextStepEvent.TimeStamp = internalEvent.TimeStamp

	return *frontendNextStepEvent
}
