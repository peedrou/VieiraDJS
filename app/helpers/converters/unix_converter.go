package converters

import (
	"strconv"
	"time"
)

func ConvertExecutionTimeToUNIX(execution_time time.Time) int64 {
	converted_time := execution_time.Unix()
	converted_timeWithoutSeconds := converted_time / 60
	return converted_timeWithoutSeconds
}

func CalculateNextExecutionTimeInUnix(interval string) int64 {
	currentTime := time.Now()
	numbers := interval[:len(interval)-1]
	character := interval[len(interval)-1:]

	value, err := strconv.Atoi(numbers)
	if err != nil {
		return 0
	}

	var nextTime time.Time
	if character == "M" {
		nextTime = currentTime.Add(time.Duration(value) * time.Minute)
	} else if character == "H" {
		nextTime = currentTime.Add(time.Duration(value) * time.Hour)
	} else {
		nextTime = currentTime.Add(time.Duration(value) * 24 * time.Hour)
	}

	return nextTime.Unix() / 60
}
