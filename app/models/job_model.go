package models

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	JobID       uuid.UUID
	IsRecurring bool
	StartTime   time.Time
	Interval    string
	MaxRetries  int
	CreatedTime time.Time
}
