package jobs

import (
	"testing"

	"github.com/interline-io/transitland-jobs/internal/jobtest"
)

func TestLocalJobs(t *testing.T) {
	newQueue := func(queueName string) JobQueue {
		q := NewLocalJobs()
		q.AddQueue("default", 4)
		return q
	}
	jobtest.TestJobQueue(t, newQueue)
}
