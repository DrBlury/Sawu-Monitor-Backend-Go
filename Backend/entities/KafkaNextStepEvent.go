package entities

// KafkaNextStepEvent is the type compatible with the Sawu messages sent in kafka
type KafkaNextStepEvent struct {
	ID                string `json:"id"`
	TimeStamp         string `json:"timestamp"`
	ProcessName       string `json:"processname"`
	ComingFromID      string `json:"comingfromid"`
	ProcessStep       string `json:"processstep"`
	ProcessStepClass  string `json:"processstepclass"`
	ProcessInstanceID string `json:"processinstanceid"`
	CorrelationState  string `json:"correlationstate"`
	CorrelationID     string `json:"correlationid"`
	RetryCount        string `json:"retrycount"`
	NextRetryAt       string `json:"nextretryat"`
	Internal          string `json:"internal"`
	WaitID            string `json:"waitid"`
	Error             string `json:"error"`
	Data              string `json:"data"`
}
