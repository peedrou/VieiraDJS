package validators

import (
	"VieiraDJS/app/models"
	"errors"
	"strconv"
)

type ValidatedJob struct {
	Job models.Job
}

func (j ValidatedJob) ValidateJob() error {
	if j.Job.IsRecurring && j.Job.Interval == "" {
		return errors.New("interval must be set when is_recurring is true")
	}

	lastChar := j.Job.Interval[len(j.Job.Interval)-1:]
	restOfInterval := j.Job.Interval[:len(j.Job.Interval)-1]

	if lastChar != "M" && lastChar != "H" && lastChar != "D" && lastChar != "W" {
		return errors.New("interval is not in a valid format")
	}

	_, err := strconv.Atoi(restOfInterval)

	if err != nil {
		return errors.New("interval does not contain a number")
	}

	return nil
}
