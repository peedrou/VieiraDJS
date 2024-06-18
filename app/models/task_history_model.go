package models

import (
	"time"
)

type TaskStatus string

const (
	TaskStatusUndefined TaskStatus = "UNDEFINED"
	TaskStatusError     TaskStatus = "ERROR"
	TaskStatusPending   TaskStatus = "PENDING"
	TaskStatusRunning   TaskStatus = "RUNNING"
	TaskStatusFailed    TaskStatus = "FAILED"
	TaskStatusCompleted TaskStatus = "COMPLETED"
)

type TaskHistory struct {
	JobId          int
	ExecutionTime  time.Time
	Status         TaskStatus
	RetryCount     int
	LastUpdateTime time.Time
}
