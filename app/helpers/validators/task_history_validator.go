package validators

import (
	"VieiraDJS/app/models"
	"errors"
	"reflect"

	"github.com/gocql/gocql"
)

type ValidatedTaskHistory struct {
	TaskHistory models.TaskHistory
}

func (th ValidatedTaskHistory) ValidateTaskHistory() error {

	if reflect.TypeOf(th.TaskHistory.JobId) == reflect.TypeOf(gocql.UUID{}) {
		return errors.New("JobID is not of type gocql.UUID")
	}

	if th.TaskHistory.RetryCount < 0 {
		return errors.New("RetryCount must be a non-negative integer")
	}

	if th.TaskHistory.ExecutionTime.IsZero() {
		return errors.New("ExecutionTime must be a valid time")
	}

	if th.TaskHistory.LastUpdateTime.IsZero() {
		return errors.New("LastUpdateTime must be a valid time")
	}

	if !isValidTaskStatus(th.TaskHistory.Status) {
		return errors.New("status must be a valid TaskStatus")
	}

	return nil
}

func isValidTaskStatus(status models.TaskStatus) bool {
	switch status {
	case models.TaskStatusCompleted, models.TaskStatusScheduled, models.TaskStatusFailed, models.TaskStatusPending, models.TaskStatusRunning, models.TaskStatusUndefined:
		return true
	}
	return false
}
