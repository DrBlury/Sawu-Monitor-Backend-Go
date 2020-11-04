package mapper

import (
	"sawu-monitor/entities"
)

func MapFrontendToInternal(frontendEvent entities.FrontendNextStepEvent) entities.NextStepEvent {
	nextStepEvent := new(entities.NextStepEvent)
	nextStepEvent.ComingFromID = frontendEvent.ComingFromID
	nextStepEvent.CorrelationID = frontendEvent.CorrelationID
	nextStepEvent.CorrelationState = frontendEvent.CorrelationState
	nextStepEvent.Data = frontendEvent.Data
	nextStepEvent.ID = frontendEvent.ID
	nextStepEvent.NextRetryAt = frontendEvent.NextRetryAt
	nextStepEvent.RetryCount = frontendEvent.RetryCount
	nextStepEvent.ProcessName = frontendEvent.ProcessName
	nextStepEvent.ProcessStep = frontendEvent.ProcessStep
	nextStepEvent.ProcessInstanceID = frontendEvent.ProcessInstanceID
	nextStepEvent.WaitID = frontendEvent.WaitID
	nextStepEvent.TimeStamp = frontendEvent.TimeStamp

	return *nextStepEvent
}
