package scheduler

import (
	crud "VieiraDJS/app/db/CRUD"
	"VieiraDJS/app/helpers/builders"
	"VieiraDJS/app/helpers/converters"
	"VieiraDJS/app/kafka"
	"VieiraDJS/app/models"
	"VieiraDJS/app/services/jobs"
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

func CheckPendingTasks(session *gocql.Session) ([]interface{}, error) {
	currentTime := time.Now()
	nextMinuteExecutionTime := converters.ConvertExecutionTimeToUNIX(currentTime)

	result, err := crud.ReadModel(
		session,
		"task_schedule",
		[]string{"job_id"},
		[]string{"next_execution_time"},
		nextMinuteExecutionTime)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch this minute's tasks: %v", err)
	}

	return result, nil
}

func SchedulePendingTasks(kp *kafka.KafkaProducer, tasks []interface{}, session *gocql.Session) ([]interface{}, []interface{}, error) {
	var tasksSucceeded []interface{}
	var tasksFailed []interface{}
	var retryCount int
	var isRecurring bool
	var interval string

	for _, task := range tasks {
		taskUUID, ok := task.(gocql.UUID)
		if !ok {
			return tasksSucceeded, tasksFailed, fmt.Errorf("invalid task format: %v", task)
		}

		result, err := jobs.RetrieveRetryCount(session, taskUUID)

		if err != nil {
			return tasksSucceeded, tasksFailed, fmt.Errorf("failed to retrieve retry count: %v", task)
		}

		for _, r := range result {
			retryCount = r.(int)
		}

		validatedTaskHistory, err := builders.NewTaskHistory(
			taskUUID,
			time.Now(),
			models.TaskStatusPending,
			retryCount,
			time.Now())

		if err != nil {
			return tasksSucceeded, tasksFailed, fmt.Errorf("failed to build task history: %v", err)
		}

		taskMessage, err := SendMessageToKafkaTopic("task-schedule", taskUUID, kp)

		if err != nil {
			log.Printf("failed to send task message: %v", err)
			tasksFailed = append(tasksFailed, task)
			err = jobs.UpdateTaskHistory(session, validatedTaskHistory, models.TaskStatusFailed, time.Now(), true)
			if err != nil {
				return tasksSucceeded, tasksFailed, fmt.Errorf("failed to update task history: %v", err)
			}
		} else {
			log.Printf("Task sent: %s", taskMessage)
			tasksSucceeded = append(tasksSucceeded, task)
			err = jobs.UpdateTaskHistory(session, validatedTaskHistory, models.TaskStatusScheduled, time.Now(), false)
			if err != nil {
				return tasksSucceeded, tasksFailed, fmt.Errorf("failed to update task history: %v", err)
			}

			result, err = jobs.RetrieveJobIsRecurring(session, taskUUID)
			if err != nil {
				return tasksSucceeded, tasksFailed, fmt.Errorf("failed to retrieve task is_recurring: %v", err)
			}

			for _, r := range result {
				isRecurring = r.(bool)
			}

			if isRecurring {
				result, err = jobs.RetrieveJobInterval(session, taskUUID)
				if err != nil {
					return tasksSucceeded, tasksFailed, fmt.Errorf("failed to retrieve task interval: %v", err)
				}

				for _, r := range result {
					interval = r.(string)
				}

				nextExecutionTime := converters.CalculateNextExecutionTimeInUnix(interval)
				newTaskSchedule, err := builders.NewTaskSchedule(nextExecutionTime, taskUUID)
				if err != nil {
					return tasksSucceeded, tasksFailed, fmt.Errorf("failed to create new task schedule: %v", err)
				}

				err = jobs.CreateTaskSchedule(session, newTaskSchedule)
				if err != nil {
					return tasksSucceeded, tasksFailed, fmt.Errorf("failed to create new task schedule: %v", err)
				}
			}

		}
	}

	return tasksSucceeded, tasksFailed, nil
}

func SendMessageToKafkaTopic(topic string, taskUUID gocql.UUID, kp *kafka.KafkaProducer) (string, error) {
	taskMessage := taskUUID.String()
	err := kp.SendMessage(topic, taskMessage)
	return taskMessage, err
}
