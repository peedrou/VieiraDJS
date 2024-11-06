package jobs

import (
	crud "VieiraDJS/app/db/CRUD"
	"VieiraDJS/app/helpers/builders"
	"VieiraDJS/app/helpers/validators"
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

func CreateJob(session *gocql.Session, isRecurring bool, maxRetries int, startTime time.Time, interval string) error {
	someUUID := uuid.New()
	gocqlUUID, _ := gocql.ParseUUID(someUUID.String())
	job, err := builders.NewJob(
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

func UpdateJobs(session *gocql.Session, keyToUpdate string, valueToUpdate string, IDs []interface{}) error {
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
	fields := []string{}

	fields = append(fields, "job_id", "created_time", "interval", "is_recurring", "max_retries", "start_time")

	err := crud.CreateModel(
		session,
		"jobs",
		fields,
		job.Job.JobID,
		job.Job.CreatedTime,
		job.Job.Interval,
		job.Job.IsRecurring,
		job.Job.MaxRetries,
		job.Job.StartTime)

	if err != nil {
		return err
	}

	return nil
}
