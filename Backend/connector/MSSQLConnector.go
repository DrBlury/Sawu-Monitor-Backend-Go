package connector

import (
	"fmt"
	"log"
	"os"
	"sawu-monitor/config"
	"sawu-monitor/entities"
	"sawu-monitor/mapper"
	"strings"

	// go-mssqldb is the database driver
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

var nextstepeventSchema = `
CREATE TABLE nextstepevent (
	id varchar(255) NOT NULL,
	coming_from_id varchar(255) NULL,
	correlation_id varchar(255) NULL,
	correlation_state varchar(255) NULL,
	next_retry_at varchar(255) NULL,
	process_instance_id varchar(255) NULL,
	process_name varchar(255) NULL,
	process_step varchar(255) NULL,
	retry_count varchar(255) NULL,
	time_stamp varchar(255) NULL,
	variables varchar(MAX) NULL,
	wait_id varchar(255) NULL,
	CONSTRAINT PK_nextstep PRIMARY KEY (id)
)`

var db *sqlx.DB

// ConnectDB connects the application to the database
func ConnectDB() {

	var defaults config.Conf
	defaults.GetDefaults()

	//Set default consumergroup if not set
	server, isPresent := os.LookupEnv("database_server")
	if isPresent == false {
		server = defaults.Database.Server
	}

	//Set default consumergroup if not set
	user, isPresent := os.LookupEnv("database_user")
	if isPresent == false {
		user = defaults.Database.User
	}

	//Set default consumergroup if not set
	password, isPresent := os.LookupEnv("database_password")
	if isPresent == false {
		password = defaults.Database.Password
	}

	//Set default port if not set
	port, isPresent := os.LookupEnv("database_port")
	if isPresent == false {
		port = defaults.Database.Port
	}

	connString := fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s", server, port, user, password)
	log.Println(connString)
	var err error
	db, err = sqlx.Connect("sqlserver", connString)

	if err != nil {
		log.Fatalln(err)
	}

	db.Exec(nextstepeventSchema)

	tx := db.MustBegin()
	tx.Commit()
}

func escapeDBletters(data string) string {
	return strings.Replace(data, "'", "''", -1)
}

// FindAllProcessesInstanceInfo returns a list of all process Instances
func FindAllProcessesInstanceInfo() []entities.ProcessInstanceInfo {
	// TODO: Implement
	return nil
}

// FindProcessInstanceInfoByDataValue returns a ProcessInstanceInfo that has events containing the value
func FindProcessInstanceInfoByDataValue(value string) entities.ProcessInstanceInfo {
	processInstanceInfo := new(entities.ProcessInstanceInfo)

	// TODO: Implement

	return *processInstanceInfo
}

// FindProcessEventsByProcessInstanceID returns a list of Process Events in kafka format
func FindProcessEventsByProcessInstanceID(processInstanceID string) []entities.KafkaNextStepEvent {
	selectString := fmt.Sprintf("SELECT * FROM nextstepevent WHERE process_instance_id LIKE '%%%s%%'", processInstanceID)
	fmt.Println(selectString)
	mssqlNextStepEvents := []entities.MSSQLNextStepEvent{}
	db.Select(&mssqlNextStepEvents, selectString)

	var nextStepEvents []entities.KafkaNextStepEvent
	for i := 0; i <= len(mssqlNextStepEvents)-1; i++ {
		event := mapper.MapMssqlToKafka(mssqlNextStepEvents[i])
		nextStepEvents = append(nextStepEvents, event)
	}

	return nextStepEvents
}

// CreateNewEvent triggers a new process event with the defined data (can be used to fix a process instance)
func CreateNewEvent(internalNextStepEvent entities.KafkaNextStepEvent) {
	event := mapper.MapKafkaToMssql(internalNextStepEvent)

	insertString := fmt.Sprintf(
		"INSERT INTO nextstepevent"+
			"(id, coming_from_id, correlation_id, correlation_state, next_retry_at, process_instance_id, process_name, process_step, retry_count, time_stamp, variables, wait_id)"+
			"VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');",
		event.ID, event.ComingFromID, event.CorrelationID,
		event.CorrelationState, event.NextRetryAt,
		event.ProcessInstanceID, event.ProcessName,
		event.ProcessStep, event.RetryCount,
		event.TimeStamp, escapeDBletters(event.Data),
		event.WaitID)

	tx := db.MustBegin()
	result, err := tx.Exec(insertString)

	if err != nil {
		log.Printf("Error: %s \nResult: %s", err, result)
	}
	tx.Commit()
}
