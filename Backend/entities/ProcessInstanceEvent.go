package entities

// ProcessInstanceEvent is a type containing details about a process event
type ProcessInstanceEvent struct {
	ID               string `json:"id"`
	TimeStamp        string `json:"timeStamp"`
	ComingFromID     string `json:"comingFromId"`
	ProcessStep      string `json:"processStep"`
	CorrelationState string `json:"correlationState"`
	CorrelationID    string `json:"correlationId"`
	RetryCount       string `json:"retryCount"`
	NextRetryAt      string `json:"nextRetryAt"`
	WaitID           string `json:"waitId"`
	Data             string `json:"variables"`
}
