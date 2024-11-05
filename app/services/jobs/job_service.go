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
