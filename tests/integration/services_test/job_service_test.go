package services_test

import (
	"testing"
	"time"

	db "VieiraDJS/app/db/session"
	"VieiraDJS/app/services/jobs"

	"github.com/stretchr/testify/assert"
)

func TestCreateJob(t *testing.T) {

	cassandra_session, err := db.CreateSession()
	if err != nil {
		t.Fatalf("Failed to create session: %v", err)
	}
	defer cassandra_session.Close()

	isRecurring := true
	maxRetries := 3
	startTime := time.Now()
	interval := "daily"

	err = jobs.CreateJob(cassandra_session, isRecurring, maxRetries, startTime, interval)

	assert.NoError(t, err, "CreateJob() returned an error: %v", err)
}
