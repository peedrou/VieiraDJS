package jobs

import (
	crud "VieiraDJS/app/db/CRUD"
	"VieiraDJS/app/helpers/builders"
	"VieiraDJS/app/helpers/converters"
	"VieiraDJS/app/helpers/validators"
	"VieiraDJS/app/models"
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

func CreateJob(session *gocql.Session, userID gocql.UUID, isRecurring bool, maxRetries int, startTime time.Time, interval string) error {
	someUUID := uuid.New()
	gocqlUUID, _ := gocql.ParseUUID(someUUID.String())
	job, err := builders.NewJob(
		userID,
		gocqlUUID,
		isRecurring,
		maxRetries,
		startTime,
		interval,
	)

	if err != nil {
		return fmt.Errorf("failed to create job: %v", err)
	}

	if job == nil {
		return fmt.Errorf("job returned nil")
	}

	err = InsertJobInDB(session, *job)

	if err != nil {
		return fmt.Errorf("there was an error inserting the Job in the database: %v", err)
	}

	return nil

}

func RemoveJobs(session *gocql.Session, IDs []interface{}) error {
	if len(IDs) == 0 {
		return fmt.Errorf("no IDs were provided for job removal")
	}

	err := crud.RemoveModel(session, "jobs", "job_id", IDs)

	if err != nil {
		return fmt.Errorf("there was an error removing the Job(s) from the database: %v", err)
	}

	return nil
}

func UpdateJobs(session *gocql.Session, keyToUpdate string, valueToUpdate interface{}, IDs []interface{}) error {
	if len(IDs) == 0 {
		return fmt.Errorf("no IDs were provided for job update")
	}

	err := crud.UpdateModelBatch(session, "jobs", keyToUpdate, valueToUpdate, "job_id", IDs)

	if err != nil {
		return fmt.Errorf("there was an error updating the Job(s) from the database: %v", err)
	}

	return nil
}

func ReadJobs(session *gocql.Session, keys []string, values ...interface{}) ([]interface{}, error) {
	if len(keys) == 0 {
		return nil, fmt.Errorf("no keys were provided for job reading")
	}

	result, err := crud.ReadModel(session, "jobs", []string{"job_id"}, keys, values...)

	if err != nil {
		return nil, fmt.Errorf("there was an error reading the Job(s) from the database: %v", err)
	}

	return result, nil
}

func InsertJobInDB(session *gocql.Session, job validators.ValidatedJob) error {
	id := job.Job.JobID
	fields := []string{}

	fields = append(fields, "user_id", "job_id", "created_time", "interval", "is_recurring", "max_retries", "start_time")

	err := crud.CreateModel(
		session,
		"jobs",
		fields,
		job.Job.UserID,
		id,
		job.Job.CreatedTime,
		job.Job.Interval,
		job.Job.IsRecurring,
		job.Job.MaxRetries,
		job.Job.StartTime)

	if err != nil {
		return err
	}

	taskSchedule, _ := builders.NewTaskSchedule(
		converters.ConvertExecutionTimeToUNIX(job.Job.StartTime),
		id)

	err = crud.CreateModel(
		session,
		"task_schedule",
		[]string{"next_execution_time", "job_id"},
		taskSchedule.TaskSchedule.NextExecutionTime,
		taskSchedule.TaskSchedule.JobId)

	if err != nil {
		_ = crud.RemoveModel(session, "jobs", "job_id", []interface{}{id})
		return err
	}

	taskHistory, _ := builders.NewTaskHistory(
		id,
		job.Job.CreatedTime,
		models.TaskStatusPending,
		job.Job.MaxRetries,
		job.Job.CreatedTime)

	err = crud.CreateModel(
		session,
		"task_history",
		[]string{"job_id", "execution_time", "status", "retry_count", "last_update_time"},
		id,
		taskHistory.TaskHistory.ExecutionTime,
		taskHistory.TaskHistory.Status,
		taskHistory.TaskHistory.RetryCount,
		taskHistory.TaskHistory.LastUpdateTime)

	if err != nil {
		_ = crud.RemoveModel(session, "jobs", "job_id", []interface{}{id})
		_ = crud.RemoveModel(session, "task_schedule", "job_id", []interface{}{id})
		return err
	}

	return nil
}
