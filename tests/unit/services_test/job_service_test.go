package services_test

import (
	"testing"
	"time"

	"VieiraDJS/app/services/jobs"

	"github.com/golang/mock/gomock"
)

func TestCreateJob(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock session
	mockSession := wrappers.NewMockSessionInterface(ctrl)

	mockQuery := wrappers.NewMockQueryInterface(ctrl) 

	mockSession.EXPECT().
		Query(gomock.Any(), gomock.Any()).
		Return(mockQuery) 

	mockQuery.EXPECT().Exec().Return(nil) 

	isRecurring := true
	maxRetries := 3
	startTime := time.Now()
	interval := "daily"

	err := jobs.CreateJob(mockSession, isRecurring, maxRetries, startTime, interval)

	if err != nil {
		t.Errorf("CreateJob() returned an error: %v", err)
	}
}