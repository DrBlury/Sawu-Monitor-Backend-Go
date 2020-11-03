package entities

type MSSQLNextStepEvent struct {
	ID                string `json:"id" db:"id"`
	TimeStamp         string `json:"timeStamp" db:"time_stamp"`
	ProcessName       string `json:"processName" db:"process_name"`
	ComingFromID      string `json:"comingFromId" db:"coming_from_id"`
	ProcessStep       string `json:"processStep" db:"process_step"`
	ProcessInstanceID string `json:"processInstanceId" db:"process_instance_id"`
	CorrelationState  string `json:"correlationState" db:"correlation_state"`
	CorrelationID     string `json:"correlationId" db:"correlation_id"`
	RetryCount        string `json:"retryCount" db:"retry_count"`
	NextRetryAt       string `json:"nextRetryAt" db:"next_retry_at"`
	WaitID            string `json:"waitId" db:"wait_id"`
	Data              string `json:"variables" db:"variables"`
}
