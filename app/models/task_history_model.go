package models

import (
	"time"

	"github.com/gocql/gocql"
)

type TaskStatus string

const (
	TaskStatusUndefined TaskStatus = "UNDEFINED"
	TaskStatusScheduled TaskStatus = "SCHEDULED"
	TaskStatusPending   TaskStatus = "PENDING"
	TaskStatusRunning   TaskStatus = "RUNNING"
	TaskStatusFailed    TaskStatus = "FAILED"
	TaskStatusCompleted TaskStatus = "COMPLETED"
)

type TaskHistory struct {
	JobId          gocql.UUID
	ExecutionTime  time.Time
	Status         TaskStatus
	RetryCount     int
	LastUpdateTime time.Time
}
