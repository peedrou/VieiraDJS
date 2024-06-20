package builders

import (
	"VieiraDJS/app/helpers/validators"
	"VieiraDJS/app/models"
)

func NewTaskSchedule(partition int64, jobId int) (*validators.ValidatedTaskSchedule, error) {
	taskSchedule := &models.TaskSchedule{
		Partition: partition,
		JobId: jobId,
	}

	validatedTaskSchedule := &validators.ValidatedTaskSchedule{
		TaskSchedule: *taskSchedule,
	}

	if err := validatedTaskSchedule.ValidateTaskSchedule(); err != nil {
		return nil, err
	}

	return validatedTaskSchedule, nil
}