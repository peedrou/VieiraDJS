package models

import (
	"time"
)

type Job struct {
	IsRecurring bool
	StartTime time.Time
	Interval string
	MaxRetries int
	CreatedTime time.Time
}
