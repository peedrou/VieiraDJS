package converters

import (
	"time"
)

func ConvertExecutionTimeToUNIX(execution_time time.Time) int64 {
	converted_time := execution_time.Unix()
	converted_timeWithoutSeconds := converted_time / 60
	return converted_timeWithoutSeconds
}
