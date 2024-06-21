package builders

import (
	"VieiraDJS/app/helpers/builders"
	"testing"
	"time"
)

func TestJobBuilder(t *testing.T) {
	_, err := builders.NewJob(
		true,
		2,
		time.Now(),
		"PT3H",
	)

	if err != nil {
		t.Errorf("NewJob() returned an unexpected error: %v", err)
	}
}
