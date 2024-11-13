package models

import (
	"time"

	"github.com/gocql/gocql"
)

type Job struct {
	UserID      gocql.UUID
	JobID       gocql.UUID
	IsRecurring bool
	StartTime   time.Time
	Interval    string
	MaxRetries  int
	CreatedTime time.Time
}
