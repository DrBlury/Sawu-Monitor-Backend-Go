package entities

// KafkaNextStepEvent is the type compatible with the Sawu messages sent in kafka
type KafkaNextStepEvent struct {
	ID                string `json:"id"`
	TimeStamp         string `json:"timestamp"`
	ProcessName       string `json:"processName"`
	ComingFromID      string `json:"comingFromId"`
	ProcessStep       string `json:"processStep"`
	ProcessInstanceID string `json:"processInstanceId"`
	CorrelationState  string `json:"correlationState"`
	CorrelationID     string `json:"correlationId"`
	RetryCount        string `json:"retryCount"`
	NextRetryAt       string `json:"nextRetryAt"`
	Internal          string `json:"internal"`
	WaitID            string `json:"waitId"`
	Error             string `json:"error"`
	Data              string `json:"data"`
	SourceTopic       string `json:"sourceTopic"`
}
