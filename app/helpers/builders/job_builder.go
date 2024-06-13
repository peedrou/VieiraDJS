package builders

import (
	"VieiraDJS/app/helpers/validators"
	"VieiraDJS/app/models"
	"time"
)

func NewJob(isRecurring bool, maxRetries int, startTime time.Time, interval string) (*validators.ValidatedJob, error) {
	job := &models.Job{
		IsRecurring: isRecurring,
		MaxRetries: maxRetries,
		StartTime:   startTime,
		Interval:    interval,
		CreatedTime: time.Now(),
	}

	validated_job := &validators.ValidatedJob{
		Job: *job,
	}

	if err := validated_job.ValidateJob(); err != nil {
		return nil, err
	}

	return validated_job, nil
}
