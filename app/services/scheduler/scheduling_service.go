package scheduler

import (
	crud "VieiraDJS/app/db/CRUD"
	"VieiraDJS/app/helpers/converters"
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

func CheckPendingTasks(session *gocql.Session) ([]interface{}, error) {
	currentTime := time.Now()
	nextMinuteExecutionTime := converters.ConvertExecutionTimeToUNIX(currentTime)

	result, err := crud.ReadModel(
		session,
		"task_schedule",
		[]string{"job_id"},
		[]string{"next_execution_time"},
		nextMinuteExecutionTime)
	
	if err != nil {
		return nil, fmt.Errorf("failed to fetch this minute's tasks: %v", err)
	}

	return result, nil
}