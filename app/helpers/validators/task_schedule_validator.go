package validators

import (
	"VieiraDJS/app/models"
	"errors"
	"time"
)

type ValidatedTaskSchedule struct {
	TaskSchedule models.TaskSchedule
}

func (ts ValidatedTaskSchedule) ValidateTaskSchedule() error {
	if !isValidUNIXTimestamp(ts.TaskSchedule.Partition) {
		return errors.New("partition is not a valid UNIX timestamp")
	}

	if ts.TaskSchedule.JobId < 0 {
		return errors.New("the Job ID cannot be below 0")
	}

	return nil
}

func isValidUNIXTimestamp(timestamp int64) bool {

	if timestamp <= 0 {
		return false
	}

	now := time.Now().Unix()

	return timestamp <= now
}
