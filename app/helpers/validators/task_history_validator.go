package validators

import (
	"VieiraDJS/app/models"
	"errors"
)

type ValidatedTaskHistory struct {
	TaskHistory models.TaskHistory
}

func (th ValidatedTaskHistory) ValidateTaskHistory() error {

	if th.TaskHistory.JobId < 0 {
        return errors.New("JobId must be a non-negative integer")
    }

    if th.TaskHistory.JobId < 0 {
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
	case models.TaskStatusCompleted, models.TaskStatusError, models.TaskStatusFailed, models.TaskStatusPending, models.TaskStatusRunning, models.TaskStatusUndefined:
		return true
	}
	return false
}