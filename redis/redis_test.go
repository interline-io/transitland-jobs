package jobs

import (
	"testing"

	"github.com/interline-io/transitland-dbutil/testutil"
	"github.com/interline-io/transitland-jobs/internal/jobtest"
	"github.com/interline-io/transitland-jobs/jobs"
)

func TestRedisJobs(t *testing.T) {
	// redis jobs and cache
	if a, ok := testutil.CheckTestRedisClient(); !ok {
		t.Skip(a)
		return
	}
	client := testutil.MustOpenTestRedisClient(t)
	newQueue := func(prefix string) jobs.JobQueue {
		q := jobs.NewJobLogger(NewRedisJobs(client, prefix))
		q.AddQueue("default", 4)
		return q
	}
	jobtest.TestJobQueue(t, newQueue)
}
