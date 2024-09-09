package jobs

import (
	crud "VieiraDJS/app/db/CRUD"
	"VieiraDJS/app/helpers/builders"
	"VieiraDJS/app/helpers/validators"
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

func CreateJob(session *gocql.Session, isRecurring bool, maxRetries int, startTime time.Time, interval string) error {
	job, err := builders.NewJob(
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
	err := crud.CreateModel(session, "jobs", job)

	if err != nil {
		return err
	}

	return nil
}