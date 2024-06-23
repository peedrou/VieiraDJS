package jobs

// import (
// 	"VieiraDJS/app/helpers/builders"
// 	"VieiraDJS/app/helpers/validators"
// 	"time"

// 	"github.com/gocql/gocql"
// )

// func CreateJob(session *gocql.Session, isRecurring bool, maxRetries int, startTime time.Time, interval string) error {
// 	job, _ := builders.NewJob(
// 		isRecurring,
// 		maxRetries,
// 		startTime,
// 		interval,
// 	)

// 	if job != nil {
// 		response := InsertJobInDB(*job)
// 	}

// }

// func InsertJobInDB(job validators.ValidatedJob) error {

// }