package builders

import (
	"VieiraDJS/app/helpers/validators"
	"VieiraDJS/app/models"

	"github.com/gocql/gocql"
)

func NewTaskSchedule(nextExecutionTime int64, jobId gocql.UUID) (*validators.ValidatedTaskSchedule, error) {
	taskSchedule := &models.TaskSchedule{
		NextExecutionTime: nextExecutionTime,
		JobId:             jobId,
	}

	validatedTaskSchedule := &validators.ValidatedTaskSchedule{
		TaskSchedule: *taskSchedule,
	}

	if err := validatedTaskSchedule.ValidateTaskSchedule(); err != nil {
		return nil, err
	}

	return validatedTaskSchedule, nil
}
