package converters

import (
	"time"
)

func ConvertExecutionTimeToUNIX(execution_time time.Time) int64 {
	converted_time := execution_time.Unix()
	return converted_time
}
