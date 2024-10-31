package builders

import (
	"VieiraDJS/app/helpers/validators"
	"VieiraDJS/app/models"
	"time"

	"github.com/google/uuid"
)

func NewJob(jobID uuid.UUID, isRecurring bool, maxRetries int, startTime time.Time, interval string) (*validators.ValidatedJob, error) {
	job := &models.Job{
		JobID:       jobID,
		IsRecurring: isRecurring,
		MaxRetries:  maxRetries,
		StartTime:   startTime,
		Interval:    interval,
		CreatedTime: time.Now(),
	}

	validatedJob := &validators.ValidatedJob{
		Job: *job,
	}

	if err := validatedJob.ValidateJob(); err != nil {
		return nil, err
	}

	return validatedJob, nil
}
