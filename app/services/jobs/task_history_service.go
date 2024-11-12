package jobs

import (
	crud "VieiraDJS/app/db/CRUD"
	"VieiraDJS/app/helpers/validators"
	"VieiraDJS/app/models"
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

func UpdateTaskHistory(session *gocql.Session, taskHistory *validators.ValidatedTaskHistory, newStatus models.TaskStatus, new_execution_time time.Time, decreaseCounter bool) error {
	retryCount := taskHistory.TaskHistory.RetryCount

	if decreaseCounter {
		retryCount = retryCount - 1
	}

	err := crud.CreateModel(
		session,
		"task_history",
		[]string{"job_id", "execution_time", "status", "retry_count", "last_update_time"},
		taskHistory.TaskHistory.JobId,
		new_execution_time,
		newStatus,
		retryCount,
		time.Now())

	if err != nil {
		return fmt.Errorf("there was a problem updating the task history with a new entry: %v", err)
	}

	return nil
}
