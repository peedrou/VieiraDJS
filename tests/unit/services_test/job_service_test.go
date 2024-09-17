package services_test

import (
	"testing"
	"time"

	"VieiraDJS/tests/mocks"

	"VieiraDJS/app/services/jobs"

	"github.com/gocql/gocql"
	"github.com/golang/mock/gomock"
)

func TestCreateJob(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSessionCreator := mocks.NewMockSessionCreator(ctrl)
	mockSessionCreator.EXPECT().CreateSession().Return(&gocql.Session{}, nil)

	isRecurring := true
	maxRetries := 3
	startTime := time.Now()
	interval := "daily"

	// Call the function under test
	err := jobs.CreateJob(mockSessionCreator, isRecurring, maxRetries, startTime, interval)

	// Check for errors
	if err != nil {
		t.Errorf("CreateJob() returned an error: %v", err)
	}
}