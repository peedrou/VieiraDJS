package jobs

import (
	crud "VieiraDJS/app/db/CRUD"
	"VieiraDJS/app/helpers/validators"
	"fmt"

	"github.com/gocql/gocql"
)

func UpdateTaskSchedule(session *gocql.Session, taskSchedule *validators.ValidatedTaskSchedule, new_execution_time int64) error {
	err := crud.UpdateModelBatch(
		session,
		"task_schedule",
		"next_execution_time",
		new_execution_time,
		"job_id",
		[]interface{}{taskSchedule.TaskSchedule.JobId})

	if err != nil {
		return fmt.Errorf("there was a problem updating the task schedule: %v", err)
	}

	return nil
}

func RemoveTaskSchedule(session *gocql.Session, taskSchedule *validators.ValidatedTaskSchedule) error {
	err := crud.RemoveModel(
		session,
		"task_schedule",
		"job_id",
		[]interface{}{taskSchedule.TaskSchedule.JobId})

	if err != nil {
		return fmt.Errorf("there was a problem removing the task schedule: %v", err)
	}

	return nil
}
