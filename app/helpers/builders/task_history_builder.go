package builders

import (
	"VieiraDJS/app/helpers/validators"
	"VieiraDJS/app/models"
	"time"
)

func NewTaskHistory(jobId int, executionTime time.Time, status models.TaskStatus, retryCount int, lastUpdateTime time.Time) (*validators.ValidatedTaskHistory, error) {
	taskHistory := &models.TaskHistory{
		JobId:          jobId,
		ExecutionTime:  executionTime,
		Status:         status,
		RetryCount:     retryCount,
		LastUpdateTime: lastUpdateTime,
	}

	validatedTaskHistory := &validators.ValidatedTaskHistory{
		TaskHistory: *taskHistory,
	}

	if err := validatedTaskHistory.ValidateTaskHistory(); err != nil {
		return nil, err
	}

	return validatedTaskHistory, nil
}
