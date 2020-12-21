package entities

// ProcessInstanceInfo is a type that generally informs about a process instance
type ProcessInstanceInfo struct {
	ID                string `json:"id"`
	TimeStamp         string `json:"timestamp"`
	ProcessName       string `json:"processname"`
	ProcessInstanceID string `json:"processinstanceid"`
	Status            string `json:"status"`
}
