package entities

// ProcessInstanceInfo is a type that generally informs about a process instance
type ProcessInstanceInfo struct {
	ID                string `json:"id" db:"id"`
	TimeStamp         string `json:"timestamp" db:"time_stamp"`
	ProcessName       string `json:"processname"  db:"process_name"`
	ProcessInstanceID string `json:"processinstanceid" db:"process_instance_id"`
	Status            string `json:"status" db:"status"`
}
