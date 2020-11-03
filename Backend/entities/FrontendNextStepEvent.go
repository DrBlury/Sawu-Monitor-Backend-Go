package entities

type FrontendNextStepEvent struct {
	ID                string `json:"id"`
	TimeStamp         string `json:"timeStamp"`
	ProcessName       string `json:"processName"`
	ComingFromID      string `json:"comingFromId"`
	ProcessStep       string `json:"processStep"`
	ProcessInstanceID string `json:"processInstanceId"`
	CorrelationState  string `json:"correlationState"`
	CorrelationID     string `json:"correlationId"`
	RetryCount        string `json:"retryCount"`
	NextRetryAt       string `json:"nextRetryAt"`
	WaitID            string `json:"waitId"`
	Data              string `json:"variables"`
}
