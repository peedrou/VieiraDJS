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

	for _, task := range tasks {
		taskUUID, ok := task.(gocql.UUID)
		if !ok {
			return tasksSucceeded, tasksFailed, fmt.Errorf("invalid task format: %v", task)
		}

		result, err := crud.ReadModel(
			session,
			"task_history",
			[]string{"retry_count"},
			[]string{"job_id"},
			taskUUID)

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

		topic := "task-schedule"
		taskMessage := taskUUID.String()

		err = kp.SendMessage(topic, taskMessage)
		if err != nil {
			log.Printf("failed to send task message: %v", err)
			tasksFailed = append(tasksFailed, task)
			err = jobs.UpdateTaskHistory(session, validatedTaskHistory, models.TaskStatusFailed, time.Now(), true)
		} else {
			log.Printf("Task sent: %s", taskMessage)
			tasksSucceeded = append(tasksSucceeded, task)
		}
	}

	return tasksSucceeded, tasksFailed, nil
}
