package scheduler

import (
	crud "VieiraDJS/app/db/CRUD"
	"VieiraDJS/app/helpers/converters"
	"VieiraDJS/app/kafka"
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

func SchedulePendingTasks(kp *kafka.KafkaProducer, tasks []interface{}) ([]interface{}, []interface{}, error) {
	var tasksSucceeded []interface{}
	var tasksFailed []interface{}

	for _, task := range tasks {
		taskMessage, ok := task.(string)
		if !ok {
			return tasksSucceeded, tasksFailed, fmt.Errorf("invalid task format: %v", task)
		}

		topic := "task-schedule"

		err := kp.SendMessage(topic, taskMessage)
		if err != nil {
			log.Printf("failed to send task message: %v", err)
			tasksFailed = append(tasksFailed, task)
		} else {
			log.Printf("Task sent: %s", taskMessage)
			tasksSucceeded = append(tasksSucceeded, task)
		}
	}

	return tasksSucceeded, tasksFailed, nil
}
