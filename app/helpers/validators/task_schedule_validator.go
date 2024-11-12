package validators

import (
	"VieiraDJS/app/models"
	"errors"
	"reflect"
	"time"

	"github.com/gocql/gocql"
)

type ValidatedTaskSchedule struct {
	TaskSchedule models.TaskSchedule
}

func (ts ValidatedTaskSchedule) ValidateTaskSchedule() error {
	if !isValidUNIXTimestamp(ts.TaskSchedule.NextExecutionTime) {
		return errors.New("partition is not a valid UNIX timestamp")
	}

	if reflect.TypeOf(ts.TaskSchedule.JobId) != reflect.TypeOf(gocql.UUID{}) {
		return errors.New("JobID is not of type gocql.UUID")
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
