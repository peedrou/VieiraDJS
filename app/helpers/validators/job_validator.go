package validators

import (
	"VieiraDJS/app/models"
	"errors"
)

type ValidatedJob struct {
	Job models.Job
}

func (j ValidatedJob) ValidateJob() error {
	if j.Job.IsRecurring && j.Job.Interval == "" {
		return errors.New("interval must be set when is_recurring is true")
	}

	return nil
}
