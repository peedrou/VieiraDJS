package jobs

import (
	"VieiraDJS/app/helpers/builders"
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

func CreateJob(session *gocql.Session, isRecurring bool, maxRetries int, startTime time.Time, interval string) (string, error) {
	job, err := builders.NewJob(
		isRecurring,
		maxRetries,
		startTime,
		interval,
	)

	if err != nil {
		return "", fmt.Errorf("failed to create job: %v", err)
	}

	if job == nil {
		return "", fmt.Errorf("job returned nil")
	}

	// response, err := InsertJobInDB(*job)

	if err != nil {
		return "", fmt.Errorf("there was an error inserting the Job in the database: %v", err)
	}

	return "", nil

}
