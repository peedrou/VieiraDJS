package models

import (
	"github.com/gocql/gocql"
)

type TaskSchedule struct {
	NextExecutionTime int64
	JobId             gocql.UUID
}
